using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using System.Threading.Tasks.Dataflow;
using Absa.Cib.ServiceBus;
using Absa.Cib.ServiceBus.Model;
using Absa.Cib.ServiceBus.Provider;
using Bagl.Cib.MIT.IoC;
using StackExchange.Redis;
using Unity.Interception.Utilities;

namespace ConsoleApp1
{
    public class WhatEver : IWhatEver
    {
        private readonly ISystemInformation _systemInformation;
        public WhatEver(ISystemInformation systemInformation)
        {
            _systemInformation = systemInformation;
        }



        int _outstandingMessages = 0;
        int _produceCount = 0;
        int _consumeCount = 0;
        void ActualTest(ConcurrentStack<ISendData> stack, int messageCount, int packageSize, int maxDegreeOfParallelism, int connectionCount, Action<ISendData, ManualResetEvent> act)
        {
            
            ManualResetEvent ev = new ManualResetEvent(false);
            ActionBlock<string> sendDataToRedisBlock = new ActionBlock<string>(
                s =>
                {
                    try
                    {
                        while (true)
                        {
                            if (stack.TryPop(out ISendData item))
                            {
                                try
                                {
                                    if (!item.Active)
                                    {
                                        continue;
                                    }
                                    if (item.Active)
                                    {
                                        try
                                        {
                                            Interlocked.Increment(ref _produceCount);
                                            Interlocked.Increment(ref _outstandingMessages);
                                            act(item, ev);
                                        }
                                        catch (Exception e)
                                        {
                                            // ignored
                                        }
                                    }
                                }
                                finally
                                {
                                    if (item.Active)
                                    {
                                        stack.Push(item);
                                    }
                                }
                                break;
                            }
                        }
                    }
                    catch (Exception)
                    {
                        // ignored
                    }
                },
                new ExecutionDataflowBlockOptions
                {
                    MaxDegreeOfParallelism = maxDegreeOfParallelism,
                });


            
            Stopwatch overallStopwatch = new Stopwatch();
            overallStopwatch.Start();
            
            for (int i = 0; i < messageCount; i++)
            {
                Interlocked.Increment(ref _outstandingMessages);
                while (true)
                {
                    if (!sendDataToRedisBlock.Post(""))
                    {
                        continue;
                    }
                    break;
                }
                Interlocked.Decrement(ref _outstandingMessages);
            }

            sendDataToRedisBlock.Complete();
            Task.WaitAll(sendDataToRedisBlock.Completion);
            ev.WaitOne();
            overallStopwatch.Stop();
            Console.WriteLine($"({messageCount}, {packageSize}, {maxDegreeOfParallelism}, {connectionCount}): {overallStopwatch.Elapsed}, {overallStopwatch.Elapsed.TotalMilliseconds/messageCount}");
        }






        


        

        public void WriteToQueueUsingSocket(string queueName, int messageCount, int packageSize, int maxDegreeOfParallelism, int connectionCount)
        {
            Console.WriteLine($"Begin {packageSize}");
            ConcurrentStack<ISendData> stack = new ConcurrentStack<ISendData>();
            for (int i = 0; i < connectionCount; i++)
            {

                {
                    IRedisConnection item = new RedisConnection($"connection_{i}", "22.242.118.171", null);
                    stack.Push(new SendDataRedisConnection(item));
                }

                //{
                //    IRedisConnection item = new RedisConnection($"connection_{i}", "127.0.0.1", null);
                //    stack.Push(new SendDataRedisConnection(item));
                //}
                //{
                //    IRedisConnection item = new RedisConnection("10.114.94.98");
                //    stack.Push(new SendDataRedisConnection(item));
                //}
            }
            try
            {
                StringBuilder sb = new StringBuilder(packageSize);
                sb.Append('a', packageSize);
                string sbValue = sb.ToString();
                //ActualTest(stack, messageCount, packageSize, maxDegreeOfParallelism, connectionCount,
                //    (item, ev) =>
                //    {
                //        item.WriteHSetCommand("ABC", Guid.Empty.ToString(), sbValue,
                //            (answer) =>
                //            {
                //                Interlocked.Increment(ref _consumeCount);
                //                int value = Interlocked.Decrement(ref _outstandingMessages);
                //                if (value == 0)
                //                {
                //                    ev.Set();
                //                }
                //            });
                //    });


                if (stack.TryPop(out ISendData writeItem))
                {
                    try
                    {
                        writeItem.WriteHSetCommand("DEF", "ABC", sbValue).Wait();
                    }
                    finally
                    {
                        stack.Push(writeItem);
                    }
                }



                ActualTest(stack, messageCount, packageSize, maxDegreeOfParallelism, connectionCount,
                    (item, ev) =>
                    {
                        var dd = item.GetHSetCommand("DEF", "ABC",
                            (answer) =>
                            {
                                if (answer.Answer != sbValue)
                                {
                                    Console.WriteLine("read errr");
                                }
                                Interlocked.Increment(ref _consumeCount);
                                int value = Interlocked.Decrement(ref _outstandingMessages);
                                if (value == 0)
                                {
                                    ev.Set();
                                }
                            });

                       // dd.Wait();
                    });
            }
            finally
            {
                while (true)
                {
                    if (stack.TryPop(out ISendData item))
                    {
                        item.Dispose();
                        continue;
                    }
                    break;
                }
            }
            Console.WriteLine($"finish {packageSize}");
        }

