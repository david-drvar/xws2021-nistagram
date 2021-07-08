package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RegistrationRequestRepository interface {
	CreateRegistrationRequest(context.Context, string) error
	GetAllPendingRequests(context.Context) ([]persistence.RegistrationRequest, error)
	UpdateRequest(context.Context, persistence.RegistrationRequest) (*persistence.RegistrationRequest,error)

}

type registrationRequestRepository struct {
	DB                *gorm.DB
	redisServer       saga.RedisServer

}

func NewRegistrationRequestRepo(db *gorm.DB, redisServer *saga.RedisServer) (*registrationRequestRepository, error) {
	if db == nil {
		panic("RegistrationRequestRepo not created, gorm.DB is nil")
	}
	return &registrationRequestRepository{DB: db, redisServer: *redisServer}, nil
}

func (repo *registrationRequestRepository) 	CreateRegistrationRequest(ctx context.Context, userId  string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateRegistrationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repo.DB.Create(&persistence.RegistrationRequest{CreatedAt: time.Now(),Id:uuid.New().String(), UserId: userId, Status: "Pending"})
	if result.Error != nil {
		return errors.New("Could not create registration request!")
	}else if result.RowsAffected != 1 {
		return errors.New("Could not create registration request!")
	}

	return nil

}
func (repo *registrationRequestRepository) GetAllPendingRequests(ctx context.Context) ([]persistence.RegistrationRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPendingRequests")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var requests []persistence.RegistrationRequest
	result := repo.DB.Where("status = ?", "Pending").Find(&requests)
	if result.Error != nil {
		return nil, errors.New("Could not load pending requests!")
	}
	return requests, nil
}

func (repo *registrationRequestRepository) 	UpdateRequest(ctx context.Context, request persistence.RegistrationRequest) (*persistence.RegistrationRequest, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repo.DB.Where("id = ?", request.Id).Updates(&persistence.RegistrationRequest{Status: request.Status})
	if result.Error != nil {
		return nil, errors.New("Could not update request")
	}else if result.RowsAffected == 0 {
		return nil, errors.New("Could not update request")
	}
	result = repo.DB.Where("id = ?", request.Id).Find(&request)
	return &request, nil
}