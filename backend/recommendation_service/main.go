package main

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/util/setup"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"os"
	"time"
)

func main(){
	if os.Getenv("Docker_env") == "" {
		SetupEnvVariables()
	}

	time.Sleep(30*time.Second)

	driver, _ := setup.CreateConnection(os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_PW"))
	defer setup.CloseConnection(driver)
	for  i := 1; i < 5; i++ {
		duration := 2 << i - 1
		time.Sleep(time.Duration(duration)*time.Second)
		err := CreateUniqueConstraint(driver)
		if err != nil {
			log.Println("Retrying to connect to Neo4j...")
			continue
		}else {
			log.Println("Successfully connected to Neo4j!")
		}
		break
	}

	setup.GRPCServer(driver)
}

func SetupEnvVariables() {
	os.Setenv("DB_HOST", "bolt://localhost:7687")
	os.Setenv("DB_NAME", common.RecommendationDatabaseName)
	os.Setenv("DB_PW", "root")
	os.Setenv("USER_SERVICE", grpc_common.Users_service_address)
	os.Setenv("CONTENT_SERVICE", grpc_common.Content_service_address)
	os.Setenv("CHAT_SERVICE", grpc_common.Chat_service_address)
	os.Setenv("AGENT_APPLICATION", grpc_common.Agent_service_address)
	os.Setenv("REDIS_HOST", "localhost:6379")
	os.Setenv("REDIS_PW", "")
}

func CreateUniqueConstraint(driver neo4j.Driver) error {
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			"CREATE CONSTRAINT constraint_id IF NOT EXISTS ON (u:User) ASSERT u.id IS UNIQUE",
			map[string]interface{}{

			})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return  err
	}
	return  nil
}