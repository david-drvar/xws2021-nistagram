package setup

import (
	"github.com/david-drvar/xws2021-nistagram/chat_service/model"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Message{},
		&model.MessageRequest{},
		&model.ChatRoom{},
	)

	return err
}
