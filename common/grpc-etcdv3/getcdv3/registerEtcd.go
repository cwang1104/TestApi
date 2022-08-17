package getcdv3

import (
	"TestApi/common/config"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

type EtcdClient struct {
	address  []string
	kv       clientv3.KV
	client   *clientv3.Client
	ctx      context.Context
	lease    clientv3.Lease
	leaseID  clientv3.LeaseID
	leaseTTL int64
	username string
	password string
}

var (
	healthProvider     *HealthProvider
	healthProviderOnce sync.Once
)

type HealthProvider struct {
	EtcdClient *EtcdClient
}

func RegisterRpc(serviceName, port string, leaseTTl int64) {
	provider := GetHealthProvider(leaseTTl)
	go HealthCheck(provider, serviceName, port)
	//defer provider.EtcdClient.Close()
}

func GetHealthProvider(leaseTTl int64) *HealthProvider {
	healthProviderOnce.Do(func() {
		healthProvider = &HealthProvider{
			EtcdClient: NewEtcdClient(leaseTTl),
		}
	})
	return healthProvider
}

func NewEtcdClient(leaseTTl int64) *EtcdClient {
	addr := make([]string, 1)
	addr[0] = config.EtcdAddress

	var client = &EtcdClient{
		ctx:      context.Background(),
		address:  addr,
		leaseTTL: leaseTTl,
	}

	err := client.connect()
	if err != nil {
		panic(err)
	}
	return client
}

func (e *EtcdClient) connect() (err error) {
	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   e.address,
		DialTimeout: 5 * time.Second,
		TLS:         nil,
		Username:    e.username,
		Password:    e.password,
	})
	if err != nil {
		return
	}
	e.kv = clientv3.NewKV(e.client)
	e.ctx = context.Background()
	return
}

func (e *EtcdClient) register(serviceName string, port string) (*clientv3.PutResponse, error) {
	e.lease = clientv3.NewLease(e.client)
	leaseResp, err := e.lease.Grant(e.ctx, e.leaseTTL)
	if err != nil {
		return nil, err
	}

	e.leaseID = leaseResp.ID
	seviceKey := GetIpPrefix(serviceName)
	serviceValue := config.ServerIp + ":" + port

	return e.kv.Put(e.ctx, seviceKey, serviceValue, clientv3.WithLease(e.leaseID))
}

func (e *EtcdClient) LeaseKeepAlive(serviceName string, port string) error {

	if e.lease == nil {
		_, err := e.register(serviceName, port)
		if err != nil {
			return err
		}
	}
	_, err := e.lease.KeepAlive(e.ctx, e.leaseID)
	if err != nil {
		return err
	}
	return nil
}

func HealthCheck(provider *HealthProvider, serviceName string, port string) {
	var tick = time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			err := provider.EtcdClient.LeaseKeepAlive(serviceName, port)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}

// 格式化字符串为:/grpc/serviceName/ip:port/模式
func GetIpPrefix(serviceName string) string {
	return fmt.Sprintf("/grpc/%s/", serviceName)
}

func (e *EtcdClient) Close() (err error) {
	return e.client.Close()
}
