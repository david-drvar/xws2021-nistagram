package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"gorm.io/gorm"
)

type NotificationService struct {
	repository  repositories.NotificationRepository
	userService *UserService
}

func NewNotificationService(db *gorm.DB, redis *saga.RedisServer) (*NotificationService, error) {
	repository, err := repositories.NewNotificationRepo(db)
	service, err := NewUserService(db, redis)
	return &NotificationService{
		repository:  repository,
		userService: service,
	}, err
}

func (s NotificationService) CreateNotification(ctx context.Context, domainNotification *domain.UserNotification) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	userNotification := &persistence.UserNotification{}
	if domainNotification.NotificationType == "MessageRequest" {
		userNotification.Text = " wants to send you a message."
	} else if domainNotification.NotificationType == "FollowPublic" {
		userNotification.Text = " started following you."
	} else if domainNotification.NotificationType == "FollowPrivate" {
		userNotification.Text = " wants to follow you."
	} else if domainNotification.NotificationType == "Like" {
		userNotification.Text = " liked your post."
	} else if domainNotification.NotificationType == "Dislike" {
		userNotification.Text = " disliked your post."
	} else if domainNotification.NotificationType == "Comment" {
		userNotification.Text = " commented on your post."
	} else if domainNotification.NotificationType == "Post" {
		userNotification.Text = " shared a post."
	} else if domainNotification.NotificationType == "Story" {
		userNotification.Text = " shared a story."
	} else if domainNotification.NotificationType == "Campaign" {
		userNotification.Text = " wants to hire you for campaign ."
	} else {
		return errors.New("Bad notification type")
	}

	userNotification.UserId = domainNotification.UserId
	userNotification.CreatorId = domainNotification.CreatorId
	userNotification.Type = domainNotification.NotificationType
	userNotification.IsRead = false
	userNotification.ContentId = domainNotification.ContentId

	err := s.repository.CreateNotification(ctx, userNotification)
	if err != nil {
		return err
	}

	return nil
}

func (s NotificationService) GetUserNotifications(ctx context.Context, userId string) ([]persistence.UserNotification, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserNotifications")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	notifications, err := s.repository.GetUserNotifications(ctx, userId)
	if err != nil {
		return nil, err
	}

	for _, n := range notifications {
		creatorUsername, err := s.userService.GetUsername(ctx, n.CreatorId)
		if err == nil {
			n.Text = creatorUsername + " " + n.Text
		}
	}

	return notifications, nil
}

func (s *NotificationService) ReadAllNotifications(ctx context.Context, userId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "ReadAllNotifications")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.ReadAllNotifications(ctx, userId)

}

func (s NotificationService) DeleteNotification(ctx context.Context, id string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UnBlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.DeleteNotification(ctx, id)
}

func (s NotificationService) DeleteByTypeAndCreator(ctx context.Context, notification *persistence.UserNotification) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteByTypeAndCreator")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.DeleteByTypeAndCreator(ctx, notification)

}

func (s NotificationService) UpdateNotification(ctx context.Context, notification *persistence.UserNotification) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.UpdateNotification(ctx, notification)
}

func (s NotificationService) GetByTypeAndCreator(ctx context.Context, notification *persistence.UserNotification) (*persistence.UserNotification, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetByTypeAndCreator")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.GetByTypeAndCreator(ctx, notification)
}
