package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	//"github.com/david-drvar/xws2021-nistagram/user_service/util/setup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserGrpcController struct {
	service        *services.UserService
	requestService *services.RegistrationRequestService
	jwtManager     *common.JWTManager
	logger         *logger.Logger
}

func NewUserController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger, redis *saga.RedisServer) (*UserGrpcController, error) {
	service, err := services.NewUserService(db, redis)
	if err != nil {
		return nil, err
	}
	requestService, err := services.NewRegistrationRequestService(db, redis)

	return &UserGrpcController{
		service,
		requestService,
		jwtManager,
		logger,
	}, nil
}

func (s *UserGrpcController) CreateAgentUser(ctx context.Context, in *protopb.CreateUserRequest) (*protopb.UsersDTO, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateAdminUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	in.User.Role = "Agent"
	user, err := s.CreateUser(ctx, in)
	if err != nil {
		return user, err
	}

	err = s.requestService.CreateRegistrationRequest(ctx, user.Id)
	if err != nil {
		return user, err
	}

	return user, nil

}

func (s *UserGrpcController) CreateUser(ctx context.Context, in *protopb.CreateUserRequest) (*protopb.UsersDTO, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("CreateUser", "User registration attempt: "+in.User.Email, logger.Info)

	var user *persistence.User
	var userAdditionalInfo persistence.UserAdditionalInfo
	user = user.ConvertFromGrpc(in.User)
	if user == nil {
		return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, "cannot create user")
	}

	userAdditionalInfo = *userAdditionalInfo.ConvertFromGrpc(in.User)

	userDomain, err := s.service.CreateUserWithAdditionalInfo(ctx, user, &userAdditionalInfo)
	if err != nil {
		s.logger.ToStdoutAndFile("CreateUser", "User registration failed: "+in.User.Email, logger.Error)
		return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, err.Error())
	}

	s.logger.ToStdoutAndFile("CreateUser", "User registration successful: "+in.User.Email, logger.Info)
	userProto := userDomain.ConvertToGrpc()
	return userProto, nil
}

func (s *UserGrpcController) GetAllUsers(ctx context.Context, in *protopb.EmptyRequest) (*protopb.UsersResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	users, err := s.service.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var usersList []*protopb.UsersDTO
	for _, user := range users {
		usersList = append(usersList, user.ConvertToGrpc())
	}

	finalResponse := protopb.UsersResponse{
		Users: usersList,
	}

	return &finalResponse, nil
}

func (s *UserGrpcController) UpdateUserProfile(ctx context.Context, in *protopb.CreateUserDTORequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserProfile")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user domain.User
	user = user.ConvertFromGrpc(in.User)
	if user.Id == "" {
		return &protopb.EmptyResponse{}, status.Errorf(codes.Unknown, "cannot convert user from grpc")
	}

	_, err := s.service.UpdateUserProfile(ctx, user)
	if err != nil {
		return &protopb.EmptyResponse{}, status.Errorf(codes.Unknown, "Could not create user")
	}

	return &protopb.EmptyResponse{}, nil
}

func (s *UserGrpcController) UpdateUserPassword(ctx context.Context, in *protopb.CreatePasswordRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("UpdateUserPassword", "Updating password attempt by user with id "+in.Password.Id, logger.Info)

	var password domain.Password
	password = password.ConvertFromGrpc(in.Password)
	if password.Id == "" {
		return &protopb.EmptyResponse{}, status.Errorf(codes.InvalidArgument, "Could not create user")
	}

	_, err := s.service.UpdateUserPassword(ctx, password)
	if err != nil {
		s.logger.ToStdoutAndFile("UpdateUserPassword", "Updating password failed by user with id "+in.Password.Id, logger.Error)
		return &protopb.EmptyResponse{}, status.Errorf(codes.InvalidArgument, "Could not create user")
	}

	s.logger.ToStdoutAndFile("UpdateUserPassword", "Updating password successful by user with id "+in.Password.Id, logger.Info)
	return &protopb.EmptyResponse{}, nil
}

func (s *UserGrpcController) SearchUser(ctx context.Context, in *protopb.SearchUserDtoRequest) (*protopb.UsersResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SearchUser")
	defer span.Finish()
	claims, _ := s.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user domain.User
	user = user.ConvertFromGrpc(in.User)

	users, err := s.service.SearchUsersByUsernameAndName(ctx, &user, claims.UserId)
	if err != nil {
		return nil, err
	}

	var usersList []*protopb.UsersDTO
	for _, user := range users {
		usersList = append(usersList, user.ConvertToGrpc())
	}

	finalResponse := protopb.UsersResponse{
		Users: usersList,
	}

	return &finalResponse, nil
}

