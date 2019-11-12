using System;
using System.Threading.Tasks;
using Absa.Cib.ServiceBus.Model;

namespace ConsoleApp1
{
    public interface IRedisConnection: IDisposable
    {
        Task<string> WriteToSocket(string[] command, TimeSpan timeOut, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null);
        bool Active { get; }
    }

    public static class RedisConnectionExtensions
    {
        public static Task<string> WriteMultiCommand(this IRedisConnection connection, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null)
        {
            return connection.WriteToSocket(new []{"MULTI"}, TimeSpan.FromSeconds(60), cb, cbKey);
        }

        public static Task<string> WriteExecCommand(this IRedisConnection connection, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null)
        {
            return connection.WriteToSocket(new []{ "EXEC" }, TimeSpan.FromSeconds(60), cb, cbKey);
        }

        public static void WriteAppendCommand(this IRedisConnection connection, string key, string value, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null)
        {
            connection.WriteToSocket(new[]{"APPEND", key, value}, TimeSpan.FromSeconds(60), cb, cbKey);
        }

        public static Task<string> WriteHSetCommand(this IRedisConnection connection, string myHash, string field1,
            string value, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null)
        {
            return connection.WriteToSocket(new[] { "HSET", myHash,field1, value}, TimeSpan.FromSeconds(60), cb, cbKey);
        }

        public static Task<string> GetHSetCommand(this IRedisConnection connection, string myHash, string field1, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null)
        {
            return connection.WriteToSocket(new []{"HGET", myHash, field1}, TimeSpan.FromSeconds(60), cb, cbKey);
        }

        public static Task<string> WriteHSetNxCommand(this IRedisConnection connection, string myHash, string field1, string value, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null)
        {
            return connection.WriteToSocket(new[]{"HSETNX", myHash, field1, value}, TimeSpan.FromSeconds(60), cb, cbKey);
        }
    }
}