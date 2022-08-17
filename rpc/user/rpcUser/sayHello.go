package rpcUser

import (
	pbUser "TestApi/proto/user"
	"context"
	"fmt"
)

func (s *SayHelloServer) SayHelloToWho(_ context.Context, pb *pbUser.HelloRequest) (*pbUser.HelloResponse, error) {
	resp := pbUser.HelloResponse{
		Resp: fmt.Sprintf("hello，i am %s, i am %d years old", pb.Name, pb.Age),
	}
	return &resp, nil
}
