using System;
using System.Collections;
using System.IO;
using System.Threading;
using System.Threading.Tasks;
using Absa.Cib.ServiceBus;
using Absa.Cib.ServiceBus.Provider;
using Bagl.Cib.MIT.IoC;
using Bagl.Cib.MIT.IoC.Models;
using Bagl.Cib.MIT.Logging;
using Bagl.Cib.MIT.Redis;
using Bagl.Cib.MIT.Redis.Serialization;
using Bagl.Cib.MSF.Contracts.Model;
using StackExchange.Redis;
using Unity;
using Unity.Injection;
using Unity.Registration;
using Unity.Resolution;

namespace ConsoleApp1
{
    class Program
    {
        static void Test0001(ISystemInformation systemInformation)
        {
            IWhatEver we = systemInformation.Resolve<IWhatEver>();
            //we.SharedPool(1, 80, 400, 40);
            //we.SharedPool(2, 80, 400, 40);
            //we.SharedPool(4, 80, 400, 40);
            //we.SharedPool(8, 80, 400, 40);
            //we.SharedPool(16, 80, 400,40);
            //we.SharedPool(32, 80, 400, 40);
            //we.SharedPool(64, 80, 400, 40);
            //we.SharedPool(128, 80, 400, 40);

            we.SharedPool(128, 1, 1, 0);
        }
        static void Test0002(ISystemInformation systemInformation)
        {
            IWhatEver we = systemInformation.Resolve<IWhatEver>();
            //we.SeperatePools(1, 40, 40, 0, 400);
            //we.SeperatePools(2, 40, 40,0, 400);
            //we.SeperatePools(4, 40, 40, 0, 400);
            //we.SeperatePools(8, 40, 40, 0, 400);
            //we.SeperatePools(16, 40, 40, 0, 400);
            we.SeperatePools(32, 40, 40, 0, 400);
            //we.SeperatePools(64, 40, 40, 0, 400);
            //we.SeperatePools(128, 40, 40, 0, 400);
            //we.SeperatePools(256, 40, 40, 0, 400);
            //we.SeperatePools(512, 40, 40, 0, 400);
            //we.SeperatePools(1024, 40, 40, 0, 400);
        }

        static void Test0003(ISystemInformation systemInformation)
        {
            IWhatEver we = systemInformation.Resolve<IWhatEver>();
            we.ddddd(1, 1, 100, 100);
        }

        //static void Test0002(ISystemInformation systemInformation)
        //{
        //    IWhatEver we = systemInformation.Resolve<IWhatEver>();

        //    we.WriteToQueueUsingSocket("ABC", 50000, (int)Math.Pow(2, 8), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 50000, (int)Math.Pow(2, 9), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 50000, (int)Math.Pow(2, 10), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 50000, (int)Math.Pow(2, 11), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 50000, (int)Math.Pow(2, 12), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 10000, (int)Math.Pow(2, 13), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 10000, (int)Math.Pow(2, 14), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 10000, (int)Math.Pow(2, 15), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 10000, (int)Math.Pow(2, 16), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 10000, (int)Math.Pow(2, 17), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 10000, (int)Math.Pow(2, 18), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 10000, (int)Math.Pow(2, 19), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 10000, (int)Math.Pow(2, 20), 150, 50);
        //    we.WriteToQueueUsingSocket("ABC", 10000, (int)Math.Pow(2, 21), 150, 50);

        //    Console.WriteLine("finish");
        //    Console.ReadLine();

        //}
        static void Main(string[] args)
        {
            ISystemInformation systemInformation = new SystemInformation("ddd", "UAT", SessionKeyType.Application, null, null);

            Absa.Cib.MIT.Redis.IoC.Registrations.Register(systemInformation);
            Bagl.Cib.MSF.ClientAPI.Registrations.Register(systemInformation);
            Absa.Cib.JwtAuthentication.Registrations.Register(systemInformation);
            Absa.Cib.ServiceBus.IoC.Registrations.Register(systemInformation);
            systemInformation.RegisterType<IWhatEver, WhatEver>(Scope.Instance);

            systemInformation.RegisterType<IServiceBus, RedisServiceBus>(Scope.ContainerSingleton);
            systemInformation.RegisterType<IRedisBusConnectionProvider>(Scope.ContainerSingleton, new InjectionMember[1]
            {
                (InjectionMember) new InjectionFactory((Func<IUnityContainer, object>) (e =>
                {
                    RedisConfiguration redisConfiguration = (RedisConfiguration) e.Resolve(typeof (RedisConfiguration), (string) null, (ResolverOverride[]) null);
                    return (object) new RedisBusConnectionProvider((ILoggingService) e.Resolve(typeof (ILoggingService), (string) null, (ResolverOverride[]) null), redisConfiguration.RedisOptions);
                }))
            });

            ILoggingService loggingService = systemInformation.Resolve<ILoggingService>();
            loggingService.Initialize(systemInformation.LoggingInformation);


            //Test0001(systemInformation);
            //Test0002(systemInformation);
            Test0003(systemInformation);

            //using (IRedisConnection conn = new RedisConnection())
            //{
            //    conn.WriteMultiCommand();
            //    conn.WriteAppendCommand("AA", "ABC");
            //    conn.WriteAppendCommand("AA", "ABC");
            //    conn.WriteAppendCommand("AA", "ABC");
            //    conn.WriteAppendCommand("AA", "ABC");
            //    conn.WriteAppendCommand("AA", "ABC");
            //    conn.WriteAppendCommand("AA", "ABC");
            //    conn.WriteAppendCommand("AA", "ABC");
            //    conn.WriteAppendCommand("AA", "ABC");
            //    conn.WriteHSetCommand("ABC", "DEF", "GHI");
            //    conn.WriteHSetCommand("ABC", "DEF", "JKL");
            //    conn.WriteHSetCommand("ABC", "DEF", "MNO");
            //    conn.WriteExecCommand();
            //    conn.Complete();
            //    conn.Wait();
            //}

            Console.ReadLine();

        }
    }
}