        async Task WriteAction(ConcurrentQueue<ISendData> stack, string stringData)
        {
            List<Task> tasks = new List<Task>();
            for (int i = 0; i < 60; i++)
            {
                Task<string> t = null;
                if (stack.TryDequeue(out ISendData item))
                {
                    try
                    {
                        t = item.WriteHSetCommand("ABC", i.ToString(), stringData);
                    }
                    finally
                    {
                        stack.Enqueue(item);
                    }

                    tasks.Add(t);
                    Interlocked.Add(ref overallWriteCount, 1);
                }
            }

            await Task.WhenAll(tasks.ToArray());
        }
        int overallWriteCount = 0;
        int overallReadCount = 0;


        public void SharedPool(int size, int connections, int readConsumers, int writeProducers)
        {

            int len = size * 1024;
            StringBuilder sbBuilder = new StringBuilder(len);
            sbBuilder.Append((char)(97), len);
            string stringData = sbBuilder.ToString();


            Console.WriteLine($"Begin ");
            ConcurrentQueue<ISendData> stack = new ConcurrentQueue<ISendData>();
            for (int i = 0; i < connections; i++)
            {
                //RedisConnection item = new RedisConnection($"connection_{i}", "127.0.0.1", null);
                IRedisConnection item = new RedisConnection($"connection_{i}", "22.242.118.171", null);
                //IRedisConnection item = new RedisConnection($"connection_{i}", "JHBPSM020001301", null);
                
               stack.Enqueue(new SendDataRedisConnection(item));
            }
            try
            {

                Task t = WriteAction(stack, stringData);
                //t.Start();
                t.Wait();

                overallWriteCount = 0;
                overallReadCount = 0;

                CancellationTokenSource cancellationTokenSource = new CancellationTokenSource(new TimeSpan(0, 0, 0, 30));
                List<Task> outerTasks = new List<Task>();
                for (int writeCount = 0; writeCount < writeProducers; writeCount++)
                {
                    outerTasks.Add(Task.Run(async () =>
                    {
                        

                        while (!cancellationTokenSource.IsCancellationRequested)
                        {
                            await WriteAction(stack, stringData);
                        }
                    }));
                }

                
                for (int counter = 0; counter < readConsumers; counter++)
                {
                    outerTasks.Add(Task.Run(async () =>
                    {
                        while (!cancellationTokenSource.IsCancellationRequested)
                        {
                            List<Task> tasks = new List<Task>();
                            for (int i = 0; i < 15; i++)
                            {
                                if (cancellationTokenSource.IsCancellationRequested)
                                {
                                    break;
                                }
                                Task<string> task=null;
                                if (stack.TryDequeue(out ISendData item))
                                {
                                    try
                                    {
                                        
                                        task = item.GetHSetCommand("ABC", i.ToString());
                                    }
                                    finally
                                    {
                                        stack.Enqueue(item);
                                    }
                                    tasks.Add(task);
                                    Interlocked.Add(ref overallReadCount, 1);
                                }
                            }
                            await Task.WhenAll(tasks.ToArray());
                        }
                    }, cancellationTokenSource.Token));
                }
                Task.WaitAll(outerTasks.ToArray());
            }
            finally
            {
                while (true)
                {
                    if (stack.TryDequeue(out ISendData item))
                    {
                        item.Dispose();
                        continue;
                    }
                    break;
                }
            }
            Console.WriteLine($"finish:({size}, {connections}, {readConsumers}, {writeProducers}) {overallWriteCount}, {overallReadCount}");
        }

