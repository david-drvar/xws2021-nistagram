package controllers

import (
	"context"
	"encoding/json"
	"github.com/david-drvar/xws2021-nistagram/chat_service/model"
	"github.com/david-drvar/xws2021-nistagram/chat_service/service"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type MessageController struct {
	Service *service.MessageService
}

func NewMessageController(db *gorm.DB) (*MessageController, error) {
	service, err := service.NewMessageService(db)
	if err != nil {
		return nil, err
	}

	return &MessageController{
		Service: service,
	}, err
}

func (c MessageController) SaveMessage(w http.ResponseWriter, r *http.Request)  {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var message model.Message
	json.NewDecoder(r.Body).Decode(&message)

	err := c.Service.SaveMessage(ctx, message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	json.NewEncoder(w).Encode(message)

}

func (c MessageController)	DeleteMessage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	id := mux.Vars(r)["id"]

	err := c.Service.DeleteMessage(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}


}

func (c MessageController) GetMessagesForChatRoom(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "GetMessagesForChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	roomId := mux.Vars(r)["roomId"]

	messages, err := c.Service.GetMessagesForChatRoom(ctx, roomId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	json.NewEncoder(w).Encode(messages)
}

func (c MessageController) GetChatRoomsForUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "GetChatRoomsForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	userId := mux.Vars(r)["userId"]

	rooms, err := c.Service.GetChatRoomsForUser(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	json.NewEncoder(w).Encode(rooms)

}

func (c MessageController) StartConversation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "StartConversation")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var room model.ChatRoom
	json.NewDecoder(r.Body).Decode(&room)

	//proveri jel postoji room
	chatRoom, err := c.Service.GetChatRoomByUsers(ctx, room)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}else if chatRoom.Id != "" {
		json.NewEncoder(w).Encode(chatRoom)
		return
	}

	//proveri sve silne provere za privacy usera
	connection, err := grpc_common.CheckFollowInteraction(ctx, room.Person1, room.Person2)
	if err != nil || connection.UserId == "" {
		privacy, err := grpc_common.GetUserPrivacy(ctx, room.Person2)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte{})
			return
		}
		if !privacy.IsProfilePublic || !privacy.IsDmPublic{
			_, err := c.Service.CreateMessageRequest(ctx, &model.MessageRequest{SenderId: room.Person1, ReceiverId: room.Person2})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte{})
				return
			}
			err = grpc_common.CreateNotification(ctx, room.Person2, room.Person1, "MessageRequest", "" )
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte{})
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte{})
			return
		}
		// kreira chatRoom ukoliko je sve public kod usera
		roomRetVal, err := c.Service.CreateChatRoom(ctx, room)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte{})
			return
		}
		json.NewEncoder(w).Encode(roomRetVal)

	}

}

func (c MessageController) CreateChatRoom(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var room model.ChatRoom
	json.NewDecoder(r.Body).Decode(&room)

	var roomRetVal *model.ChatRoom
	roomRetVal, err := c.Service.CreateChatRoom(ctx, room)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	json.NewEncoder(w).Encode(roomRetVal)
}

func (c *MessageController) CreateMessageRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateMessageRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var messageRequest *model.MessageRequest
	json.NewDecoder(r.Body).Decode(&messageRequest)

	messageRequest, err := c.Service.CreateMessageRequest(ctx, messageRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	json.NewEncoder(w).Encode(messageRequest)

}

func (c *MessageController) AcceptMessageRequest(w http.ResponseWriter, r *http.Request)  {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "AcceptMessageRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var messageRequest model.MessageRequest
	json.NewDecoder(r.Body).Decode(&messageRequest)

	err := c.Service.AcceptMessageRequest(ctx, messageRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	err = grpc_common.DeleteByTypeAndCreator(ctx, "MessageRequest", messageRequest.ReceiverId, messageRequest.SenderId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	json.NewEncoder(w).Encode(messageRequest)
}

func (c *MessageController) DeclineMessageRequest(w http.ResponseWriter, r *http.Request)  {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "AcceptMessageRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var messageRequest model.MessageRequest
	json.NewDecoder(r.Body).Decode(&messageRequest)

	err := c.Service.DeclineMessageRequest(ctx, messageRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	err = grpc_common.DeleteByTypeAndCreator(ctx, "MessageRequest", messageRequest.ReceiverId, messageRequest.SenderId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	json.NewEncoder(w).Encode(messageRequest)
}

func (c *MessageController) SeenPhotoOrVideo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "SeenPhotoOrVideo")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	messageId := mux.Vars(r)["messageId"]

	err := c.Service.SeenPhotoOrVideo(ctx, messageId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

}