func (s *UserGrpcController) GetUserById(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.UsersDTO, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserById")
	defer span.Finish()
	claims, _ := s.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if claims.UserId == "" {
		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
		if err != nil {
			return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, err.Error())
		}
		if !isPublic {
			return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, "this user is private")
		}
	} else if claims.UserId != in.Id {
		//isBlocked, err := grpc_common.CheckIfBlocked(ctx, in.Id, claims.UserId)
		//if err != nil {
		//	return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, err.Error())
		//}
		//
		//// If used is blocked or his profile is private and did not approve your request
		//if isBlocked {
		//	return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, "cannot retrieve this user, no connection available")
		//}
	}

	user, err := s.service.GetUser(ctx, in.Id)
	if err != nil {
		return &protopb.UsersDTO{}, err
	}

	userResponse := user.ConvertToGrpc()

	return userResponse, nil
}

func (s *UserGrpcController) GetUsernameById(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.UsersDTO, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsernameById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	username, err := s.service.GetUsername(ctx, in.Id)
	if err != nil {
		return &protopb.UsersDTO{}, err
	}

	userResponse := &protopb.UsersDTO{
		Username: username,
	}

	return userResponse, nil
}

func (s *UserGrpcController) GetPhotoById(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.UserPhoto, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPhotoById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	photo, err := s.service.GetUserPhoto(ctx, in.Id)
	if err != nil {
		return &protopb.UserPhoto{}, err
	}

	userResponse := &protopb.UserPhoto{
		Photo: photo,
	}

	return userResponse, nil
}

func (s *UserGrpcController) LoginUser(ctx context.Context, in *protopb.LoginRequest) (*protopb.LoginResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "LoginUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("LoginUser", "Login attempt by "+in.Email, logger.Info)

	var request domain.LoginRequest
	request = request.ConvertFromGrpc(in)

	user, err := s.service.LoginUser(ctx, request)
	if err != nil {
		s.logger.ToStdoutAndFile("LoginUser", "Login failed by "+in.Email, logger.Warn)
		return &protopb.LoginResponse{}, err
	}

	token, err := s.jwtManager.GenerateJwt(user.Id, user.Role.String())
	if err != nil {
		s.logger.ToStdoutAndFile("LoginUser", "JWT generate failed", logger.Error)
		return &protopb.LoginResponse{}, err
	}

	photo, err := s.service.GetUserPhoto(ctx, user.Id)
	if err != nil {
		s.logger.ToStdoutAndFile("LoginUser", "Could not retrieve user's photo", logger.Error)
		return &protopb.LoginResponse{}, err
	}

	s.logger.ToStdoutAndFile("LoginUser", "Successful login by "+in.Email, logger.Info)

	return &protopb.LoginResponse{
		AccessToken: token,
		UserId:      user.Id,
		Username:    user.Username,
		Role:        user.Role.String(),
		IsSSO:       false,
		Photo:       photo,
	}, nil
}

func (s *UserGrpcController) GetUserByEmail(ctx context.Context, in *protopb.RequestEmailUser) (*protopb.UsersDTO, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserByEmail")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	user, err := s.service.GetUserByEmail(ctx, in.Email)

	if err != nil {
		return &protopb.UsersDTO{}, err
	}

	userResponse := user.ConvertToGrpc()

	return userResponse, nil
}

func (s *UserGrpcController) ValidateResetCode(ctx context.Context, in *protopb.RequestResetCode) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "ValidateResetCode")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := s.service.ValidateResetCode(ctx, in.ResetCode, in.Email)

	if err != nil {
		return &protopb.EmptyResponse{}, status.Errorf(codes.Unknown, "Could not create user")
	}

	return &protopb.EmptyResponse{}, nil
}

func (s *UserGrpcController) ChangeForgottenPass(ctx context.Context, in *protopb.CreatePasswordRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "ChangeForgottenPass")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("ChangeForgottenPass", "Password change attempt: "+in.Password.Id, logger.Info)

	var password domain.Password
	password = password.ConvertFromGrpc(in.Password)
	if password.Id == "" {
		return &protopb.EmptyResponse{}, status.Errorf(codes.InvalidArgument, "Could not create user")
	}
	_, err := s.service.ChangeForgottenPass(ctx, password)
	if err != nil {
		s.logger.ToStdoutAndFile("ChangeForgottenPass", "Password change failed: "+in.Password.Id, logger.Error)
		return &protopb.EmptyResponse{}, status.Errorf(codes.InvalidArgument, "Could not create user")
	}

	s.logger.ToStdoutAndFile("ChangeForgottenPass", "Password change successful: "+in.Password.Id, logger.Info)
	return &protopb.EmptyResponse{}, nil
}

