// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// FollowersClient is the client API for Followers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowersClient interface {
	CreateUserConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetAllFollowers(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type followersClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowersClient(cc grpc.ClientConnInterface) FollowersClient {
	return &followersClient{cc}
}

func (c *followersClient) CreateUserConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Followers/CreateUserConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetAllFollowers(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/proto.Followers/GetAllFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowersServer is the server API for Followers service.
// All implementations must embed UnimplementedFollowersServer
// for forward compatibility
type FollowersServer interface {
	CreateUserConnection(context.Context, *CreateFollowerRequest) (*EmptyResponse, error)
	GetAllFollowers(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	mustEmbedUnimplementedFollowersServer()
}

// UnimplementedFollowersServer must be embedded to have forward compatible implementations.
type UnimplementedFollowersServer struct {
}

func (UnimplementedFollowersServer) CreateUserConnection(context.Context, *CreateFollowerRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserConnection not implemented")
}
func (UnimplementedFollowersServer) GetAllFollowers(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFollowers not implemented")
}
func (UnimplementedFollowersServer) mustEmbedUnimplementedFollowersServer() {}

// UnsafeFollowersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowersServer will
// result in compilation errors.
type UnsafeFollowersServer interface {
	mustEmbedUnimplementedFollowersServer()
}

func RegisterFollowersServer(s grpc.ServiceRegistrar, srv FollowersServer) {
	s.RegisterService(&Followers_ServiceDesc, srv)
}

func _Followers_CreateUserConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).CreateUserConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/CreateUserConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).CreateUserConnection(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetAllFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetAllFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/GetAllFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetAllFollowers(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Followers_ServiceDesc is the grpc.ServiceDesc for Followers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Followers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Followers",
	HandlerType: (*FollowersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserConnection",
			Handler:    _Followers_CreateUserConnection_Handler,
		},
		{
			MethodName: "GetAllFollowers",
			Handler:    _Followers_GetAllFollowers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "followers.proto",
}
