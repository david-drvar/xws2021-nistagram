package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"

	//"github.com/david-drvar/xws2021-nistagram/user_service/util/setup"
	otgo "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"io"
)

type Server struct {
	protopb.UnimplementedUsersServer
	protopb.UnimplementedPrivacyServer
	userController         *UserGrpcController
	privacyController      *PrivacyGrpcController
	emailController   	   *EmailGrpcController
	notificationController *NotificationGrpcController
	registrationRequestController *RegistrationRequestController
	apiTokenController *ApiTokenGrpcController
	tracer            otgo.Tracer
	closer            io.Closer
	verificationController *VerificationGrpcController

}

func NewServer(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger, redis *saga.RedisServer) (*Server, error) {
	newUserController, _ := NewUserController(db, jwtManager, logger, redis)
	newPrivacyController, _ := NewPrivacyController(db, redis)
	newEmailController, _ := NewEmailController(db, redis)
	notificationController, _ := NewNotificationController(db, redis)
	newVerificationController, _ := NewVerificationController(db, jwtManager, logger, redis)
	newRegistrationRequestController, _ := NewRegistrationRequestController(db, jwtManager, logger, redis)
	newApiTokenController, _ := NewApiTokenGrpcController(db, jwtManager, logger)

	tracer, closer := tracer.Init("userService")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		userController:         newUserController,
		privacyController:      newPrivacyController,
		emailController:        newEmailController,
		notificationController: notificationController,
		verificationController: newVerificationController,
		registrationRequestController: newRegistrationRequestController,
		apiTokenController: newApiTokenController,
		tracer:                 tracer,
		closer:                 closer,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

func (s *Server) CreateUser(ctx context.Context, in *protopb.CreateUserRequest) (*protopb.UsersDTO, error) {
	return s.userController.CreateUser(ctx, in)
}

func (s *Server) GetAllUsers(ctx context.Context, in *protopb.EmptyRequest) (*protopb.UsersResponse, error) {
	return s.userController.GetAllUsers(ctx, in)
}

func (s *Server) GetUserById(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.UsersDTO, error) {
	return s.userController.GetUserById(ctx, in)
}

func (s *Server) GetUsernameById(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.UsersDTO, error) {
	return s.userController.GetUsernameById(ctx, in)
}

func (s *Server) GetPhotoById(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.UserPhoto, error) {
	return s.userController.GetPhotoById(ctx, in)
}

func (s *Server) UpdateUserProfile(ctx context.Context, in *protopb.CreateUserDTORequest) (*protopb.EmptyResponse, error) {
	return s.userController.UpdateUserProfile(ctx, in)
}

func (s *Server) UpdateUserPassword(ctx context.Context, in *protopb.CreatePasswordRequest) (*protopb.EmptyResponse, error) {
	return s.userController.UpdateUserPassword(ctx, in)
}

func (s *Server) CreatePrivacy(ctx context.Context, in *protopb.CreatePrivacyRequest) (*protopb.EmptyResponsePrivacy, error) {
	return s.privacyController.CreatePrivacy(ctx, in)
}

func (s *Server) UpdatePrivacy(ctx context.Context, in *protopb.CreatePrivacyRequest) (*protopb.EmptyResponsePrivacy, error) {
	return s.privacyController.UpdatePrivacy(ctx, in)
}

func (s *Server) BlockUser(ctx context.Context, in *protopb.CreateBlockRequest) (*protopb.EmptyResponsePrivacy, error) {
	return s.privacyController.BlockUser(ctx, in)
}

func (s *Server) UnBlockUser(ctx context.Context, in *protopb.CreateBlockRequest) (*protopb.EmptyResponsePrivacy, error) {
	return s.privacyController.UnBlockUser(ctx, in)
}

func (s *Server) CheckIfBlocked(ctx context.Context, in *protopb.CreateBlockRequest) (*protopb.BooleanResponse, error) {
	return s.privacyController.CheckIfBlocked(ctx, in)
}

func (s *Server) SearchUser(ctx context.Context, in *protopb.SearchUserDtoRequest) (*protopb.UsersResponse, error) {
	return s.userController.SearchUser(ctx, in)
}

func (s *Server) CheckUserProfilePublic(ctx context.Context, in *protopb.PrivacyRequest) (*protopb.BooleanResponse, error) {
	return s.privacyController.CheckUserProfilePublic(ctx, in)
}

func (s *Server) GetAllPublicUsers(ctx context.Context, in *protopb.RequestIdPrivacy) (*protopb.StringArray, error) {
	return s.privacyController.GetAllPublicUsers(ctx, in)
}

func (s *Server) LoginUser(ctx context.Context, in *protopb.LoginRequest) (*protopb.LoginResponse, error) {
	return s.userController.LoginUser(ctx, in)
}

func (s *Server) SendEmail(ctx context.Context, in *protopb.SendEmailDtoRequest) (*protopb.EmptyResponse, error) {
	return s.emailController.SendEmail(ctx, in)
}

func (s *Server) GetUserByEmail(ctx context.Context, in *protopb.RequestEmailUser) (*protopb.UsersDTO, error) {
	return s.userController.GetUserByEmail(ctx, in)
}

func (s *Server) ValidateResetCode(ctx context.Context, in *protopb.RequestResetCode) (*protopb.EmptyResponse, error) {
	return s.userController.ValidateResetCode(ctx, in)
}

func (s *Server) ChangeForgottenPass(ctx context.Context, in *protopb.CreatePasswordRequest) (*protopb.EmptyResponse, error) {
	return s.userController.ChangeForgottenPass(ctx, in)
}

func (s *Server) ApproveAccount(ctx context.Context, in *protopb.CreatePasswordRequest) (*protopb.EmptyResponse, error) {
	return s.userController.ApproveAccount(ctx, in)
}

func (s *Server) GoogleAuth(ctx context.Context, in *protopb.GoogleAuthRequest) (*protopb.LoginResponse, error) {
	return s.userController.GoogleAuth(ctx, in)
}

func (s *Server) UpdateUserPhoto(ctx context.Context, in *protopb.UserPhotoRequest) (*protopb.EmptyResponse, error) {
	return s.userController.UpdateUserPhoto(ctx, in)
}

func (s *Server) CreateNotification(ctx context.Context, in *protopb.CreateNotificationRequest) (*protopb.EmptyResponse, error) {
	return s.notificationController.CreateNotification(ctx, in)
}

func (s *Server) DeleteNotification(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.EmptyResponse, error) {
	return s.notificationController.DeleteNotification(ctx, in)
}

func (s *Server) CheckIsApproved(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.BooleanResponseUsers, error) {
	return s.userController.CheckIsApproved(ctx, in)
}
func (s *Server) GetUserByUsername(ctx context.Context, in *protopb.RequestUsernameUser) (*protopb.UsersDTO, error) {
	return s.userController.GetUserByUsername(ctx, in)
}
func (s *Server) SubmitVerificationRequest(ctx context.Context, in *protopb.VerificationRequest) (*protopb.EmptyResponse, error) {
	return s.verificationController.SubmitVerificationRequest(ctx, in)
}

func (s *Server) GetPendingVerificationRequests(ctx context.Context, in *protopb.EmptyRequest) (*protopb.VerificationRequestsArray, error) {
	return s.verificationController.GetPendingVerificationRequests(ctx, in)
}

func (s *Server) ChangeVerificationRequestStatus(ctx context.Context, in *protopb.VerificationRequest) (*protopb.EmptyResponse, error) {
	return s.verificationController.ChangeVerificationRequestStatus(ctx, in)
}

func (s *Server) GetVerificationRequestsByUserId(ctx context.Context, in *protopb.VerificationRequest) (*protopb.VerificationRequestsArray, error) {
	return s.verificationController.GetVerificationRequestsByUserId(ctx, in)
}

func (s *Server) GetAllVerificationRequests(ctx context.Context, in *protopb.EmptyRequest) (*protopb.VerificationRequestsArray, error) {
	return s.verificationController.GetAllVerificationRequests(ctx, in)
}

func (s *Server) GetUserNotifications(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.CreateNotificationResponse, error) {
	return s.notificationController.GetUserNotifications(ctx, in)
}

func (s *Server) GetBlockedUsers(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.ResponseIdUsers, error) {
	return s.privacyController.GetBlockedUsers(ctx, in)
}

func (s *Server) GetUserPrivacy(ctx context.Context, in *protopb.RequestIdPrivacy) (*protopb.PrivacyMessage, error) {
	return s.privacyController.GetUserPrivacy(ctx, in)
}

func (s *Server) ReadAllNotifications(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.EmptyResponse,error) {
	return s.notificationController.ReadAllNotifications(ctx, in)
}

func (s *Server) DeleteByTypeAndCreator(ctx context.Context, in *protopb.Notification) (*protopb.EmptyResponse, error) {
	return s.notificationController.DeleteByTypeAndCreator(ctx, in)
}

func (s *Server) GetByTypeAndCreator(ctx context.Context,in *protopb.Notification) (*protopb.Notification, error) {
	return s.notificationController.GetByTypeAndCreator(ctx, in)
}

func (s *Server) UpdateNotification(ctx context.Context, in *protopb.Notification) (*protopb.EmptyResponse, error) {
	return s.notificationController.UpdateNotification(ctx, in)
}
func (s *Server) CheckIsActive(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.BooleanResponseUsers, error) {
 	return s.userController.CheckIsActive(ctx, in)
}

func (s *Server) ChangeUserActiveStatus(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.EmptyResponse ,error) {
	return s.userController.ChangeUserActiveStatus(ctx, in)
}

func (s *Server) CreateAgentUser(ctx context.Context, in *protopb.CreateUserRequest) (*protopb.UsersDTO, error) {
	return s.userController.CreateAgentUser(ctx, in)
}

func (s *Server) GetAllPendingRequests(ctx context.Context, in *protopb.EmptyRequest) (*protopb.ResponseRequests, error) {
	return s.registrationRequestController.GetAllPendingRequests(ctx, in)
}

func (s *Server) UpdateRequest(ctx context.Context, in *protopb.RegistrationRequest) (*protopb.EmptyResponse, error) {
	return s.registrationRequestController.UpdateRequest(ctx, in)
}

