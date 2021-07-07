package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type VerificationGrpcController struct {
	service    *services.VerificationService
	jwtManager *common.JWTManager
	logger     *logger.Logger
}

func NewVerificationController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger, redis *saga.RedisServer) (*VerificationGrpcController, error) {
	service, err := services.NewVerificationService(db, redis)
	if err != nil {
		return nil, err
	}

	return &VerificationGrpcController{
		service,
		jwtManager,
		logger,
	}, nil
}

func (s *VerificationGrpcController) SubmitVerificationRequest(ctx context.Context, in *protopb.VerificationRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SubmitVerificationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("SubmitVerificationRequest", "User verification request submit attempt: "+in.UserId, logger.Info)

	var verificationRequest domain.VerificationRequest
	verificationRequest = verificationRequest.ConvertFromGrpc(in)

	err := s.service.CreateVerificationRequest(ctx, verificationRequest)
	if err != nil {
		return &protopb.EmptyResponse{}, status.Errorf(codes.Unknown, "Could not create verification request")
	}

	return &protopb.EmptyResponse{}, nil
}

func (s *VerificationGrpcController) GetPendingVerificationRequests(ctx context.Context, in *protopb.EmptyRequest) (*protopb.VerificationRequestsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPendingVerificationRequests")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("GetPendingVerificationRequests", "Get pending requests attempt", logger.Info)

	verificationRequests, err := s.service.GetPendingVerificationRequests(ctx)
	if err != nil {
		return &protopb.VerificationRequestsArray{}, status.Errorf(codes.Unknown, "Could not get pending verification requests")
	}

	responseVerificationRequests := []*protopb.VerificationRequest{}
	for _, verificationRequest := range verificationRequests {
		responseVerificationRequests = append(responseVerificationRequests, verificationRequest.ConvertToGrpc())
	}

	return &protopb.VerificationRequestsArray{
		VerificationRequests: responseVerificationRequests,
	}, nil
}

func (s *VerificationGrpcController) ChangeVerificationRequestStatus(ctx context.Context, in *protopb.VerificationRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "ChangeVerificationRequestStatus")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("ChangeVerificationRequestStatus", "Verification request status change attempt: "+in.UserId, logger.Info)

	var verificationRequest domain.VerificationRequest
	verificationRequest = verificationRequest.ConvertFromGrpc(in)

	err := s.service.ChangeVerificationRequestStatus(ctx, verificationRequest)
	if err != nil {
		return &protopb.EmptyResponse{}, status.Errorf(codes.Unknown, "Could not change verification request status")
	}

	return &protopb.EmptyResponse{}, nil
}

func (s *VerificationGrpcController) GetVerificationRequestsByUserId(ctx context.Context, in *protopb.VerificationRequest) (*protopb.VerificationRequestsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetVerificationRequestsByUserId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("GetVerificationRequestsByUserId", "Get verification request by user attempt: "+in.UserId, logger.Info)

	verificationRequests, err := s.service.GetVerificationRequestsByUserId(ctx, in.UserId)
	if err != nil {
		return &protopb.VerificationRequestsArray{}, status.Errorf(codes.Unknown, "Could not get verification requests by user")
	}

	responseVerificationRequests := []*protopb.VerificationRequest{}
	for _, verificationRequest := range verificationRequests {
		responseVerificationRequests = append(responseVerificationRequests, verificationRequest.ConvertToGrpc())
	}

	return &protopb.VerificationRequestsArray{
		VerificationRequests: responseVerificationRequests,
	}, nil
}

func (s *VerificationGrpcController) GetAllVerificationRequests(ctx context.Context, in *protopb.EmptyRequest) (*protopb.VerificationRequestsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllVerificationRequests")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("GetAllVerificationRequests", "Get all requests attempt", logger.Info)

	verificationRequests, err := s.service.GetAllVerificationRequests(ctx)
	if err != nil {
		return &protopb.VerificationRequestsArray{}, status.Errorf(codes.Unknown, "Could not get pending verification requests")
	}

	responseVerificationRequests := []*protopb.VerificationRequest{}
	for _, verificationRequest := range verificationRequests {
		responseVerificationRequests = append(responseVerificationRequests, verificationRequest.ConvertToGrpc())
	}

	return &protopb.VerificationRequestsArray{
		VerificationRequests: responseVerificationRequests,
	}, nil
}
