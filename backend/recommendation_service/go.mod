module github.com/david-drvar/xws2021-nistagram/recommendation_service

go 1.16

replace github.com/david-drvar/xws2021-nistagram/common => ./../common

require (
	github.com/david-drvar/xws2021-nistagram/common v0.0.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/neo4j/neo4j-go-driver/v4 v4.3.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/sirupsen/logrus v1.8.1 // indirect
	google.golang.org/genproto v0.0.0-20210524171403-669157292da3 // indirect
	google.golang.org/grpc v1.38.0
	gorm.io/gorm v1.21.10 // indirect

)
