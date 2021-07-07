package grpc_common

import (
	"context"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	otgo "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"log"
)

func CreateGrpcConnection(address string) (*grpc.ClientConn, error) {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	return grpc.DialContext(
		context.Background(),
		address,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(4 << 30),
			grpc.MaxCallRecvMsgSize(4 << 30),
			),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		),
		grpc.WithStreamInterceptor(
			grpc_opentracing.StreamClientInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		),
	)
}

// MUST use defer outside the function, not in it
func GetClientConnection(address string) (*grpc.ClientConn, error){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s ", err)
	}

	return conn, err
}

func GetFollowersClient(conn *grpc.ClientConn) protopb.FollowersClient{
	return protopb.NewFollowersClient(conn)
}

func GetUsersClient(conn *grpc.ClientConn) protopb.UsersClient{
	return protopb.NewUsersClient(conn)
}

func GetPrivacyClient(conn *grpc.ClientConn) protopb.PrivacyClient{
	return protopb.NewPrivacyClient(conn)
}

func GetContentClient(conn *grpc.ClientConn) protopb.ContentClient {
	return protopb.NewContentClient(conn)
}