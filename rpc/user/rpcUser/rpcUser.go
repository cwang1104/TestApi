package rpcUser

import (
	"TestApi/common/config"
	"TestApi/common/grpc-etcdv3/getcdv3"
	"TestApi/common/utils"
	pbUser "TestApi/proto/user"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"net"
)

type SayHelloServer struct {
	rpcPort string
	rpcName string

	pbUser.UnimplementedSayHelloServer
}

func NewSayHelloServer() *SayHelloServer {
	return &SayHelloServer{
		rpcPort: config.UserPort,
		rpcName: config.UserName,
	}
}

func (s *SayHelloServer) Run() {

	address := utils.ServerIP + ":" + s.rpcPort
	logs.Info("init rpc service: %s", address)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		logs.Debug("user listen err", err)
		return
	}

	server := grpc.NewServer()
	defer server.GracefulStop()
	pbUser.RegisterSayHelloServer(server, s)

	getcdv3.RegisterRpc(s.rpcName, s.rpcPort, int64(10))

	logs.Info("start server %s", s.rpcName)
	err = server.Serve(listen)
	if err != nil {
		logs.Debug("serve err", err)
		return
	}

}