        public void SeperatePools(int size, int readConnections, int writeConnections, int readers, int writers)
        {
            int len = size * 1024;
            StringBuilder sbBuilder = new StringBuilder(len);
            sbBuilder.Append((char)(97), len);
            string stringData = sbBuilder.ToString();

            string PC = "22.242.118.171";
            //string PC = "JHBPSM020001301";
            Console.WriteLine($"Begin ");
            ConcurrentQueue<ISendData> readerStack = new ConcurrentQueue<ISendData>();
            for (int i = 0; i < readConnections; i++)
            {
                IRedisConnection item = new RedisConnection($"connection_{i}", PC, null);
                readerStack.Enqueue(new SendDataRedisConnection(item));
            }

            ConcurrentQueue<ISendData> writerStack = new ConcurrentQueue<ISendData>();
            for (int i = 0; i < readConnections; i++)
            {
                IRedisConnection item = new RedisConnection($"connection_{i}", PC, null);
                writerStack.Enqueue(new SendDataRedisConnection(item));
            }

            try
            {

                Task t = WriteAction(writerStack, stringData);
                //t.Start();
                t.Wait();

                overallWriteCount = 0;
                overallReadCount = 0;

                CancellationTokenSource cancellationTokenSource = new CancellationTokenSource(new TimeSpan(0, 0, 0, 30));
                List<Task> outerTasks = new List<Task>();
                for (int writeCount = 0; writeCount < writers; writeCount++)
                {
                    outerTasks.Add(Task.Run(async () =>
                    {


                        while (!cancellationTokenSource.IsCancellationRequested)
                        {
                            await WriteAction(writerStack, stringData);
                        }
                    }));
                }


                for (int counter = 0; counter < readers; counter++)
                {
                    outerTasks.Add(Task.Run(async () =>
                    {
                        while (!cancellationTokenSource.IsCancellationRequested)
                        {
                            List<Task> tasks = new List<Task>();
                            for (int i = 0; i < 15; i++)
                            {
                                if (cancellationTokenSource.IsCancellationRequested)
                                {
                                    break;
                                }
                                Task<string> task = null;
                                if (readerStack.TryDequeue(out ISendData item))
                                {
                                    try
                                    {

                                        task = item.GetHSetCommand("ABC", i.ToString());
                                    }
                                    finally
                                    {
                                        readerStack.Enqueue(item);
                                    }
                                    tasks.Add(task);
                                    Interlocked.Add(ref overallReadCount, 1);
                                }
                            }
                            await Task.WhenAll(tasks.ToArray());
                        }
                    }, cancellationTokenSource.Token));
                }
                Task.WaitAll(outerTasks.ToArray());
            }
            finally
            {
                while (true)
                {
                    if (readerStack.TryDequeue(out ISendData item))
                    {
                        item.Dispose();
                        continue;
                    }
                    break;
                }
                while (true)
                {
                    if (writerStack.TryDequeue(out ISendData item))
                    {
                        item.Dispose();
                        continue;
                    }
                    break;
                }
            }
            Console.WriteLine($"finish:({size}, {readConnections},{writeConnections}, {readers}, {writers}) {overallWriteCount}, {overallReadCount}");
        }

        public void ddddd(int size, int connections, int readers, int writers)
        {
            StringBuilder sb = new StringBuilder();
            sb.Append('a', size * 1024);
            //string sbValue = sb.ToString();
            string sbValue = "ABC";

            ConnectionMultiplexer connection = ConnectionMultiplexer.Connect("22.242.118.171");
            int sendDataOverall = 0;
            List<Task> outerList = new List<Task>();
            for (int i = 0; i < 20; i++)
            {
                outerList.Add(
                    Task.Run(async () =>
                    {

                        int sendData = 0;
                        for (int ii = 0; i < 1000000; i++)
                        {
                            var t = connection.GetDatabase(0).HashSetAsync("ABC", "DEF", sbValue, flags: CommandFlags.FireAndForget);
                            sendData += sbValue.Length;
                            //l.Add(t);

                        }
                        //await Task.WhenAll(l.ToArray());
                        Interlocked.Add(ref sendDataOverall, sendData);
                    }));
            }
            


            Task.WaitAll(outerList.ToArray());


            Console.WriteLine($"ddd {sendDataOverall}");


        }

        
    }
}