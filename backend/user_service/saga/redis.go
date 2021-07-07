package saga

import (
	"encoding/json"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"log"
)

type RedisServer struct {
	Orchestrator *Orchestrator
	db           *gorm.DB
}

func NewRedisServer(db *gorm.DB) *RedisServer {
	return &RedisServer{
		Orchestrator: NewOrchestrator(),
		db:           db,
	}
}

func (rs *RedisServer) RedisConnection() {
	// create client and ping redis
	var err error
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	if _, err = client.Ping().Result(); err != nil {
		log.Fatalf("error creating redis client %s", err)
	}

	// subscribe to the required channels
	pubsub := client.Subscribe(UserChannel, ReplyChannel)
	if _, err = pubsub.Receive(); err != nil {
		log.Fatalf("error subscribing %s", err)
	}
	defer func() { _ = pubsub.Close() }()
	ch := pubsub.Channel()

	log.Println("starting user service")
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
			case UserChannel:

				// Happy Flow
				if m.Action == ActionStart {
					//user service nema action start jer on ne mora nista da uradi posle recommendation servisa
				}

				// Rollback flow
				if m.Action == ActionRollback {
					//brises user-a jer je recommendation service puknuo
					print("AAA FROM USER SERVER")
					var user persistence.User
					var userPrivacy persistence.Privacy
					var userAdditionalInfo persistence.UserAdditionalInfo
					rs.db.Where("id = ?", m.UserId).Find(&user)
					rs.db.Where("id = ?", m.UserId).Find(&userAdditionalInfo)
					rs.db.Where("user_id = ?", m.UserId).Find(&userPrivacy)
					rs.db.Delete(&user)
					rs.db.Delete(&userAdditionalInfo)
					rs.db.Delete(&userPrivacy)
				}

			}
		}
	}
}

func sendToReplyChannel(client *redis.Client, m *Message, action string, service string, senderService string) {
	var err error
	m.Action = action
	m.Service = service
	m.SenderService = senderService
	if err = client.Publish(ReplyChannel, m).Err(); err != nil {
		log.Printf("error publishing done-message to %s channel", ReplyChannel)
	}
	log.Printf("done message published to channel :%s", ReplyChannel)
}
