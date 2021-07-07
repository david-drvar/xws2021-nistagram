package main

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/interceptors/rbac"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/setup"
	"os"
)

func main() {
	if os.Getenv("Docker_env") == "" {
		SetupEnvVariables()
	}

	db := common.InitDatabase(common.UserDatabase)

	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}

	err = rbac.SetupUsersRBAC(db)
	if err != nil {
		panic("Cannot setup rbac tables. Error message: " + err.Error())
	}

	setup.GRPCServer(db)
}

func SetupEnvVariables() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", common.UsersDatabaseName)
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PW", "root")
	os.Setenv("RECOMMENDATION_SERVICE", grpc_common.Recommendation_service_address)
	os.Setenv("CONTENT_SERVICE", grpc_common.Content_service_address)
	os.Setenv("CHAT_SERVICE", grpc_common.Chat_service_address)
	os.Setenv("AGENT_APPLICATION", grpc_common.Agent_service_address)
	os.Setenv("REDIS_HOST", "localhost:6379")
	os.Setenv("REDIS_PW", "")
}
