namespace ConsoleApp1
{
    public interface IWhatEver
    {
        void WriteToQueueUsingSocket(string queueName, int messageCount, int packageSize, int maxDegreeOfParallelism, int connectionCount);
        void SharedPool(int size, int connections, int readers, int writers);
        void SeperatePools(int size, int readConnections, int writeConnections, int readers, int writers);
        void ddddd(int size, int connections, int readers, int writers);

    }
}