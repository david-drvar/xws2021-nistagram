package controllers

import (
	"context"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"gorm.io/gorm"
)

type NotificationGrpcController struct {
	service *services.NotificationService
}

func NewNotificationController(db *gorm.DB, redis *saga.RedisServer) (*NotificationGrpcController, error) {
	service, err := services.NewNotificationService(db, redis)
	if err != nil {
		return nil, err
	}

	return &NotificationGrpcController{
		service: service,
	}, nil
}

func (c *NotificationGrpcController) CreateNotification(ctx context.Context, in *protopb.CreateNotificationRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var notification *domain.UserNotification
	notification = notification.ConvertFromGrpc(*in)
	err := c.service.CreateNotification(ctx, notification)
	if err != nil {
		return &protopb.EmptyResponse{}, err
	}

	return &protopb.EmptyResponse{}, nil
}

func (c *NotificationGrpcController) GetUserNotifications(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.CreateNotificationResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserNotifications")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	notifications, err := c.service.GetUserNotifications(ctx, in.Id)
	if err != nil {
		return &protopb.CreateNotificationResponse{}, err
	}

	var response []*protopb.Notification
	for _, n := range notifications {
		response = append(response, n.ConvertToGrpc())
	}

	return &protopb.CreateNotificationResponse{
		Notifications: response,
	}, nil
}
func (c *NotificationGrpcController) DeleteNotification(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UnBlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := c.service.DeleteNotification(ctx, in.Id)
	if err != nil {
		return &protopb.EmptyResponse{}, err
	}

	return &protopb.EmptyResponse{}, nil
}

func (c *NotificationGrpcController) ReadAllNotifications(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "ReadAllNotifications")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := c.service.ReadAllNotifications(ctx, in.Id)
	if err != nil {
		return &protopb.EmptyResponse{}, err
	}

	return &protopb.EmptyResponse{}, nil
}

func (c *NotificationGrpcController) DeleteByTypeAndCreator(ctx context.Context, in *protopb.Notification) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteByTypeAndCreator")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := c.service.DeleteByTypeAndCreator(ctx, &persistence.UserNotification{CreatorId: in.CreatorId, Type: in.Type, UserId: in.UserId})
	if err != nil {
		return &protopb.EmptyResponse{}, err
	}

	return &protopb.EmptyResponse{}, nil
}

func (c *NotificationGrpcController) UpdateNotification(ctx context.Context, in *protopb.Notification) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := c.service.UpdateNotification(ctx, &persistence.UserNotification{NotificationId: in.Id, Type: in.Type, Text: in.Text})
	if err != nil {
		return &protopb.EmptyResponse{}, err
	}
	return &protopb.EmptyResponse{}, nil
}

func (c *NotificationGrpcController) GetByTypeAndCreator(ctx context.Context, in *protopb.Notification) (*protopb.Notification, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetByTypeAndCreator")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	notification, err := c.service.GetByTypeAndCreator(ctx, &persistence.UserNotification{UserId: in.UserId, CreatorId: in.CreatorId, Type: in.Type})
	if err != nil {
		return nil, err
	}

	return notification.ConvertToGrpc(), nil
}
