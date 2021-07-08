package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"gorm.io/gorm"
)

type RegistrationRequestService struct {
	registrationRequestRepository repositories.RegistrationRequestRepository
	service *UserService
}

func NewRegistrationRequestService(db *gorm.DB,redis *saga.RedisServer) (*RegistrationRequestService, error) {
	registrationRequestRepo, err := repositories.NewRegistrationRequestRepo(db,redis)
	service, err := NewUserService(db,redis)
	return &RegistrationRequestService{
		registrationRequestRepository: registrationRequestRepo,
		service: service,
	}, err
}

func (service *RegistrationRequestService) CreateRegistrationRequest(ctx context.Context, userId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateRegistrationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.registrationRequestRepository.CreateRegistrationRequest(ctx, userId)
	return err

}

func (service *RegistrationRequestService) GetAllPendingRequests(ctx context.Context) ([]domain.RegistrationRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPendingRequests")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	requests, err := service.registrationRequestRepository.GetAllPendingRequests(ctx)
	if err != nil {
		return nil, err
	}

	var retVal []domain.RegistrationRequest
	for _, req := range requests {
		user, _ := service.service.GetUser(ctx, req.UserId)
		retVal = append(retVal, domain.RegistrationRequest{Id: req.Id, UserId: req.UserId,
			CreatedAt: req.CreatedAt, Email: user.Email,
			Username: user.Username, Website: user.Website,
			FirstName: user.FirstName, LastName: user.LastName, Status: req.Status})
	}

	return retVal, nil

}

func (service *RegistrationRequestService) UpdateRequest(ctx context.Context, request persistence.RegistrationRequest) (*persistence.RegistrationRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.registrationRequestRepository.UpdateRequest(ctx, request)
}