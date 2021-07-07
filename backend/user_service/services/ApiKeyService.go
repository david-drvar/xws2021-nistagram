package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"gorm.io/gorm"
	"time"
)

type ApiKeyService struct {
	repository repositories.ApiKeyRepository
	jwtManager *common.JWTManager
}

func NewApiTokenService (db *gorm.DB) (*ApiKeyService, error) {
	repo, err := repositories.NewApiTokenRepository(db)
	jwtManager := common.NewJWTManager("xs-dawer", 500000 * time.Minute)

	return &ApiKeyService{repository: repo, jwtManager: jwtManager}, err
}

func (s ApiKeyService) GenerateApiToken(ctx context.Context, id string) (string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GenerateApiToken")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	token, err := s.jwtManager.GenerateJwt(id, "Agent")
	if err != nil {
		return "", err
	}

	err = s.repository.SaveApiToken(ctx,&persistence.APIKeys{APIKey: token, UserId: id})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s ApiKeyService) GetKeyByUserId(ctx context.Context, id string) (string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetKeyByUserId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return s.repository.GetKeyByUserId(ctx, id)
}

func (s ApiKeyService) ValidateKey(ctx context.Context, token string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "ValidateKey")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := s.jwtManager.ValidateJWT(token)
	return err
}
