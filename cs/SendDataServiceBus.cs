using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Absa.Cib.ServiceBus;
using Absa.Cib.ServiceBus.Model;
using Absa.Cib.ServiceBus.Provider;
using Absa.Cib.ServiceBus.Serialization;
using Absa.Cib.ServiceBus.Serialization.Json;
using StackExchange.Redis;

namespace ConsoleApp1
{
    //public class SendDataServiceBus : ISendData
    //{
    //    private readonly IServiceBus _item;
    //    private IRedisBusConnectionProvider _redisBusConnectionProvider;
        
    //    private string _events = "Events:";

    //    public SendDataServiceBus(IServiceBus item, IRedisBusConnectionProvider redisBusConnectionProvider)
    //    {
    //        _item = item;
    //        _redisBusConnectionProvider = redisBusConnectionProvider;
    //    }
        

    //    public Task<bool> WriteHSetCommand(string myHash, string field1, string value)
    //    {
    //        return _redisBusConnectionProvider.GetConnection(0).HashSetAsync(myHash, field1, value);
    //    }


    //    public void Dispose()
    //    {
            
    //    }
    //}
}