package getcdv3

//
//import (
//	"TestApi/common/config"
//	"TestApi/common/utils"
//	"fmt"
//	clientv3 "go.etcd.io/etcd/client/v3"
//	"golang.org/x/net/context"
//	"sync"
//	"time"
//)
//
//var (
//	healthProvider     *HealthProvider
//	healthProviderOnce sync.Once
//)
//
//type HealthProvider struct {
//	EtcdClient *EtcdClient
//}
//
//type EtcdClient struct {
//	address  []string
//	kv       clientv3.KV
//	client   *clientv3.Client
//	ctx      context.Context
//	lease    clientv3.Lease
//	leaseID  clientv3.LeaseID
//	leaseTTL int64
//	username string
//	password string
//}
//
//func GetHealthProvider() *HealthProvider {
//	healthProviderOnce.Do(func() {
//		healthProvider = &HealthProvider{
//			EtcdClient: NewEtcdClientnew(10),
//		}
//	})
//	return healthProvider
//}
//
//func Registernew() {
//	provider := GetHealthProvider()
//	go HealthCheck(provider)
//	defer provider.EtcdClient.Close()
//}
//
////k-v 的key前缀
////格式：   %s:///%s/
//func GetKeyPrefix(scheme, serviceName string) string {
//	return fmt.Sprintf("%s:///%s/", scheme, serviceName)
//}
//
////传入etcd名，地址，服务ip，端口号，注册的服务名称，租约过期时间
////func NewEtcdClintall(schema, etcdAddr string, serverIp, port, serviceName string, leaseTTl int64) (*EtcdClient, error) {
////	var addr = make([]string, 1)
////	addr[0] = etcdAddr
////
////	//创建新的etcd客户端
////	client, err := clientv3.New(clientv3.Config{
////		Endpoints:   addr,
////		TLS:         nil,
////		DialTimeout: time.Second * 5,
////	})
////	if err != nil {
////		logs.Info("new clientv3 client err", err)
////		return nil, err
////	}
////
////	//创建kv
////	kv := clientv3.NewKV(client)
////	ctx := context.Background()
////
////	//创建租约并获得租约id
////	lease := clientv3.NewLease(client)
////	leaseResp, err := lease.Grant(ctx, leaseTTl)
////	if err != nil {
////		logs.Info("grant lease err", err)
////		return nil, err
////	}
////	leaseId := leaseResp.ID
////
////	//设置k-v的键值
////	//  schema:///serviceName/ip:port ->ip:port
////	serviceValue := net.JoinHostPort(serverIp, port)               //格式=> ip:port
////	serviceKey := GetKeyPrefix(schema, serviceName) + serviceValue //格式=> fm:///FMlogin/ip:port
////
////	//将kv存储进去
////	_, err = kv.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(leaseId))
////	if err != nil {
////		logs.Error("clientv3 put err ", err)
////		return nil, err
////	}
////
////	//设置keepalive
////	kresp, err := client.KeepAlive(ctx, leaseId)
////	if err != nil {
////		logs.Error("keepAlive err", err)
////		return nil, err
////	}
////
////	go func() {
////	Flood:
////		for {
////			select {
////			case _, ok := <-kresp:
////				if ok == true {
////				} else {
////					break Flood
////				}
////			}
////		}
////	}()
////
////	etcdClient := &EtcdClient{
////		address:  addr,
////		leaseTTl: leaseTTl,
////		kv:       kv,
////		client:   client,
////		ctx:      ctx,
////		leaseId:  leaseId,
////	}
////	return etcdClient, nil
////}
//
//// 格式化字符串为:/grpc/serviceName/ip:port/模式
//func GetIpPrefix(serviceName string) string {
//	return fmt.Sprintf("/grpc/%s/", serviceName)
//}
//
////value值为ip及端口
////传入地址，服务ip，端口号，注册的服务名称，租约过期时间
////值为ip形式
////func RegistEtcdClintByip(etcdAddr string, serverIp, port, serviceName string, leaseTTl int64) (*EtcdClient, error) {
////	var addr = make([]string, 1)
////	addr[0] = etcdAddr
////	fmt.Println("RegistEtcdClintByip -- 1")
////
////	//创建新的etcd客户端
////	client, err := clientv3.New(clientv3.Config{
////		Endpoints:   addr,
////		TLS:         nil,
////		DialTimeout: time.Second * 5,
////		Username:    "",
////		Password:    "",
////	})
////	fmt.Println("RegistEtcdClintByip -- 2")
////
////	if err != nil {
////		logs.Info("new clientv3 client err", err)
////		return nil, err
////	}
////
////	//创建kv
////	kv := clientv3.NewKV(client)
////
////	fmt.Println("RegistEtcdClintByip -- 3")
////
////	//创建租约并获得租约id
////	//lease := clientv3.NewLease(client)
////	fmt.Println("RegistEtcdClintByip -- 4")
////	ctx, cancel := context.WithCancel(context.Background())
////	defer cancel()
////	leaseResp, err := client.Grant(ctx, leaseTTl)
////
////	fmt.Println("RegistEtcdClintByip -- 5")
////
////	if err != nil {
////		logs.Info("grant lease err", err)
////		return nil, err
////	}
////	leaseId := leaseResp.ID
////
////	//设置k-v的键值
////	//  schema:///serviceName/ip:port ->ip:port
////	serviceValue := net.JoinHostPort(serverIp, port) //格式=> ip:port
////	serviceKey := GetIpPrefix(serviceName)           //格式=> "/grpc/serviceName/
////	//prefix := "/test/loginName/"
////	fmt.Println("--prefix-register-key", serviceKey)
////	fmt.Println("-prefix-register-value", serviceValue)
////	//将kv存储进去
////	_, err = kv.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(leaseId))
////	if err != nil {
////		logs.Error("clientv3 put err ", err)
////		return nil, err
////	}
////	resp, _ := kv.Get(ctx, serviceKey, clientv3.WithPrefix())
////	servers := make([]string, 0)
////	for _, value := range resp.Kvs {
////		if value != nil {
////			servers = append(servers, string(value.Value))
////		}
////	}
////	fmt.Println("register get servers", servers)
////	//设置keepalive
////	kresp, err := client.KeepAlive(ctx, leaseId)
////	if err != nil {
////		logs.Error("keepAlive err", err)
////		return nil, err
////	}
////
////	go func() {
////	Flood:
////		for {
////			select {
////			case _, ok := <-kresp:
////				if ok == true {
////				} else {
////					break Flood
////				}
////			}
////		}
////	}()
////
////	etcdClient := &EtcdClient{
////		address:  addr,
////		leaseTTl: leaseTTl,
////		kv:       kv,
////		client:   client,
////		ctx:      ctx,
////		leaseId:  leaseId,
////	}
////
////	resp2, _ := kv.Get(ctx, serviceKey, clientv3.WithPrefix())
////	servers2 := make([]string, 0)
////	for _, value := range resp2.Kvs {
////		if value != nil {
////			servers = append(servers, string(value.Value))
////		}
////	}
////	fmt.Println("register get servers2", servers2)
////	return etcdClient, nil
////}
//
////创建新的client
//func NewEtcdClientnew(leaseTTL int64) *EtcdClient {
//	var address = make([]string, 1)
//	address[0] = config.EtcdAddress
//	var client = &EtcdClient{
//		ctx:      context.Background(),
//		address:  address,
//		leaseTTL: leaseTTL,
//	}
//
//	err := client.connect()
//	if err != nil {
//		panic(err)
//	}
//	return client
//}
//
//func (e *EtcdClient) connect() (err error) {
//	e.client, err = clientv3.New(clientv3.Config{
//		Endpoints:   e.address,
//		DialTimeout: 5 * time.Second,
//		TLS:         nil,
//		Username:    e.username,
//		Password:    e.password,
//	})
//	if err != nil {
//		return
//	}
//	e.kv = clientv3.NewKV(e.client)
//	e.ctx = context.Background()
//	return
//}
//
//func (e *EtcdClient) Close() (err error) {
//	return e.client.Close()
//}
//
//func (e *EtcdClient) register(addr string) (*clientv3.PutResponse, error) {
//	e.lease = clientv3.NewLease(e.client)
//	leaseResp, err := e.lease.Grant(e.ctx, e.leaseTTL)
//	if err != nil {
//		return nil, err
//	}
//
//	e.leaseID = leaseResp.ID
//	prefix := GetIpPrefix(config.LoginName)
//	return e.kv.Put(e.ctx, prefix, addr, clientv3.WithLease(e.leaseID))
//}
//
//func (e *EtcdClient) LeaseKeepAlive() error {
//	addr := utils.ServerIP + ":" + config.LoginName
//	if e.lease == nil {
//		_, err := e.register(addr)
//		if err != nil {
//			return err
//		}
//	}
//	_, err := e.lease.KeepAlive(e.ctx, e.leaseID)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func HealthCheck(provider *HealthProvider) {
//	var tick = time.NewTicker(time.Second)
//	for {
//		select {
//		case <-tick.C:
//			err := provider.EtcdClient.LeaseKeepAlive()
//			if err != nil {
//				fmt.Println(err.Error())
//				return
//			}
//		}
//	}
//}
