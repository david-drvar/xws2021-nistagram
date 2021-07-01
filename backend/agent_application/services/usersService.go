package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/domain"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/agent_application/repositories"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(db *gorm.DB) (*UserService, error) {
	userRepository, err := repositories.NewUserRepo(db)

	return &UserService{
		userRepository,
	}, err
}

func (service *UserService) LoginUser(ctx context.Context, request domain.LoginRequest) (persistence.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "LoginUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if request.Email == "" || request.Password == "" {
		return persistence.User{}, errors.New("empty login request")
	}

	return service.userRepository.LoginUser(ctx, request)
}

func (service *UserService) CreateUserInAgentApp(ctx context.Context, user persistence.User) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserInAgentApp")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) GetUserPhoto(ctx context.Context, userId string) (string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if userId == "" {
		return "", errors.New("can't find photo from non-existing user")
	}

	return service.userRepository.GetUserPhoto(ctx, userId)
}
