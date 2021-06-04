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
const _ = grpc.SupportPackageIsVersion7

// ContentClient is the client API for Content service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentClient interface {
	CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetAllPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ReducedPostArray, error)
	CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetCommentsForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*CommentsArray, error)
	CreateLike(ctx context.Context, in *Like, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetLikesForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*LikesArray, error)
	GetDislikesForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*LikesArray, error)
	SearchContentByLocation(ctx context.Context, in *SearchLocationRequest, opts ...grpc.CallOption) (*ReducedPostArray, error)
}

type contentClient struct {
	cc grpc.ClientConnInterface
}

func NewContentClient(cc grpc.ClientConnInterface) ContentClient {
	return &contentClient{cc}
}

func (c *contentClient) CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/protoe.Content/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetAllPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ReducedPostArray, error) {
	out := new(ReducedPostArray)
	err := c.cc.Invoke(ctx, "/protoe.Content/GetAllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/protoe.Content/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetCommentsForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*CommentsArray, error) {
	out := new(CommentsArray)
	err := c.cc.Invoke(ctx, "/protoe.Content/GetCommentsForPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateLike(ctx context.Context, in *Like, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/protoe.Content/CreateLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetLikesForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*LikesArray, error) {
	out := new(LikesArray)
	err := c.cc.Invoke(ctx, "/protoe.Content/GetLikesForPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetDislikesForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*LikesArray, error) {
	out := new(LikesArray)
	err := c.cc.Invoke(ctx, "/protoe.Content/GetDislikesForPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) SearchContentByLocation(ctx context.Context, in *SearchLocationRequest, opts ...grpc.CallOption) (*ReducedPostArray, error) {
	out := new(ReducedPostArray)
	err := c.cc.Invoke(ctx, "/protoe.Content/SearchContentByLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServer is the server API for Content service.
// All implementations must embed UnimplementedContentServer
// for forward compatibility
type ContentServer interface {
	CreatePost(context.Context, *Post) (*EmptyResponse, error)
	GetAllPosts(context.Context, *EmptyRequest) (*ReducedPostArray, error)
	CreateComment(context.Context, *Comment) (*EmptyResponse, error)
	GetCommentsForPost(context.Context, *RequestId) (*CommentsArray, error)
	CreateLike(context.Context, *Like) (*EmptyResponse, error)
	GetLikesForPost(context.Context, *RequestId) (*LikesArray, error)
	GetDislikesForPost(context.Context, *RequestId) (*LikesArray, error)
	SearchContentByLocation(context.Context, *SearchLocationRequest) (*ReducedPostArray, error)
	mustEmbedUnimplementedContentServer()
}

// UnimplementedContentServer must be embedded to have forward compatible implementations.
type UnimplementedContentServer struct {
}

func (UnimplementedContentServer) CreatePost(context.Context, *Post) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedContentServer) GetAllPosts(context.Context, *EmptyRequest) (*ReducedPostArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPosts not implemented")
}
func (UnimplementedContentServer) CreateComment(context.Context, *Comment) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedContentServer) GetCommentsForPost(context.Context, *RequestId) (*CommentsArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsForPost not implemented")
}
func (UnimplementedContentServer) CreateLike(context.Context, *Like) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLike not implemented")
}
func (UnimplementedContentServer) GetLikesForPost(context.Context, *RequestId) (*LikesArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikesForPost not implemented")
}
func (UnimplementedContentServer) GetDislikesForPost(context.Context, *RequestId) (*LikesArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDislikesForPost not implemented")
}
func (UnimplementedContentServer) SearchContentByLocation(context.Context, *SearchLocationRequest) (*ReducedPostArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchContentByLocation not implemented")
}
func (UnimplementedContentServer) mustEmbedUnimplementedContentServer() {}

// UnsafeContentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServer will
// result in compilation errors.
type UnsafeContentServer interface {
	mustEmbedUnimplementedContentServer()
}

func RegisterContentServer(s *grpc.Server, srv ContentServer) {
	s.RegisterService(&_Content_serviceDesc, srv)
}

func _Content_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoe.Content/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetAllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetAllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoe.Content/GetAllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetAllPosts(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoe.Content/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetCommentsForPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetCommentsForPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoe.Content/GetCommentsForPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetCommentsForPost(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoe.Content/CreateLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateLike(ctx, req.(*Like))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetLikesForPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetLikesForPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoe.Content/GetLikesForPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetLikesForPost(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetDislikesForPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetDislikesForPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoe.Content/GetDislikesForPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetDislikesForPost(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_SearchContentByLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).SearchContentByLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoe.Content/SearchContentByLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).SearchContentByLocation(ctx, req.(*SearchLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Content_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoe.Content",
	HandlerType: (*ContentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _Content_CreatePost_Handler,
		},
		{
			MethodName: "GetAllPosts",
			Handler:    _Content_GetAllPosts_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _Content_CreateComment_Handler,
		},
		{
			MethodName: "GetCommentsForPost",
			Handler:    _Content_GetCommentsForPost_Handler,
		},
		{
			MethodName: "CreateLike",
			Handler:    _Content_CreateLike_Handler,
		},
		{
			MethodName: "GetLikesForPost",
			Handler:    _Content_GetLikesForPost_Handler,
		},
		{
			MethodName: "GetDislikesForPost",
			Handler:    _Content_GetDislikesForPost_Handler,
		},
		{
			MethodName: "SearchContentByLocation",
			Handler:    _Content_SearchContentByLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content.proto",
}