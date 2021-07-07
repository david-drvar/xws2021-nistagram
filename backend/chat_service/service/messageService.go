package service

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/chat_service/model"
	"github.com/david-drvar/xws2021-nistagram/chat_service/repository"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"gorm.io/gorm"
)

type MessageService struct {
	repository repository.MessageRepository
}

func NewMessageService(db *gorm.DB) (*MessageService, error){
	repository, _ := repository.NewMessageRepository(db)
	return &MessageService{
		repository: repository,
	},nil
}

func (s MessageService) SaveMessage(ctx context.Context, message model.Message) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.SaveMessage(ctx, message)
}

func (s MessageService)	DeleteMessage(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.DeleteMessage(ctx, id)

}

func (s MessageService) GetMessagesForChatRoom(ctx context.Context, roomId string) ([]model.Message, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetMessagesForChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.GetMessagesForChatRoom(ctx, roomId)

}

func (s MessageService) GetChatRoomsForUser(ctx context.Context, userId string) ([]model.ChatRoom, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetChatRoomsForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.GetChatRoomsForUser(ctx, userId)
}

func (s MessageService) CreateChatRoom(ctx context.Context, room model.ChatRoom) (*model.ChatRoom, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.CreateChatRoom(ctx, room)

}
func (s MessageService) GetChatRoomByUsers(ctx context.Context, room model.ChatRoom) (*model.ChatRoom, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.GetChatRoomByUsers(ctx, room)
}

//Message requests
func (s *MessageService) CreateMessageRequest(ctx context.Context, messageRequest *model.MessageRequest) (*model.MessageRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateMessageRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.CreateMessageRequest(ctx, messageRequest)

}

func (s *MessageService) AcceptMessageRequest(ctx context.Context, messageRequest model.MessageRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "AcceptMessageRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.AcceptMessageRequest(ctx, messageRequest)
}

func (s *MessageService) DeclineMessageRequest(ctx context.Context, messageRequest model.MessageRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "AcceptMessageRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.DeclineMessageRequest(ctx, messageRequest)
}