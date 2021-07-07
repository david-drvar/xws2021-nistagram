package setup

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/controllers"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/saga"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func GRPCServer(driver neo4j.Driver) {
	customLogger := logger.NewLogger()

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", grpc_common.Recommendation_service_address)
	if err != nil {
		customLogger.ToStdoutAndFile("Recommendation GRPC Server", "Couldn't listen to "+grpc_common.Recommendation_service_address, logger.Fatal)
		return
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	redisServer := saga.NewRedisServer(driver)
	go redisServer.RedisConnection()

	server, err := controllers.NewServer(driver, customLogger)
	if err != nil {
		customLogger.ToStdoutAndFile("Recommendation GRPC Server", "Couldn't create server", logger.Fatal)
		return
	}

	protopb.RegisterFollowersServer(s, server)

	customLogger.ToStdoutAndFile("Recommendation GRPC Server", "Serving gRPC on "+grpc_common.Recommendation_service_address, logger.Info)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Recommendation_service_address)
	if err != nil {
		customLogger.ToStdoutAndFile("Recommendation GRPC Server", "Couldn't connect to "+grpc_common.Recommendation_service_address, logger.Fatal)
		return
	}

	gatewayMux := runtime.NewServeMux()
	// Register Greeter
	err = protopb.RegisterFollowersHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		customLogger.ToStdoutAndFile("Recommendation GRPC Server", "Couldn't register gateway", logger.Fatal)
	}

	c := common.SetupCors()
	pool := x509.NewCertPool()

	// Here is the certificate provided by the loading client, preferably the root certificate provided by the client.
	addTrust(pool, "./../common/sslFile/gateway.p12")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	gwServer := &http.Server{
		Addr:    grpc_common.Recommendation_gateway_address,
		Handler: tracer.TracingWrapper(c.Handler(gatewayMux)),
		/*TLSConfig: &tls.Config{
			ClientCAs: pool,
			ClientAuth:  tls.RequireAndVerifyClientCert,
		},*/
	}

	customLogger.ToStdoutAndFile("Recommendation GRPC Server", "Serving gRPC-Gateway on "+grpc_common.Recommendation_gateway_address, logger.Info)
	//log.Fatalln(gwServer.ListenAndServeTLS("./../common/sslFile/gateway.crt", "./../common/sslFile/gateway.key"))
	log.Fatalln(gwServer.ListenAndServe())
}

func addTrust(pool *x509.CertPool, path string) {
	aCrt, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(aCrt)

}
