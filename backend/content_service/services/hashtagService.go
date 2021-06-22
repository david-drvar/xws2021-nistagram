package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type HashtagService struct {
	hashtagRepository repositories.HashtagRepository
}

func NewHashtagService(db *gorm.DB) (*HashtagService, error) {
	hashtagRepository, err := repositories.NewHashtagRepo(db)
	if err != nil {
		return nil, err
	}

	return &HashtagService{
		hashtagRepository,
	}, err
}

func (service HashtagService) CreateHashtag(ctx context.Context, text string) (*domain.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if text == "" {
		return nil, errors.New("cannot create empty hashtag")
	}

	hashtag, err := service.hashtagRepository.CreateHashtag(ctx, text)
	if err != nil {
		return nil, errors.New("hashtag does not exist")
	}

	return &domain.Hashtag{Id: hashtag.Id, Text: hashtag.Text}, nil
}

func (service HashtagService) GetHashtagByText(ctx context.Context, text string) (*domain.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetHashtagByText")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if text == "" {
		return nil, errors.New("cannot get empty hashtag")
	}

	hashtag, err := service.hashtagRepository.GetHashtagByText(ctx, text)
	if err != nil {
		return nil, errors.New("hashtag does not exist")
	}

	return &domain.Hashtag{Id: hashtag.Id, Text: hashtag.Text}, nil
}

func (service HashtagService) GetAllHashtags(ctx context.Context) ([]domain.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHashtags")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	hashtags, err := service.hashtagRepository.GetAllHashtags(ctx)
	if err != nil {
		return nil, errors.New("hashtag does not exist")
	}

	return hashtags, nil
}

func (service HashtagService) GetPostHashtags(ctx context.Context, postId string) ([]domain.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostHashtags")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.hashtagRepository.GetPostHashtags(ctx, postId)
}
