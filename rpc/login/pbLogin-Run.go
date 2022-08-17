package main

import (
	rpcLogin "TestApi/rpc/login/rpcLogin"
	"fmt"
)

func main() {
	//rpcPort := flag.String("port", config.LoginPort, "RpcLogin default listen port 50101")
	//flag.Parse()
	rpcServer := rpcLogin.NewRpcLoginServer()
	rpcServer.Run()
	fmt.Println("rpc run sucess")

}
