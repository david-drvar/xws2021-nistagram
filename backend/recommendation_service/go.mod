module github.com/david-drvar/xws2021-nistagram/recommendation_service

go 1.16

replace github.com/david-drvar/xws2021-nistagram/common => ./../common

require (
	github.com/david-drvar/xws2021-nistagram/common v0.0.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/neo4j/neo4j-go-driver/v4 v4.3.0
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/genproto v0.0.0-20210524171403-669157292da3 // indirect
	google.golang.org/grpc v1.38.0
	gorm.io/gorm v1.21.10 // indirect

)
