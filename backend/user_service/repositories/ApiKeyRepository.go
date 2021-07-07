package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"gorm.io/gorm"
)

type ApiKeyRepository interface {
	SaveApiToken( context.Context, *persistence.APIKeys) error
	GetKeyByUserId( context.Context,  string) (string, error)

}

type apiKeyRepository struct {
	DB *gorm.DB
}

func NewApiTokenRepository(db *gorm.DB) (ApiKeyRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}

	return &apiKeyRepository{DB: db}, nil
}

func (repo *apiKeyRepository) SaveApiToken( ctx context.Context, apiKey *persistence.APIKeys) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveApiToken")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)


	result := repo.DB.Where("user_id = ?", apiKey.UserId).Delete(&persistence.APIKeys{})
	if result.Error != nil {
		return errors.New("Could not delete api key!")
	}

	result = repo.DB.Create(&apiKey)
	if result.Error != nil {
		return result.Error
	}else if result.RowsAffected == 0 {
		return errors.New("Could not save api key!")
	}
	return nil
}

func (repo *apiKeyRepository) GetKeyByUserId(ctx context.Context, id string) (string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetTokenByUserId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var key *persistence.APIKeys
	result := repo.DB.Where("user_id = ?", id).Find(&key)
	if result.Error != nil {
		return "", errors.New("Could not load api key for user")
	}

	return key.APIKey, nil
}



