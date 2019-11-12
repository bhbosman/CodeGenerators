using System;
using System.Diagnostics;
using System.Threading.Tasks;
using System.Threading.Tasks.Dataflow;

namespace ConsoleApp1
{
    public class QueueTimeActionBlock<T> : ITargetBlock<T>
    {
        private class ContextForActionBlock
        {
            public T Data { get; }
            public Stopwatch Stopwatch { get;  }
            public ContextForActionBlock(T data)
            {
                Data = data;
                Stopwatch = new Stopwatch();
                Stopwatch.Start();
            }
        }
        private readonly ITargetBlock<ContextForActionBlock> _actionBlock;
        private readonly string _connectionName;
        private readonly Action<T> _action;
        private double _queueTime;
        private double _processTime;
        private double _overallTime;


        public QueueTimeActionBlock(string connectionName, Action<T> action, ExecutionDataflowBlockOptions dataflowBlockOptions)
        {
            _connectionName = connectionName;
            _action = action;
            _actionBlock = new ActionBlock<ContextForActionBlock>(new Action<ContextForActionBlock>(Callback), dataflowBlockOptions);
        }

        private int _counter = 0;
        void Callback(ContextForActionBlock block)
        {
            block.Stopwatch.Stop();
            _queueTime += block.Stopwatch.Elapsed.TotalMilliseconds;
            block.Stopwatch.Start();
            try
            {
                Stopwatch processTime = Stopwatch.StartNew();
                try
                {
                    _action(block.Data);
                }
                finally
                {
                    processTime.Stop();
                    _processTime += processTime.Elapsed.TotalMilliseconds;
                }
            }
            finally
            {
                block.Stopwatch.Stop();
            }
           
            _overallTime += block.Stopwatch.Elapsed.TotalMilliseconds;

            _counter++;
            if (_counter % 1000 == 0)
            {
            //    Console.WriteLine($"{_connectionName}, {_counter},{_overallTime/ _counter} , {_queueTime/ _counter}, {_processTime/ _counter}");
            }
        }

        DataflowMessageStatus ITargetBlock<T>.OfferMessage(DataflowMessageHeader messageHeader, T messageValue, ISourceBlock<T> source,
            bool consumeToAccept)
        {
            ContextForActionBlock context = new ContextForActionBlock(messageValue);
            return _actionBlock.OfferMessage(messageHeader, context, null, consumeToAccept);
        }

        void IDataflowBlock.Complete()
        {
            _actionBlock.Complete();
        }

        void IDataflowBlock.Fault(Exception exception)
        {
            
        }

        Task IDataflowBlock.Completion => _actionBlock.Completion.ContinueWith(task =>
        {
            //Console.WriteLine($"{_connectionName}, {_counter},{_overallTime / _counter} , {_queueTime / _counter}, {_processTime / _counter}");
        });
    }
}