package main

import (
	"TestApi/rpc/user/rpcUser"
	"fmt"
)

func main() {
	rpcServer := rpcUser.NewSayHelloServer()
	rpcServer.Run()
	fmt.Println("rpc run sucess")
}