func (s *UserGrpcController) ApproveAccount(ctx context.Context, in *protopb.CreatePasswordRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "ApproveAccount")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("ApproveAccount", "Account approval attempt: "+in.Password.Id, logger.Info)

	var password domain.Password
	password = password.ConvertFromGrpc(in.Password)
	if password.Id == "" {
		return &protopb.EmptyResponse{}, status.Errorf(codes.InvalidArgument, "Could not create user")
	}
	_, err := s.service.ApproveAccount(ctx, password)
	if err != nil {
		s.logger.ToStdoutAndFile("ApproveAccount", "Account approval failed: "+in.Password.Id, logger.Error)
		return &protopb.EmptyResponse{}, status.Errorf(codes.InvalidArgument, "Could not create user")
	}

	s.logger.ToStdoutAndFile("ApproveAccount", "Account approval success: "+in.Password.Id, logger.Info)
	return &protopb.EmptyResponse{}, nil
}

func (s *UserGrpcController) GoogleAuth(ctx context.Context, in *protopb.GoogleAuthRequest) (*protopb.LoginResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GoogleAuth")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("GoogleAuth", "Google SSO attempt", logger.Info)

	googleToken := in.Token

	user, err := s.service.GoogleSignIn(ctx, googleToken)
	if err != nil {
		s.logger.ToStdoutAndFile("GoogleAuth", "Google SSO attempt failed", logger.Error)
		return &protopb.LoginResponse{}, status.Errorf(codes.InvalidArgument, "could not create user")
	}

	token, err := s.jwtManager.GenerateJwt(user.Id, user.Role.String())
	if err != nil {
		s.logger.ToStdoutAndFile("GoogleAuth", "JWT generate failed", logger.Error)
		return &protopb.LoginResponse{}, err
	}

	s.logger.ToStdoutAndFile("GoogleAuth", "Google SSO attempt success by "+user.Email, logger.Info)
	return &protopb.LoginResponse{
		AccessToken: token,
		UserId:      user.Id,
		Username:    user.Username,
		Role:        user.Role.String(),
		IsSSO:       true,
	}, nil
}

func (s *UserGrpcController) UpdateUserPhoto(ctx context.Context, in *protopb.UserPhotoRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := s.service.UpdateUserPhoto(ctx, in.UserId, in.Photo)

	if err != nil {
		return &protopb.EmptyResponse{}, status.Errorf(codes.InvalidArgument, "Bad request")
	}
	return &protopb.EmptyResponse{}, nil
}

func (s *UserGrpcController) CheckIsApproved(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.BooleanResponseUsers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsernameById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	isApproved, err := s.service.CheckIsApproved(ctx, in.Id)

	if err != nil {
		return &protopb.BooleanResponseUsers{Response: true}, nil
	}

	return &protopb.BooleanResponseUsers{
		Response: isApproved,
	}, nil
}

func (s *UserGrpcController) GetUserByUsername(ctx context.Context, in *protopb.RequestUsernameUser) (*protopb.UsersDTO, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserByEmail")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	user, err := s.service.GetUserByUsername(ctx, in.Username)

	if err != nil {
		return &protopb.UsersDTO{}, err
	}

	userResponse := user.ConvertToGrpc()

	return userResponse, nil
}

func (s *UserGrpcController) CheckIsActive(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.BooleanResponseUsers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIsActive")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	retVal, err := s.service.CheckIsActive(ctx, in.Id)

	if err != nil {
		return &protopb.BooleanResponseUsers{}, err
	}

	return &protopb.BooleanResponseUsers{Response: retVal}, nil
}

func (s *UserGrpcController) ChangeUserActiveStatus(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIsActive")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := s.service.ChangeUserActiveStatus(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	_, err = grpc_common.DeleteComplaintByUserId(ctx, in.Id)
	return &protopb.EmptyResponse{}, err
}

func (s *UserGrpcController) GetAllInfluncers(ctx context.Context, in *protopb.EmptyRequest) (*protopb.InfluencerSearchResult, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllInfluncers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	users, err := s.service.GetAllInfluncers(ctx)

	if err != nil {
		return &protopb.InfluencerSearchResult{}, err
	}

	var usersList []*protopb.InfluencerSearch
	for _, user := range users {
		usersList = append(usersList, user.ConvertToGrpc())
	}

	finalResponse := protopb.InfluencerSearchResult{Users: usersList}

	return &finalResponse, nil
}
