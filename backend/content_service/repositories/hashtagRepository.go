package repositories

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HashtagRepository interface {
	CreateHashtag(context.Context, string) (*domain.Hashtag, error)
	GetHashtagByText(context.Context, string) (*domain.Hashtag, error)
	GetPostIdsByHashtag(context.Context, persistence.Hashtag) ([]string, error)
	GetAllHashtags(context.Context) ([]domain.Hashtag, error)
	BindPostWithHashtags(context.Context, string, []persistence.Hashtag) error
	GetPostHashtags(context.Context, string) ([]domain.Hashtag, error)
}

type hashtagRepository struct {
	DB *gorm.DB
}

func NewHashtagRepo(db *gorm.DB) (*hashtagRepository, error) {
	if db == nil {
		panic("HashtagRepository not created, gorm.DB is nil")
	}

	return &hashtagRepository{DB: db}, nil
}

func (repository *hashtagRepository) CreateHashtag(ctx context.Context, text string) (*domain.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var hashtag persistence.Hashtag

	hashtag.Id = uuid.New().String()
	hashtag.Text = text
	resultHashtag := repository.DB.Create(&hashtag)
	if resultHashtag.Error != nil {
		return nil, resultHashtag.Error
	}

	return &domain.Hashtag{Id: hashtag.Id, Text: hashtag.Text}, nil
}

func (repository *hashtagRepository) GetHashtagByText(ctx context.Context, text string) (*domain.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetHashtagByText")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var hashtag persistence.Hashtag

	resultHashtag := repository.DB.Where("text = ?", text).Find(&hashtag)
	if resultHashtag.Error != nil {
		return nil, resultHashtag.Error
	}

	return &domain.Hashtag{Id: hashtag.Id, Text: hashtag.Text}, nil
}

func (repository *hashtagRepository) GetPostIdsByHashtag(ctx context.Context, hashtag persistence.Hashtag) ([]string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostIdsByHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var hashtagObjavas []persistence.HashtagObjava

	resultHashtagObjava := repository.DB.Where("hashtag_id = ?", hashtag.Id).Find(&hashtagObjavas)
	if resultHashtagObjava.Error != nil {
		return nil, resultHashtagObjava.Error
	}

	var postIds []string

	for _, hashtagObjava := range hashtagObjavas {
		postIds = append(postIds, hashtagObjava.ObjavaId)
	}

	return postIds, nil
}

func (repository *hashtagRepository) GetAllHashtags(ctx context.Context) ([]domain.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHashtags")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var hashtags []persistence.Hashtag
	var hashtagsDomain []domain.Hashtag

	repository.DB.Find(&hashtags)

	for _, hashtag := range hashtags {
		hashtagsDomain = append(hashtagsDomain, domain.Hashtag{Id: hashtag.Id, Text: hashtag.Text})
	}

	return hashtagsDomain, nil
}

func (repository *hashtagRepository) BindPostWithHashtags(ctx context.Context, postId string, hashtags []persistence.Hashtag) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "BindPostWithHashtags")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var hashtagObjavas []persistence.HashtagObjava

	for _, hashtag := range hashtags {
		hashtagObjavas = append(hashtagObjavas, persistence.HashtagObjava{HashtagId: hashtag.Id, ObjavaId: postId})
	}

	resultHashtagObjava := repository.DB.Create(&hashtagObjavas)
	if resultHashtagObjava.Error != nil {
		return resultHashtagObjava.Error
	}

	return nil
}

func (repository *hashtagRepository) GetPostHashtags(ctx context.Context, postId string) ([]domain.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostHashtags")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var hashtags []persistence.Hashtag
	var hashtagsObjavas []persistence.HashtagObjava
	var hashtagsDomain []domain.Hashtag

	result := repository.DB.Where("objava_id = ?", postId).Find(&hashtagsObjavas)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, hashtagObjava := range hashtagsObjavas {
		var hashtag persistence.Hashtag
		result = repository.DB.Where("id = ?", hashtagObjava.HashtagId).Find(&hashtag)
		if result.Error != nil {
			return nil, result.Error
		}
		hashtags = append(hashtags, hashtag)
	}

	for _, hashtag := range hashtags {
		hashtagsDomain = append(hashtagsDomain, domain.Hashtag{Id: hashtag.Id, Text: hashtag.Text})
	}

	return hashtagsDomain, nil
}
