package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/domain"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/agent_application/services"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserGrpcController struct {
	service    *services.UserService
	jwtManager *common.JWTManager
	logger     *logger.Logger
}

func NewUserController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger) (*UserGrpcController, error) {
	service, err := services.NewUserService(db)
	if err != nil {
		return nil, err
	}

	return &UserGrpcController{
		service,
		jwtManager,
		logger,
	}, nil
}

func (s *UserGrpcController) LoginUserInAgentApp(ctx context.Context, in *protopb.LoginRequestAgentApp) (*protopb.LoginResponseAgentApp, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "LoginUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var request domain.LoginRequest
	request = request.ConvertFromGrpc(in)

	user, err := s.service.LoginUser(ctx, request)
	if err != nil {
		return &protopb.LoginResponseAgentApp{}, err
	}

	token, err := s.jwtManager.GenerateJwt(user.Id, user.Role.String())
	if err != nil {
		s.logger.ToStdoutAndFile("LoginUser", "JWT generate failed", logger.Error)
		return &protopb.LoginResponseAgentApp{}, err
	}

	photo, err := s.service.GetUserPhoto(ctx, user.Id)
	if err != nil {
		s.logger.ToStdoutAndFile("LoginUser", "Could not retrieve user's photo", logger.Error)
		return &protopb.LoginResponseAgentApp{}, err
	}

	s.logger.ToStdoutAndFile("LoginUser", "Successful login by "+in.Email, logger.Info)

	return &protopb.LoginResponseAgentApp{
		AccessToken: token,
		UserId:      user.Id,
		Username:    user.Username,
		Role:        user.Role.String(),
		IsSSO:       false,
		Photo:       photo,
	}, nil
}

func (s *UserGrpcController) CreateUserInAgentApp(ctx context.Context, in *protopb.CreateUserRequestAgentApp) (*protopb.EmptyResponseAgent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserInAgentApp")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user *persistence.User
	user = user.ConvertFromGrpc(in.User)
	if user == nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, "cannot create user")
	}

	err := s.service.CreateUserInAgentApp(ctx, *user)
	if err != nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseAgent{}, nil
}

func (s *UserGrpcController) GetUserByUsername(ctx context.Context, in *protopb.RequestUsernameAgent) (*protopb.UserAgentApp, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserByEmail")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	user, err := s.service.GetUserByUsername(ctx, in.Username)

	if err != nil {
		return &protopb.UserAgentApp{}, err
	}

	userResponse := user.ConvertToGrpc()

	return userResponse, nil
}
