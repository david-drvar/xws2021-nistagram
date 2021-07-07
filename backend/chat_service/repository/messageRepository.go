package repository

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/chat_service/model"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MessageRepository interface {
	SaveMessage(context.Context, model.Message) error
	DeleteMessage(context.Context, string) error
	GetMessagesForChatRoom(context.Context, string) ([]model.Message, error)
	GetChatRoomsForUser(context.Context, string) ([]model.ChatRoom, error)
	CreateChatRoom(context.Context, model.ChatRoom) (*model.ChatRoom, error)
	CreateMessageRequest(context.Context, *model.MessageRequest) (*model.MessageRequest, error)
	AcceptMessageRequest(context.Context, model.MessageRequest) error
	DeclineMessageRequest(context.Context, model.MessageRequest) error
	GetChatRoomByUsers(ctx context.Context, room model.ChatRoom) (*model.ChatRoom, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) (MessageRepository, error) {
	if db == nil {
		panic("MessageRepository not created, gorm.DB is nil")
	}
	return &messageRepository{db: db}, nil
}

func (repo *messageRepository) SaveMessage(ctx context.Context, message model.Message) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	message.DateCreated = time.Now()
	message.Id = uuid.New().String()
	result := repo.db.Create(&message)
	if result.Error != nil {
		return errors.New("Could not save message!")
	}

	return nil
}

func (repo *messageRepository) DeleteMessage(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repo.db.Where("id = ?", id).Delete(&model.Message{})
	if result.Error != nil {
		return errors.New("Could not delete message")
	}

	return nil
}

func (repo *messageRepository) GetChatRoomByUsers(ctx context.Context, room model.ChatRoom) (*model.ChatRoom, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	var checkRoom *model.ChatRoom
	result := repo.db.Where("id = ?", room.Person1+"_"+room.Person2).Find(&checkRoom)
	if result.Error != nil {
		return nil, errors.New("Could not load chatRoom")
	} else if checkRoom.Id == "" {
		result = repo.db.Where("id = ?", room.Person2+"_"+room.Person1).Find(&checkRoom)
		if result.Error != nil {
			return nil, errors.New("Could not load chatRoom")
		} else if checkRoom.Id == "" {
			return &model.ChatRoom{}, nil
		}
	}

	return checkRoom, nil
}

func (repo *messageRepository) CreateChatRoom(ctx context.Context, room model.ChatRoom) (*model.ChatRoom, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var checkRoom *model.ChatRoom
	result := repo.db.Where("id = ?", room.Person1+"_"+room.Person2).Find(&checkRoom)
	if result.Error != nil {
		return nil, errors.New("Could not load chatRoom")
	} else if checkRoom.Id != "" {
		return nil, errors.New("Room already exists")
	}

	result = repo.db.Where("id = ?", room.Person2+"_"+room.Person1).Find(&checkRoom)
	if result.Error != nil {
		return nil, errors.New("Could not load chatRoom")
	} else if checkRoom.Id != "" {
		return nil, errors.New("Room already exists")
	}

	room.Id = room.Person1 + "_" + room.Person2
	result = repo.db.Create(&room)
	if result.Error != nil {
		return nil, errors.New("Could not create chat room!")
	} else if result.RowsAffected == 0 {
		return nil, errors.New("Could not create chat room!")
	}

	return &room, nil

}

func (repo *messageRepository) GetChatRoomsForUser(ctx context.Context, userId string) ([]model.ChatRoom, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetChatRoomsForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var rooms1 []model.ChatRoom
	result := repo.db.Where("person1 = ?", userId).Find(&rooms1)
	if result.Error != nil {
		return nil, errors.New("Could not load rooms for user")
	}
	var rooms2 []model.ChatRoom
	result = repo.db.Where("person2 = ?", userId).Find(&rooms2)
	if result.Error != nil {
		return nil, errors.New("Could not load rooms for user")
	}

	return append(rooms1, rooms2...), nil
}

func (repo *messageRepository) GetMessagesForChatRoom(ctx context.Context, roomId string) ([]model.Message, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetMessagesForChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var messages []model.Message
	result := repo.db.Where("room_id = ?", roomId).Find(&messages)
	if result.Error != nil {
		return nil, errors.New("Could not load messages for room")
	}

	return messages, nil
}

//Message requests

func (repo *messageRepository) CreateMessageRequest(ctx context.Context, messageRequest *model.MessageRequest) (*model.MessageRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateMessageRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repo.db.Create(&messageRequest)
	if result.Error != nil {
		return nil, errors.New("Could not create message request")
	} else if result.RowsAffected == 0 {
		return nil, errors.New("Could not create message request")
	}

	return messageRequest, nil

}

func (repo *messageRepository) AcceptMessageRequest(ctx context.Context, messageRequest model.MessageRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "AcceptMessageRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	messageRequest.IsAccepted = true
	result := repo.db.Where("sender_id = ?", messageRequest.SenderId).Where("receiver_id = ?", messageRequest.ReceiverId).Updates(messageRequest)
	if result.Error != nil {
		return errors.New("Could not accept message request")
	} else if result.RowsAffected == 0 {
		return errors.New("Could not accept message request")
	}

	_, err := repo.CreateChatRoom(ctx, model.ChatRoom{Person1: messageRequest.SenderId, Person2: messageRequest.ReceiverId})
	return err
}

func (repo *messageRepository) DeclineMessageRequest(ctx context.Context, messageRequest model.MessageRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeclineMessageRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repo.db.Where("sender_id = ?", messageRequest.SenderId).Where("receiver_id = ?", messageRequest.ReceiverId).Delete(messageRequest)
	if result.Error != nil {
		return errors.New("Could not accept message request")
	} else if result.RowsAffected == 0 {
		return errors.New("Could not accept message request")
	}

	return nil
}
