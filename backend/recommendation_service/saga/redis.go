package saga

import (
	"context"
	"encoding/json"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/model"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/services"
	"github.com/go-redis/redis"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

type redisServer struct {
	followersService *services.FollowersService
}

func NewRedisServer(driver neo4j.Driver) *redisServer {
	followersService, _ := services.NewFollowersService(driver)
	return &redisServer{
		followersService: followersService,
	}
}

func (os *redisServer) RedisConnection() {
	// create client and ping redis
	var err error
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	if _, err = client.Ping().Result(); err != nil {
		log.Fatalf("error creating redis client %s", err)
	}

	// subscribe to the required channels
	pubsub := client.Subscribe(RecommendationChannel, ReplyChannel)
	if _, err = pubsub.Receive(); err != nil {
		log.Fatalf("error subscribing %s", err)
	}
	defer func() { _ = pubsub.Close() }()
	ch := pubsub.Channel()

	log.Println("starting the stock service")
	for {
		select {
		case msg := <-ch:
			m := Message{}
			err := json.Unmarshal([]byte(msg.Payload), &m)
			if err != nil {
				log.Println(err)
				continue
			}

			switch msg.Channel {
			case RecommendationChannel:

				// Happy Flow
				if m.Action == ActionStart {
					//todo upisi u bazu novog usera
					print("AAA from RECOMMENDATION SERVER")
					_, err := os.followersService.CreateUser(context.Background(), model.User{UserId: m.UserId})
					if err != nil {
						sendToReplyChannel(client, m, ActionError, ServiceUser, ServiceRecommendation)
					}
					//sendToReplyChannel(client, m, ActionError, ServiceUser, ServiceRecommendation)
					sendToReplyChannel(client, m, ActionDone, ServiceUser, ServiceRecommendation)
				}

				// Rollback flow
				if m.Action == ActionRollback {
				}

			}
		}
	}
}

func sendToReplyChannel(client *redis.Client, m Message, action string, service string, senderService string) {
	var err error
	m.Action = action
	m.Service = service
	m.SenderService = senderService
	if err = client.Publish(ReplyChannel, m).Err(); err != nil {
		log.Printf("error publishing done-message to %s channel", ReplyChannel)
	}
	log.Printf("done message published to channel :%s", ReplyChannel)
}
