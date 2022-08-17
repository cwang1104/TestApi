package rpcLogin

import (
	"TestApi/common/config"
	"TestApi/common/grpc-etcdv3/getcdv3"
	"TestApi/common/utils"
	pbLogin "TestApi/proto/login"
	"fmt"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"net"
)

type rpcLogin struct {
	rpcPort         string
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        string //todo:更换为[]string

	pbLogin.UnimplementedUserLoginServer
}

func NewRpcLoginServer() *rpcLogin {
	return &rpcLogin{
		rpcPort:         config.LoginPort,
		rpcRegisterName: config.LoginName,
		etcdAddr:        config.EtcdAddress,
		etcdSchema:      config.EtcdSchema,
	}
}

// run rpc server
func (rpc *rpcLogin) Run() {

	address := utils.ServerIP + ":" + rpc.rpcPort
	fmt.Println("rpcLogin run address", address)
	fmt.Println("1")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("listen rpc network failed : ", err)
		return
	}
	fmt.Println("2")
	server := grpc.NewServer()
	defer server.GracefulStop()

	//todo:注册到etcd

	pbLogin.RegisterUserLoginServer(server, rpc)

	fmt.Println("3")
	//传入etcd名，地址，服务ip，端口号，租约过期时间
	//_, err := getcdv3.NewEtcdClint(rpc.etcdSchema, rpc.etcdAddr, config.ServerIp, rpc.rpcPort, 10)
	//_, err = getcdv3.RegistEtcdClintByip(config.EtcdAddress, config.ServerIp, rpc.rpcPort, rpc.rpcRegisterName, int64(10))
	getcdv3.RegisterRpc(rpc.rpcRegisterName, rpc.rpcPort, int64(10))
	//if err != nil {
	//	logs.Error("getConn err", err)
	//	return
	//}

	fmt.Println("4")

	err = server.Serve(lis)
	if err != nil {
		fmt.Println("serve server failed err = ", err)
		return
	}
	fmt.Println("--rpc init success")
	logs.Info("rpc init success")
}
