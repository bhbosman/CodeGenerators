// SCRIPT LOAD 'local result = redis.call("SET", ARGV[1], ARGV[2], "EX", ARGV[3]) redis.call("RPUSH", ARGV[4], ARGV[5]) redis.call("PUBLISH", ARGV[4], 1) result = redis.call("ZINCRBY", ARGV[6], ARGV[7], ARGV[8]) result = redis.call("SET", ARGV[9] .. ":" .. ARGV[8], ARGV[10]) redis.call("PUBLISH", ARGV[11], ARGV[8]) return result'



using System;
using System.CodeDom.Compiler;
using System.Threading.Tasks;
using Absa.Cib.ServiceBus.Model;
using StackExchange.Redis;

namespace ConsoleApp1
{
    public class SendDataRedisConnection : ISendData
    {
        
        private readonly IRedisConnection _item;



        public SendDataRedisConnection(IRedisConnection item)
        {
            _item = item;
            
        }

        public bool Active => _item.Active;
    

        public Task<string> WriteHSetCommand(string myHash, string field1, string value, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null)
        {
            return _item.WriteHSetCommand(myHash, field1, value, cb, cbKey);
        }

        public Task<string> GetHSetCommand(string myHash, string field1, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null)
        {
            return _item.GetHSetCommand(myHash, field1, cb, cbKey);
        }

        public void Dispose()
        {
            try
            {
                _item.Dispose();
            }
            catch (Exception e)
            {
                // ignored
            }
        }
    }
}