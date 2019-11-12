using System;
using System.Diagnostics;
using System.IO;
using System.IO.Pipelines;
using System.Net.Sockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using System.Threading.Tasks.Dataflow;

namespace ConsoleApp1
{

    public class RedisConnectionMessageAnswer
    {
        public RedisConnectionMessageAnswer(string answer)
        {
            Answer = answer;
        }

        public string Answer { get;  }
    }
    public class RedisConnection: IRedisConnection
    {
        private readonly string _connectionName;

        private class QueuePackage
        {
            public string[] Command { get; }
            public TimeSpan TimeOut { get; }
            public Action<RedisConnectionMessageAnswer> Cb { get; }
            public string CbKey { get; }
            public TaskWrapper Task { get; }
            public CancellationTokenSource CancellationTokenSource { get; }
            public string Result { get; set; }
            public Stopwatch OverallStopwatch { get; }

            public QueuePackage(string[] command, TimeSpan timeOut, TaskWrapper task, CancellationTokenSource cancellationTokenSource, Action<RedisConnectionMessageAnswer> cb, string cbKey)
            {
                Command = command;
                TimeOut = timeOut;
                OverallStopwatch = new Stopwatch();
                OverallStopwatch.Start();
                Cb = cb;
                CbKey = cbKey;
                Task = task;
                CancellationTokenSource = cancellationTokenSource;
            }
        }

        class TaskWrapper
        {
            private string _ans;
            public TaskWrapper()
            {
                
                Task = new Task<string>(() => _ans);
            }

            public Task<string> Task { get; }

            public string Answer
            {
                set
                {
                    _ans = value;
                    Task.RunSynchronously();
                }
            }
        }

        public enum LastCommand
        {
            None,
            BlockRead,
            InCompleteHeader
        }

        private class ReadWaiter
        {
            public ReadWaiter(Action<RedisConnectionMessageAnswer> action, int dataSize, TaskWrapper task)
            {
                Action = action;
                DataSize = dataSize;
                Task = task;
            }
            public TaskWrapper Task { get; }
            public Action<RedisConnectionMessageAnswer> Action { get; }
            public int DataSize { get; }
        }


        private readonly TcpClient _tcpClient;
        private readonly byte[] _readBuffer;
        //private readonly ITargetBlock<QueuePackage> _writeDatActionBlock;
        private readonly CancellationTokenSource _cancellationTokenSource;
        private readonly Encoding _encoding;
        private bool _disposed;
        private int _messageReceivedCounter;
        private int _messageWrittenCounter;
        private TimeSpan _timeInWrite;
        private MemoryStream _memoryStream;
        
        private readonly JoinBlock<ReadWaiter, string> _answersJoinBlock;
        //private readonly ActionBlock<byte[]> _readSocketAction;
        private const int _readBufferSize = 8192 * 4;

        private LastCommand _lastCommand;

        public RedisConnection(string connectionName, string ip, CancellationTokenSource externalCancellationTokenSource)
        {
            _timeInWrite = new TimeSpan(0);
            _connectionName = connectionName;
            _encoding = Encoding.UTF8;
            _readBuffer = new byte[_readBufferSize];
            _disposed = false;
            _cancellationTokenSource = new CancellationTokenSource();
            _lastCommand = LastCommand.None;
//            _writeDatActionBlock = new ActionBlock<QueuePackage>(
//                package => { ___WriteToSocket___(package);},
//                new ExecutionDataflowBlockOptions
//                {
//                    SingleProducerConstrained = true,
//                    MaxDegreeOfParallelism = 1,
//                    BoundedCapacity = 4096,
//                    CancellationToken = _cancellationTokenSource.Token
//                });


            _answersJoinBlock = new JoinBlock<ReadWaiter, string>();
            ActionBlock<Tuple<ReadWaiter, string>> dd = new ActionBlock<Tuple<ReadWaiter, string>>(
                tuple =>
                {
                    tuple.Item1.Task.Answer = tuple.Item2;
                    if (tuple.Item1.Action != null)
                    {
                        tuple.Item1.Action.Invoke(new RedisConnectionMessageAnswer(tuple.Item2));
                    }
                }, 
                new ExecutionDataflowBlockOptions
                {
                    MaxDegreeOfParallelism = 1
                });

            _answersJoinBlock.LinkTo(dd);


            //_readSocketAction = new ActionBlock<byte[]>(
            //    data => { ___ProcessingIncomingData_____(data); });

            externalCancellationTokenSource?.Token.Register(Dispose);
            _tcpClient = new TcpClient(ip, 6379);
            _tcpClient.NoDelay = true;
            Read();
        }


