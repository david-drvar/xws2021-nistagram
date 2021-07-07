package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"gorm.io/gorm"
)

type RegistrationRequestController struct {
	service    *services.RegistrationRequestService
	apiKeyService *services.ApiKeyService
	jwtManager *common.JWTManager
	logger     *logger.Logger
}

func NewRegistrationRequestController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger, redis *saga.RedisServer) (*RegistrationRequestController, error) {
	service, err := services.NewRegistrationRequestService(db, redis)
	if err != nil {
		return nil, err
	}
	apiKeyService, err := services.NewApiTokenService(db)
	return &RegistrationRequestController{
		service,
		apiKeyService,
		jwtManager,
		logger,
	}, nil
}

func (controller *RegistrationRequestController) UpdateRequest(ctx context.Context, in *protopb.RegistrationRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	request := persistence.RegistrationRequest{Id: in.Id, Status: model.RequestStatus(in.Status)}
	result, err := controller.service.UpdateRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	if request.Status == "Accepted" {
		_, err = controller.apiKeyService.GenerateApiToken(ctx, result.UserId)
		if err != nil {
			return nil, err
		}
	}
	return &protopb.EmptyResponse{}, nil
}

func (controller *RegistrationRequestController) GetAllPendingRequests(ctx context.Context, in *protopb.EmptyRequest) (*protopb.ResponseRequests, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPendingRequests")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	requests, err := controller.service.GetAllPendingRequests(ctx)
	if err != nil {
		return nil, err
	}

	var retVal []*protopb.RegistrationRequest
	for _, req := range requests {
		retVal = append(retVal, req.ConvertToGrpc())

	}

	return &protopb.ResponseRequests{RegistrationRequests : retVal}, nil
}
