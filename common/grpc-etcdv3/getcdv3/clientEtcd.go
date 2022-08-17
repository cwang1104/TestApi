package getcdv3

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
)

func GetGrpcConn(serviceName string) (*grpc.ClientConn, error) {
	etcdClient := NewEtcdClient(10)

	link, err := GetServer(etcdClient, serviceName)
	if err != nil {
		logs.Debug("get server err", err)
		return nil, err
	}
	conn, err := grpc.Dial(link, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logs.Error("conn get failed,err", err)
		return nil, err
	}
	return conn, nil

}

func (e *EtcdClient) List(prefix string) ([]string, error) {
	resp, err := e.kv.Get(e.ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	servers := make([]string, 0)
	for _, value := range resp.Kvs {
		if value != nil {
			servers = append(servers, string(value.Value))
		}
	}
	return servers, nil
}

func genRand(num int) int {
	return int(rand.Int31n(int32(num)))
}

func GetServer(client *EtcdClient, serviceName string) (string, error) {
	key := GetIpPrefix(serviceName)
	servers, err := client.List(key)
	if err != nil {
		return "", err
	}
	fmt.Println("----------", servers)

	return servers[genRand(len(servers))], nil
}
