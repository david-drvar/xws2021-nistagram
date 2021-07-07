package setup

import (
	"context"
	"crypto/tls"
	"github.com/david-drvar/xws2021-nistagram/agent_application/controllers"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/interceptors"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"time"
)

func GRPCServer(db *gorm.DB) {
	customLogger := logger.NewLogger()

	lis, err := net.Listen("tcp", grpc_common.Agent_service_address)
	if err != nil {
		customLogger.ToStdoutAndFile("Agent GRPC Server", "Couldn't listen to "+grpc_common.Agent_service_address, logger.Fatal)
		return
	}

	jwtManager := common.NewJWTManager("somesecretkey", 60*time.Minute)
	rbacInterceptor := interceptors.NewRBACInterceptor(db, jwtManager, customLogger)

	// Create a gRPC server object
	s := grpc.NewServer(
		grpc.UnaryInterceptor(rbacInterceptor.Authorize()),
		grpc.MaxSendMsgSize(4<<30), // Default: 1024 * 1024 * 4 = 4MB -> Override to 4GBs
		grpc.MaxRecvMsgSize(4<<30), // Default: 1024 * 1024 * 4 = 4MB -> Override to 4GBs
	)

	server, err := controllers.NewServer(db, jwtManager, customLogger)
	if err != nil {
		customLogger.ToStdoutAndFile("Agent GRPC Server", "Couldn't create server", logger.Fatal)
		return
	}

	protopb.RegisterAgentServer(s, server)

	customLogger.ToStdoutAndFile("Agent GRPC Server", "Serving gRPC on "+grpc_common.Agent_service_address, logger.Info)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Agent_service_address)
	if err != nil {
		customLogger.ToStdoutAndFile("Agent GRPC Server", "Couldn't connect to "+grpc_common.Agent_service_address, logger.Fatal)
		return
	}

	gatewayMux := runtime.NewServeMux()
	// Register Greeter
	err = protopb.RegisterAgentHandler(context.Background(), gatewayMux, conn) //todo change
	if err != nil {
		customLogger.ToStdoutAndFile("Agent GRPC Server", "Failed to register gateway", logger.Fatal)
		return
	}

	c := common.SetupCors()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	gwServer := &http.Server{
		Addr:    grpc_common.Agent_gateway_address,
		Handler: tracer.TracingWrapper(c.Handler(gatewayMux)),
	}

	customLogger.ToStdoutAndFile("Agent GRPC Server", "Serving gRPC-Gateway on "+grpc_common.Agent_gateway_address, logger.Info)
	// log.Fatalln(gwServer.ListenAndServeTLS("./../common/sslFile/gateway.crt", "./../common/sslFile/gateway.key"))
	log.Fatalln(gwServer.ListenAndServe())
}