        private void ___ProcessingIncomingData_____(byte[] data)
        {
            int beginPoint = 0;
            while (true)
            {
                if (_lastCommand == LastCommand.None)
                {
                    void HeaderMissing()
                    {
                        _lastCommand = LastCommand.InCompleteHeader;
                        _memoryStream = new MemoryStream();
                        _memoryStream.Capacity = _readBufferSize * 2;
                        _memoryStream.Write(data, beginPoint, data.Length - beginPoint);
                    }

                    int SizeLeft()
                    {
                        return data.Length - beginPoint;
                    }
                    if (beginPoint < data.Length)
                    {
                        byte command = data[beginPoint];
                        if (command == '-')
                        {
                            beginPoint++;
                            Interlocked.Increment(ref _messageReceivedCounter);
                            int index = Array.FindIndex(data, beginPoint, b => b == '\n' || b == '\r');
                            int length = index - beginPoint;
                            string s = _encoding.GetString(data, beginPoint, length);
                            _answersJoinBlock.Target2.Post(s);
                            beginPoint += length;
                            beginPoint += 2;
                            if (beginPoint == data.Length)
                            {
                                break;
                            }
                            continue;
                        }
                        if (command == ':') // ans for write
                        {
                            int beginPointRollbackPoint = beginPoint;
                            beginPoint++;
                            Interlocked.Increment(ref _messageReceivedCounter);
                            int index = Array.FindIndex(data, beginPoint, b => b == '\n' || b == '\r');
                            if (index >= 0)
                            {
                                int length = index - beginPoint;
                                string s = _encoding.GetString(data, beginPoint, length);
                                beginPoint += length;
                                if (data.Length - beginPoint < 2)
                                {
                                    beginPoint = beginPointRollbackPoint;
                                    HeaderMissing();
                                    break;
                                }
                                beginPoint += 2;
                                _answersJoinBlock.Target2.Post(s);
                                if (beginPoint == data.Length)
                                {
                                    break;
                                }
                                continue;
                            }

                            beginPoint = beginPointRollbackPoint;
                            HeaderMissing();
                            break;

                        }
                        if (command == '$') // ans for read
                        {
                            int beginPointRollbackPoint = beginPoint;
                            beginPoint++;
                            Interlocked.Increment(ref _messageReceivedCounter);
                            int index = Array.FindIndex(data, beginPoint, b => b == '\n' || b == '\r');
                            if (index >= 0)
                            {
                                int dataSize = 0;
                            
                                int length = index - beginPoint;
                                string s = _encoding.GetString(data, beginPoint, length);
                                beginPoint += length;
                                if (data.Length - beginPoint < 2)
                                {
                                    beginPoint = beginPointRollbackPoint;
                                    HeaderMissing();
                                    break;
                                }
                                beginPoint += 2;
                                dataSize = int.Parse(s);
                            
                                

                                if (dataSize == -1)
                                {
                                    _answersJoinBlock.Target2.Post(null);
                                    if (beginPoint == data.Length)
                                    {
                                        break;
                                    }
                                    continue;
                                }
                                int leftOver = data.Length - beginPoint;
                                if (dataSize + 2 <= leftOver)
                                {
                                    string stringData = _encoding.GetString(data, beginPoint, dataSize);
                                    _answersJoinBlock.Target2.Post(stringData);
                                    beginPoint += dataSize;
                                    beginPoint += 2;
                                    if (beginPoint == data.Length)
                                    {
                                        break;
                                    }
                                    continue;
                                }
                                _memoryStream = new MemoryStream();
                                _memoryStream.Capacity = dataSize + 2;
                                if (leftOver > 0)
                                {
                                    _memoryStream.Write(data, beginPoint, leftOver);
                                }
                                _lastCommand = LastCommand.BlockRead;
                                break;
                            }
                            beginPoint = beginPointRollbackPoint;
                            HeaderMissing();
                            break;
                        }
                        throw new Exception("implement command");

                        //                    throw new Exception();

                        //_lastCommand = LastCommand.InCompleteHeader;
                        //int leftOver2 = data.Length - beginPoint;
                        //_memoryStream = new MemoryStream();
                        //_memoryStream.Capacity = _readBufferSize * 2;
                        //_memoryStream.Write(data, beginPoint, leftOver2);
                        //_lastCommand = LastCommand.BlockRead;
                        //break;
                    }
                }

                if (_lastCommand == LastCommand.InCompleteHeader)
                {
                    _memoryStream.Write(data, 0, data.Length);
                    data = _memoryStream.ToArray();
                    _memoryStream.Dispose();
                    _memoryStream = null;
                    _lastCommand = LastCommand.None;
                    continue;
                }
                if (_lastCommand == LastCommand.BlockRead)
                {
                    if (_memoryStream == null)
                    {
                        throw new Exception();
                    }
                    int needed = (int)_memoryStream.Capacity - (int)_memoryStream.Position;
                    if (needed > 0)
                    {
                        if (data.Length >= needed)
                        {
                            _memoryStream.Write(data, 0, needed);
                            beginPoint += needed;
                            string s = _encoding.GetString(_memoryStream.ToArray(), 0, _memoryStream.Capacity - 2);
                            _answersJoinBlock.Target2.Post(s);
                            _memoryStream.Dispose();
                            _memoryStream = null;
                            _lastCommand = LastCommand.None;
                            if (beginPoint == data.Length)
                            {
                                break;
                            }
                        }
                        else
                        {
                            _memoryStream.Write(data, 0, data.Length);
                            break;
                        }
                    }
                    else if (data.Length >= 2)
                    {
                        beginPoint += 2;

                        string s = _encoding.GetString(_memoryStream.ToArray());
                        _answersJoinBlock.Target2.Post(s);
                        _memoryStream.Dispose();
                        _memoryStream = null;
                        _lastCommand = LastCommand.None;
                        if (beginPoint == data.Length)
                        {
                            break;
                        }
                    }
                }
            }
        }

