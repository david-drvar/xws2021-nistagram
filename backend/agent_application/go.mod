module github.com/david-drvar/xws2021-nistagram/agent_application

go 1.16

replace github.com/david-drvar/xws2021-nistagram/common => ./../common

require (
	github.com/david-drvar/xws2021-nistagram/common v0.0.1
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	google.golang.org/genproto v0.0.0-20210524171403-669157292da3 // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/gorm v1.21.10
)
