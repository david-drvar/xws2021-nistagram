package main

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/interceptors/rbac"
	"github.com/david-drvar/xws2021-nistagram/content_service/util/setup"
	"os"
)

func main(){
	if os.Getenv("Docker_env") == "" {
		SetupEnvVariables()
	}

	db := common.InitDatabase(common.ContentDatabase)
	err := setup.FillDatabase(db)
	if err != nil { panic("Cannot setup database tables. Error message: " + err.Error()) }

	err = rbac.SetupContentRBAC(db)
	if err != nil { panic("Cannot setup rbac tables. Error message: " + err.Error()) }

	setup.GRPCServer(db)
}


func SetupEnvVariables() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", common.ContentDatabaseName)
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PW", "root")
	os.Setenv("RECOMMENDATION_SERVICE", grpc_common.Recommendation_service_address)
	os.Setenv("USER_SERVICE", grpc_common.Users_service_address)
}