        private void ___WriteToSocket___(QueuePackage package)
        {
            if (!_disposed)
            {
                try
                {
                    Interlocked.Increment(ref _messageWrittenCounter);
                    byte[] byteArray = SerializeCommand(package.Command);
                    int n = WriteToSocket(byteArray);
                    _answersJoinBlock.Target1.Post(new ReadWaiter(package.Cb, n, package.Task));
                }
                catch (Exception)
                {
                    // ignored
                }
                package.OverallStopwatch.Stop();
                _timeInWrite = _timeInWrite.Add(package.OverallStopwatch.Elapsed);
            }
            else
            {
                package.CancellationTokenSource.Cancel();
            }
        }


        private byte[] SerializeCommand(string[] command)
        {
            int commandLength = 1;
            commandLength += _encoding.GetByteCount(command.Length.ToString());
            commandLength += 2;
            foreach (string s in command)
            {
                commandLength += _encoding.GetByteCount(s.Length.ToString());
                commandLength += _encoding.GetByteCount(s);
                commandLength += 5;
            }

            byte[] byteArray = new byte[commandLength];
            int beginPos = 0;
            // write array header
            byteArray[beginPos] = (byte)'*'; beginPos++;
            string temp = command.Length.ToString();
            _encoding.GetBytes(temp, 0, temp.Length, byteArray, beginPos); beginPos += temp.Length;
            byteArray[beginPos] = (byte)'\r'; beginPos++;
            byteArray[beginPos] = (byte)'\n'; beginPos++;
            foreach (string s in command)
            {
                byteArray[beginPos] = (byte)'$'; beginPos++;
                string cmdLength = s.Length.ToString();
                _encoding.GetBytes(cmdLength, 0, cmdLength.Length, byteArray, beginPos);
                beginPos += cmdLength.Length;
                byteArray[beginPos] = (byte)'\r'; beginPos++;
                byteArray[beginPos] = (byte)'\n'; beginPos++;

                _encoding.GetBytes(s, 0, s.Length, byteArray, beginPos);
                beginPos += s.Length;
                byteArray[beginPos] = (byte)'\r'; beginPos++;
                byteArray[beginPos] = (byte)'\n'; beginPos++;
            }

            return byteArray;
        }

