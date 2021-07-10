package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type NotificationRepository interface {
	CreateNotification(context.Context, *persistence.UserNotification) error
	GetUserNotifications( context.Context,  string) ([]persistence.UserNotification, error)
	DeleteNotification(context.Context, string ) (bool, error)
	ReadAllNotifications( context.Context,  string) error
	DeleteByTypeAndCreator(ctx context.Context, notification *persistence.UserNotification) error
	UpdateNotification(ctx context.Context, notification *persistence.UserNotification) error
	GetByTypeAndCreator( context.Context, *persistence.UserNotification) (*persistence.UserNotification, error)
}
type notificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepo(db *gorm.DB) (NotificationRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}

	return &notificationRepository{DB: db}, nil
}

func (repository *notificationRepository) CreateNotification(ctx context.Context, notification *persistence.UserNotification) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	//Check if already exists notification
	err := repository.CheckIfAlreadyExists(notification)
	if err != nil {
		return nil
	}

	notification.NotificationId = uuid.New().String()
	notification.CreatedAt = time.Now()
	result := repository.DB.Create(&notification)
	if result.Error != nil || result.RowsAffected != 1 {
		return errors.New("Could not create notification in repository")
	}
	return nil
}

func (repository *notificationRepository) CheckIfAlreadyExists(notification *persistence.UserNotification) error {
	var checkNotification *persistence.UserNotification

	if notification.Type!="Comment" {
		repository.DB.Where("creator_id = ?", notification.CreatorId).Where("type = ?", notification.Type).Where("user_id = ?", notification.UserId).Find(&checkNotification)
		if checkNotification.NotificationId != "" && (notification.Type=="FollowPublic" || notification.Type=="FollowPrivate"){
			return errors.New("Notification already exists")
		}
		repository.DB.Where("creator_id = ?", notification.CreatorId).Where("type = ?", "Like").Where("content_id = ?", notification.ContentId).Find(&checkNotification)
		if checkNotification.NotificationId != "" && notification.Type=="Like"{
			return errors.New("Notification already exists")
		}
		repository.DB.Where("creator_id = ?", notification.CreatorId).Where("type = ?", "Dislike").Where("content_id = ?", notification.ContentId).Find(&checkNotification)
		if checkNotification.NotificationId != "" && notification.Type=="Dislike"{
			return errors.New("Notification already exists")
		}

	}



	return nil
}

func (repository *notificationRepository) ReadAllNotifications(ctx context.Context, userId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "ReadAllNotifications")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repository.DB.Model(&persistence.UserNotification{}).Where("user_id = ?", userId ).Updates(persistence.UserNotification{IsRead: true})
	if result.Error != nil {
		return errors.New("Could not read notifications!")
	}

	return nil
}

func (repository *notificationRepository) GetUserNotifications(ctx context.Context,userId string) ([]persistence.UserNotification, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserNotifications")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var userNotifications []persistence.UserNotification

	result := repository.DB.Where("user_id = ?", userId).Find(&userNotifications)
	if result.Error != nil {
		return nil, errors.New("Error while loading notifications")
	}

	return userNotifications, nil
}

func (repository *notificationRepository) GetNotificationById(ctx context.Context, id string) (persistence.UserNotification, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var notification persistence.UserNotification
	result := repository.DB.Where("notification_id = ?", id).Find(&notification)
	if result.Error != nil || result.RowsAffected != 1 {
		return persistence.UserNotification{}, errors.New("cannot retrieve this notification")
	}

	return notification, nil
}

func (repository *notificationRepository) DeleteNotification(ctx context.Context, id string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UnBlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var notification persistence.UserNotification
	notification, _ = repository.GetNotificationById(ctx, id)

	db := repository.DB.Delete(&notification)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *notificationRepository) DeleteByTypeAndCreator(ctx context.Context, notification *persistence.UserNotification) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteByTypeAndCreator")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repository.DB.Where("creator_id = ?", notification.CreatorId).Where("user_id = ?", notification.UserId).Where("type = ?", notification.Type).Delete(&notification)
	if result.Error != nil {
		return errors.New("Could not delete notification!")
	}
	return nil

}

func (repository *notificationRepository) UpdateNotification(ctx context.Context, notification *persistence.UserNotification) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	db := repository.DB.Model(&notification).Where("notification_id = ?", notification.NotificationId).Updates(persistence.UserNotification{Type: notification.Type, Text: notification.Text, CreatedAt: time.Now()})
	if db.Error != nil {
		return errors.New("Could not update notification!")
	}else if db.RowsAffected == 0 {
		return errors.New("Could not update notification! Rows affected 0")
	}

	return nil
}

func (repository *notificationRepository) GetByTypeAndCreator(ctx context.Context, notification *persistence.UserNotification) (*persistence.UserNotification, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetByTypeAndCreator")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var retVal *persistence.UserNotification
	result := repository.DB.Where("creator_id = ?", notification.CreatorId).Where("user_id = ?", notification.UserId).Where("type = ?", notification.Type).Find(&retVal)
	if result.Error != nil {
		return nil, errors.New("Could not delete notification!")
	}

	return retVal, nil
}
