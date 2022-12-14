// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/login/login.proto

package pbLogin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserLoginClient is the client API for UserLogin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserLoginClient interface {
	AccountRegister(ctx context.Context, in *AccountRegisterReq, opts ...grpc.CallOption) (*AccountRegisterResp, error)
	ModifyUserInfo(ctx context.Context, in *ModifyUserInfoReq, opts ...grpc.CallOption) (*ModifyUserInfoResp, error)
	GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*AccountRegisterResp, error)
}

type userLoginClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLoginClient(cc grpc.ClientConnInterface) UserLoginClient {
	return &userLoginClient{cc}
}

func (c *userLoginClient) AccountRegister(ctx context.Context, in *AccountRegisterReq, opts ...grpc.CallOption) (*AccountRegisterResp, error) {
	out := new(AccountRegisterResp)
	err := c.cc.Invoke(ctx, "/pblogin.UserLogin/AccountRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginClient) ModifyUserInfo(ctx context.Context, in *ModifyUserInfoReq, opts ...grpc.CallOption) (*ModifyUserInfoResp, error) {
	out := new(ModifyUserInfoResp)
	err := c.cc.Invoke(ctx, "/pblogin.UserLogin/ModifyUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginClient) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*AccountRegisterResp, error) {
	out := new(AccountRegisterResp)
	err := c.cc.Invoke(ctx, "/pblogin.UserLogin/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLoginServer is the server API for UserLogin service.
// All implementations must embed UnimplementedUserLoginServer
// for forward compatibility
type UserLoginServer interface {
	AccountRegister(context.Context, *AccountRegisterReq) (*AccountRegisterResp, error)
	ModifyUserInfo(context.Context, *ModifyUserInfoReq) (*ModifyUserInfoResp, error)
	GetUserInfo(context.Context, *UserInfoReq) (*AccountRegisterResp, error)
	mustEmbedUnimplementedUserLoginServer()
}

// UnimplementedUserLoginServer must be embedded to have forward compatible implementations.
type UnimplementedUserLoginServer struct {
}

func (UnimplementedUserLoginServer) AccountRegister(context.Context, *AccountRegisterReq) (*AccountRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountRegister not implemented")
}
func (UnimplementedUserLoginServer) ModifyUserInfo(context.Context, *ModifyUserInfoReq) (*ModifyUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserInfo not implemented")
}
func (UnimplementedUserLoginServer) GetUserInfo(context.Context, *UserInfoReq) (*AccountRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserLoginServer) mustEmbedUnimplementedUserLoginServer() {}

// UnsafeUserLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLoginServer will
// result in compilation errors.
type UnsafeUserLoginServer interface {
	mustEmbedUnimplementedUserLoginServer()
}

func RegisterUserLoginServer(s grpc.ServiceRegistrar, srv UserLoginServer) {
	s.RegisterService(&UserLogin_ServiceDesc, srv)
}

func _UserLogin_AccountRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServer).AccountRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pblogin.UserLogin/AccountRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServer).AccountRegister(ctx, req.(*AccountRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLogin_ModifyUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServer).ModifyUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pblogin.UserLogin/ModifyUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServer).ModifyUserInfo(ctx, req.(*ModifyUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLogin_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pblogin.UserLogin/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServer).GetUserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserLogin_ServiceDesc is the grpc.ServiceDesc for UserLogin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserLogin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pblogin.UserLogin",
	HandlerType: (*UserLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountRegister",
			Handler:    _UserLogin_AccountRegister_Handler,
		},
		{
			MethodName: "ModifyUserInfo",
			Handler:    _UserLogin_ModifyUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserLogin_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/login/login.proto",
}