        private int WriteToSocket(byte[] byteArray)
        {

            if (!_disposed)
            {
                int begin = 0;
                int length = byteArray.Length;
                while (true)
                {
                    int leftOver = length - begin;
                    int n = _tcpClient.Client.Send(byteArray, begin, leftOver, SocketFlags.None);
                    begin += n;
                    if (begin == length)
                    {
                        break;
                    }
                }
                return length;
            }
            else
            {
                throw new TaskCanceledException("Connection to service closed");
            }
        }

        public Task<string> WriteToSocket(string[] command, TimeSpan timeOut, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null)
        {
            try
            {
                if (!_disposed)
                {
                    CancellationTokenSource cancellationToken = new CancellationTokenSource(timeOut);
                    TaskWrapper task = new TaskWrapper();

                    Post(new QueuePackage(command, timeOut, task, cancellationToken, cb, cbKey));
                    return task.Task;
                }
                cb?.Invoke(new RedisConnectionMessageAnswer(""));
                return Task.FromCanceled<string>(new CancellationToken(true));
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
                throw;
            }
            
        }

        public bool Active => !_disposed;

        private void Read()
        {
            if (!_disposed)
            {
                try
                {
                    _tcpClient.Client.BeginReceive(
                        _readBuffer,
                        0, _readBuffer.Length,
                        SocketFlags.None,
                        ReadFromSocket, null);
                }
                catch (Exception)
                {
                    Dispose();
                }
            }
        }

        private void ReadFromSocket(IAsyncResult ar)
        {
            if (!_disposed)
            {
                try
                {
                    int n = _tcpClient.Client.EndReceive(ar);
                    if (n > 0)
                    {
                        //Console.Write('R');
                        byte[] data = new byte[n];
                        Array.Copy(_readBuffer, 0, data, 0, n);
                        //_readSocketAction.Post(data);
                        ___ProcessingIncomingData_____(data);
                        
                     
                    }
                    Read();
                }
                catch (Exception e)
                {
                    Dispose();
                }
            }
        }

        public void Dispose()
        {
            if (!_disposed)
            {
                _cancellationTokenSource.Cancel();

                void ActionWrappedInExceptionLogic(Action act)
                {
                    try
                    {
                        act();
                    }
                    catch (Exception)
                    {
                        // ignored
                    }
                }
                _disposed = true;
                ActionWrappedInExceptionLogic(_tcpClient.Close);
                ActionWrappedInExceptionLogic(_tcpClient.Dispose);

                //Console.WriteLine($"{_connectionName}: {_timeInWrite}, {_messageWrittenCounter}, {_messageReceivedCounter}");
                //ActionWrappedInExceptionLogic(_writeDatActionBlock.Complete);
                //try
                //{
                //    _writeDatActionBlock.Completion.Wait(new TimeSpan(0, 0, 0, 1));
                //}
                //catch (Exception)
                //{
                //    // ignored
                //}
            }
        }

        

        private void Post(QueuePackage queuePackage)
        {
            //while (true)
            //{
            //    if (!_writeDatActionBlock.Post(queuePackage))
            //    {
            //       Thread.Sleep(100);
            //        continue;
            //    }
            //    break;
            //}
            ___WriteToSocket___(queuePackage);
            //Console.Write('.');
        }
    }
}