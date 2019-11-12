using System;
using System.Threading.Tasks;
using Absa.Cib.ServiceBus.Model;

namespace ConsoleApp1
{
    public interface ISendData: IDisposable
    {
        bool Active { get; }
        Task<string> WriteHSetCommand(string myHash, string field1, string value, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null);

        Task<string> GetHSetCommand(string myHash, string field1, Action<RedisConnectionMessageAnswer> cb = null, string cbKey = null);
    }
}