package main

import (
	"github.com/david-drvar/xws2021-nistagram/chat_service/controllers"
	"github.com/david-drvar/xws2021-nistagram/chat_service/util/setup"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"os"
)

func main(){
	if os.Getenv("Docker_env") == "" {
		SetupEnvVariables()
	}
	db := common.InitDatabase(common.ChatDatabase)
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}
	controller , _:= controllers.NewMessageController(db)
	setup.ServerSetup(controller)
}

func SetupEnvVariables() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", common.ChatDatabaseName)
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PW", "root")
	os.Setenv("RECOMMENDATION_SERVICE", grpc_common.Recommendation_service_address)
	os.Setenv("USER_SERVICE", grpc_common.Users_service_address)
	os.Setenv("CONTENT_SERVICE", grpc_common.Content_service_address)
	os.Setenv("AGENT_SERVICE", grpc_common.Agent_service_address)
}