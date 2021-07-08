package main

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/interceptors/rbac"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/content_service/util/setup"
	"gorm.io/gorm"
	"os"
	"time"
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

	customLogger := logger.NewLogger()

	go checkOngoingCampaigns(db, customLogger)
	setup.GRPCServer(db, customLogger)
}

func checkOngoingCampaigns(db *gorm.DB, customLogger *logger.Logger){
	for {
		time.Sleep(1 * time.Minute)
		now := time.Now()
		campaigns := []persistence.Campaign{}
		result := db.Model(&persistence.Campaign{}).Where("start_date <= ? AND end_date >= ?", now, now).Find(&campaigns)
		if result.Error != nil {
			continue
		} else {
			if len(campaigns) > 0 {
				customLogger.ToStdout("CheckOngoingCampaigns", "Found ongoing campaigns", "info")
				model.SetAreCampaignsOngoing(true)
			} else {
				customLogger.ToStdout("CheckOngoingCampaigns", "No ongoing campaigns", "info")
				model.SetAreCampaignsOngoing(false)
			}
		}
	}
}

func SetupEnvVariables() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", common.ContentDatabaseName)
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PW", "root")
	os.Setenv("RECOMMENDATION_SERVICE", grpc_common.Recommendation_service_address)
	os.Setenv("USER_SERVICE", grpc_common.Users_service_address)
	os.Setenv("CHAT_SERVICE", grpc_common.Chat_service_address)
	os.Setenv("AGENT_APPLICATION", grpc_common.Agent_service_address)
}

