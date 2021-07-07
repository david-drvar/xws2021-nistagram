package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/domain"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/agent_application/util"
	"github.com/david-drvar/xws2021-nistagram/agent_application/util/encryption"
	"github.com/david-drvar/xws2021-nistagram/agent_application/util/images"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	LoginUser(context.Context, domain.LoginRequest) (persistence.User, error)
	CreateUser(context.Context, persistence.User) error
	SaveUserProfilePhoto(context.Context, persistence.User) error
	GetUserPhoto(ctx context.Context, id string) (string, error)
	GetUserByUsername(ctx context.Context, username string) (persistence.User, error)
	GetUserById(ctx context.Context, id string) (persistence.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) (*userRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}

	return &userRepository{DB: db}, nil
}

func (repository *userRepository) GetUserPhoto(ctx context.Context, userId string) (string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserPhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user *persistence.User
	result := repository.DB.Model(&persistence.User{}).Where("id = ?", userId).Find(&user)
	if result.Error != nil || result.RowsAffected != 1 {
		return "", result.Error
	}

	if user.ProfilePhoto != "" {
		photo, err := images.LoadImageToBase64(user.ProfilePhoto)
		if err != nil {
			return "", err
		}
		return photo, nil
	}

	return "", nil
}

func (repository *userRepository) LoginUser(ctx context.Context, request domain.LoginRequest) (persistence.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "LoginUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var dbUser persistence.User
	result := repository.DB.Where("email = ?", request.Email).First(&dbUser)
	if result.Error != nil || result.RowsAffected != 1 {
		return persistence.User{}, result.Error
	}

	err := encryption.CompareHashAndPassword([]byte(dbUser.Password), []byte(request.Password))
	if err != nil {
		return persistence.User{}, errors.New("passwords do not match")
	}

	return dbUser, nil
}

func (repository *userRepository) CreateUser(ctx context.Context, user persistence.User) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	user.Password = encryption.HashAndSalt([]byte(user.Password))

	user.Id = uuid.New().String()
	resultUser := repository.DB.Create(&user)
	if resultUser.Error != nil {
		return resultUser.Error
	}

	if user.ProfilePhoto != "" {
		err := repository.SaveUserProfilePhoto(ctx, user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repository *userRepository) SaveUserProfilePhoto(ctx context.Context, user persistence.User) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveUserProfilePhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	mimeType, err := images.GetImageType(user.ProfilePhoto)
	if err != nil {
		return err
	}

	t := time.Now()
	formatted := fmt.Sprintf("%s%d%02d%02d%02d%02d%02d%02d", user.Id, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	name := formatted + "." + mimeType

	err = images.SaveImage(name, user.ProfilePhoto)
	if err != nil {
		return err
	}

	user.ProfilePhoto = util.GetContentLocation(name)
	db := repository.DB.Model(&user).Where("id = ?", user.Id).Updates(user)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (repository *userRepository) GetUserByUsername(ctx context.Context, username string) (persistence.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserByUsername")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var dbUser persistence.User
	result := repository.DB.Where("username = ?", username).First(&dbUser)
	if result.Error != nil || result.RowsAffected != 1 {
		return persistence.User{}, result.Error
	}

	return dbUser, nil
}

func (repository *userRepository) GetUserById(ctx context.Context, id string) (persistence.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var dbUser persistence.User
	result := repository.DB.Where("id = ?", id).First(&dbUser)
	if result.Error != nil || result.RowsAffected != 1 {
		return persistence.User{}, result.Error
	}

	return dbUser, nil
}
