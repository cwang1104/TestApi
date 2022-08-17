package getcdv3

//
//import (
//	"context"
//	"fmt"
//	"github.com/astaxie/beego/logs"
//	clientv3 "go.etcd.io/etcd/client/v3"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"time"
//)
//
//func GetConn(etcdAddr, serviceName string) *grpc.ClientConn {
//	etcdClient := NewEtcdClient(etcdAddr)
//	err := etcdClient.Connect()
//	if err != nil {
//		logs.Error(err)
//		return nil
//	}
//	defer etcdClient.Close()
//	addresses, err := etcdClient.List(serviceName)
//	if err != nil {
//		logs.Error("list address err", err)
//		return nil
//	}
//	fmt.Println("---getconn addresses", addresses)
//	address := addresses[0]
//	fmt.Println("getconn", address)
//	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		logs.Error("conn get failed,err", err)
//		return nil
//	}
//	return conn
//}
//
//func ClientNewEtcdClient(etcdAddr string) *EtcdClient {
//	var addr = make([]string, 1)
//	addr[0] = etcdAddr
//	var client = &EtcdClient{
//		ctx:     context.Background(),
//		address: addr,
//	}
//	err := client.Connect()
//	if err != nil {
//		logs.Error("etcdv3 connect err", err)
//		return nil
//	}
//	return client
//}
//
//func (e *EtcdClient) Connect() (err error) {
//	e.client, err = clientv3.New(clientv3.Config{
//		Endpoints:   e.address,
//		DialTimeout: 5 * time.Second,
//		TLS:         nil,
//		Username:    "",
//		Password:    "",
//	})
//	if err != nil {
//		return
//	}
//	e.kv = clientv3.NewKV(e.client)
//	e.ctx = context.Background()
//	return
//}
//
//func (e *EtcdClient) List(serviceName string) ([]string, error) {
//	prefix := GetIpPrefix(serviceName)
//	fmt.Println("----list - serviceName", serviceName)
//	fmt.Println("----list", prefix)
//	//prefix := "/test/loginName/"
//	resp, err := e.client.Get(e.ctx, prefix, clientv3.WithPrefix())
//	if err != nil {
//		logs.Error("list get value err", err)
//		return nil, err
//	}
//	servers := make([]string, 0)
//	for _, value := range resp.Kvs {
//		if value != nil {
//			servers = append(servers, string(value.Value))
//		}
//	}
//	fmt.Println("---", servers)
//	return servers, nil
//}
