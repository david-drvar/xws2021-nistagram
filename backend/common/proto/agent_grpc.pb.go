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

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	//    Products
	CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*EmptyResponseAgent, error)
	GetAllProductsByAgentId(ctx context.Context, in *UserAgentApp, opts ...grpc.CallOption) (*ProductsArray, error)
	GetProductById(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
	GetAllProducts(ctx context.Context, in *EmptyRequestAgent, opts ...grpc.CallOption) (*ProductsArray, error)
	DeleteProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*EmptyResponseAgent, error)
	UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*EmptyResponseAgent, error)
	//    Orders
	OrderProduct(ctx context.Context, in *Order, opts ...grpc.CallOption) (*EmptyResponseAgent, error)
	GetOrdersByUser(ctx context.Context, in *UserAgentApp, opts ...grpc.CallOption) (*OrdersArray, error)
	GetOrdersByAgent(ctx context.Context, in *UserAgentApp, opts ...grpc.CallOption) (*OrdersArray, error)
	//    Users
	LoginUserInAgentApp(ctx context.Context, in *LoginRequestAgentApp, opts ...grpc.CallOption) (*LoginResponseAgentApp, error)
	CreateUserInAgentApp(ctx context.Context, in *CreateUserRequestAgentApp, opts ...grpc.CallOption) (*EmptyResponseAgent, error)
	GetUserByUsername(ctx context.Context, in *RequestUsernameAgent, opts ...grpc.CallOption) (*UserAgentApp, error)
	// API KEY
	GetKeyByUserId(ctx context.Context, in *RequestIdAgent, opts ...grpc.CallOption) (*ApiTokenAgent, error)
	UpdateKey(ctx context.Context, in *ApiTokenAgent, opts ...grpc.CallOption) (*EmptyResponseAgent, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*EmptyResponseAgent, error) {
	out := new(EmptyResponseAgent)
	err := c.cc.Invoke(ctx, "/proto.Agent/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetAllProductsByAgentId(ctx context.Context, in *UserAgentApp, opts ...grpc.CallOption) (*ProductsArray, error) {
	out := new(ProductsArray)
	err := c.cc.Invoke(ctx, "/proto.Agent/GetAllProductsByAgentId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetProductById(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/proto.Agent/GetProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetAllProducts(ctx context.Context, in *EmptyRequestAgent, opts ...grpc.CallOption) (*ProductsArray, error) {
	out := new(ProductsArray)
	err := c.cc.Invoke(ctx, "/proto.Agent/GetAllProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeleteProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*EmptyResponseAgent, error) {
	out := new(EmptyResponseAgent)
	err := c.cc.Invoke(ctx, "/proto.Agent/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*EmptyResponseAgent, error) {
	out := new(EmptyResponseAgent)
	err := c.cc.Invoke(ctx, "/proto.Agent/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) OrderProduct(ctx context.Context, in *Order, opts ...grpc.CallOption) (*EmptyResponseAgent, error) {
	out := new(EmptyResponseAgent)
	err := c.cc.Invoke(ctx, "/proto.Agent/OrderProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetOrdersByUser(ctx context.Context, in *UserAgentApp, opts ...grpc.CallOption) (*OrdersArray, error) {
	out := new(OrdersArray)
	err := c.cc.Invoke(ctx, "/proto.Agent/GetOrdersByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetOrdersByAgent(ctx context.Context, in *UserAgentApp, opts ...grpc.CallOption) (*OrdersArray, error) {
	out := new(OrdersArray)
	err := c.cc.Invoke(ctx, "/proto.Agent/GetOrdersByAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) LoginUserInAgentApp(ctx context.Context, in *LoginRequestAgentApp, opts ...grpc.CallOption) (*LoginResponseAgentApp, error) {
	out := new(LoginResponseAgentApp)
	err := c.cc.Invoke(ctx, "/proto.Agent/LoginUserInAgentApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CreateUserInAgentApp(ctx context.Context, in *CreateUserRequestAgentApp, opts ...grpc.CallOption) (*EmptyResponseAgent, error) {
	out := new(EmptyResponseAgent)
	err := c.cc.Invoke(ctx, "/proto.Agent/CreateUserInAgentApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetUserByUsername(ctx context.Context, in *RequestUsernameAgent, opts ...grpc.CallOption) (*UserAgentApp, error) {
	out := new(UserAgentApp)
	err := c.cc.Invoke(ctx, "/proto.Agent/GetUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetKeyByUserId(ctx context.Context, in *RequestIdAgent, opts ...grpc.CallOption) (*ApiTokenAgent, error) {
	out := new(ApiTokenAgent)
	err := c.cc.Invoke(ctx, "/proto.Agent/GetKeyByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpdateKey(ctx context.Context, in *ApiTokenAgent, opts ...grpc.CallOption) (*EmptyResponseAgent, error) {
	out := new(EmptyResponseAgent)
	err := c.cc.Invoke(ctx, "/proto.Agent/UpdateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	//    Products
	CreateProduct(context.Context, *Product) (*EmptyResponseAgent, error)
	GetAllProductsByAgentId(context.Context, *UserAgentApp) (*ProductsArray, error)
	GetProductById(context.Context, *Product) (*Product, error)
	GetAllProducts(context.Context, *EmptyRequestAgent) (*ProductsArray, error)
	DeleteProduct(context.Context, *Product) (*EmptyResponseAgent, error)
	UpdateProduct(context.Context, *Product) (*EmptyResponseAgent, error)
	//    Orders
	OrderProduct(context.Context, *Order) (*EmptyResponseAgent, error)
	GetOrdersByUser(context.Context, *UserAgentApp) (*OrdersArray, error)
	GetOrdersByAgent(context.Context, *UserAgentApp) (*OrdersArray, error)
	//    Users
	LoginUserInAgentApp(context.Context, *LoginRequestAgentApp) (*LoginResponseAgentApp, error)
	CreateUserInAgentApp(context.Context, *CreateUserRequestAgentApp) (*EmptyResponseAgent, error)
	GetUserByUsername(context.Context, *RequestUsernameAgent) (*UserAgentApp, error)
	// API KEY
	GetKeyByUserId(context.Context, *RequestIdAgent) (*ApiTokenAgent, error)
	UpdateKey(context.Context, *ApiTokenAgent) (*EmptyResponseAgent, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) CreateProduct(context.Context, *Product) (*EmptyResponseAgent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedAgentServer) GetAllProductsByAgentId(context.Context, *UserAgentApp) (*ProductsArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProductsByAgentId not implemented")
}
func (UnimplementedAgentServer) GetProductById(context.Context, *Product) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedAgentServer) GetAllProducts(context.Context, *EmptyRequestAgent) (*ProductsArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedAgentServer) DeleteProduct(context.Context, *Product) (*EmptyResponseAgent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedAgentServer) UpdateProduct(context.Context, *Product) (*EmptyResponseAgent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedAgentServer) OrderProduct(context.Context, *Order) (*EmptyResponseAgent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderProduct not implemented")
}
func (UnimplementedAgentServer) GetOrdersByUser(context.Context, *UserAgentApp) (*OrdersArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersByUser not implemented")
}
func (UnimplementedAgentServer) GetOrdersByAgent(context.Context, *UserAgentApp) (*OrdersArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersByAgent not implemented")
}
func (UnimplementedAgentServer) LoginUserInAgentApp(context.Context, *LoginRequestAgentApp) (*LoginResponseAgentApp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUserInAgentApp not implemented")
}
func (UnimplementedAgentServer) CreateUserInAgentApp(context.Context, *CreateUserRequestAgentApp) (*EmptyResponseAgent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserInAgentApp not implemented")
}
func (UnimplementedAgentServer) GetUserByUsername(context.Context, *RequestUsernameAgent) (*UserAgentApp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (UnimplementedAgentServer) GetKeyByUserId(context.Context, *RequestIdAgent) (*ApiTokenAgent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeyByUserId not implemented")
}
func (UnimplementedAgentServer) UpdateKey(context.Context, *ApiTokenAgent) (*EmptyResponseAgent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKey not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetAllProductsByAgentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAgentApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetAllProductsByAgentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/GetAllProductsByAgentId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetAllProductsByAgentId(ctx, req.(*UserAgentApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/GetProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetProductById(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequestAgent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/GetAllProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetAllProducts(ctx, req.(*EmptyRequestAgent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeleteProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpdateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_OrderProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).OrderProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/OrderProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).OrderProduct(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetOrdersByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAgentApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetOrdersByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/GetOrdersByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetOrdersByUser(ctx, req.(*UserAgentApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetOrdersByAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAgentApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetOrdersByAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/GetOrdersByAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetOrdersByAgent(ctx, req.(*UserAgentApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_LoginUserInAgentApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequestAgentApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).LoginUserInAgentApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/LoginUserInAgentApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).LoginUserInAgentApp(ctx, req.(*LoginRequestAgentApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CreateUserInAgentApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequestAgentApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateUserInAgentApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/CreateUserInAgentApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateUserInAgentApp(ctx, req.(*CreateUserRequestAgentApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUsernameAgent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/GetUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetUserByUsername(ctx, req.(*RequestUsernameAgent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetKeyByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIdAgent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetKeyByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/GetKeyByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetKeyByUserId(ctx, req.(*RequestIdAgent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpdateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiTokenAgent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpdateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Agent/UpdateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpdateKey(ctx, req.(*ApiTokenAgent))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _Agent_CreateProduct_Handler,
		},
		{
			MethodName: "GetAllProductsByAgentId",
			Handler:    _Agent_GetAllProductsByAgentId_Handler,
		},
		{
			MethodName: "GetProductById",
			Handler:    _Agent_GetProductById_Handler,
		},
		{
			MethodName: "GetAllProducts",
			Handler:    _Agent_GetAllProducts_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Agent_DeleteProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Agent_UpdateProduct_Handler,
		},
		{
			MethodName: "OrderProduct",
			Handler:    _Agent_OrderProduct_Handler,
		},
		{
			MethodName: "GetOrdersByUser",
			Handler:    _Agent_GetOrdersByUser_Handler,
		},
		{
			MethodName: "GetOrdersByAgent",
			Handler:    _Agent_GetOrdersByAgent_Handler,
		},
		{
			MethodName: "LoginUserInAgentApp",
			Handler:    _Agent_LoginUserInAgentApp_Handler,
		},
		{
			MethodName: "CreateUserInAgentApp",
			Handler:    _Agent_CreateUserInAgentApp_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _Agent_GetUserByUsername_Handler,
		},
		{
			MethodName: "GetKeyByUserId",
			Handler:    _Agent_GetKeyByUserId_Handler,
		},
		{
			MethodName: "UpdateKey",
			Handler:    _Agent_UpdateKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
