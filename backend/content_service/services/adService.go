package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type AdService struct {
	adRepository    	repositories.AdRepository
	campaignRepository  repositories.CampaignRepository
	postService  		*PostService
	storyService 		*StoryService
}

func NewAdService(db *gorm.DB) (*AdService, error){
	adRepository, err := repositories.NewAdRepo(db)
	if err != nil { return nil, err }

	campaignRepository, err := repositories.NewCampaignRepo(db)
	if err != nil { return nil, err }

	postService, err := NewPostService(db)
	if err != nil { return nil, err }

	storyService, err := NewStoryService(db)
	if err != nil { return nil, err }

	return &AdService{
		adRepository,
		campaignRepository,
		postService,
		storyService,
	}, err
}


func (service *AdService) CreateAd(ctx context.Context, ad domain.Ad) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateAd")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.adRepository.CreateAd(ctx, ad)
}

func (service *AdService) GetAdsFromCampaign(ctx context.Context, campaignId string) ([]domain.Ad, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateAd")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbAds, err := service.adRepository.GetAdsFromCampaign(ctx, campaignId)
	if err != nil { return []domain.Ad{}, err }

	ads := []domain.Ad{}
	for _, dbAd := range dbAds {
		if dbAd.Type == model.TypePost.String(){
			post, err := service.postService.GetPostById(ctx, dbAd.PostId)
			if err != nil { return []domain.Ad{}, err }
			ads = append(ads, dbAd.ConvertToDomain(post.Comments, post.Likes, post.Dislikes, post.Objava))
		}else if dbAd.Type == model.TypeStory.String() {
			story, err := service.storyService.GetStoryById(ctx, dbAd.PostId)
			if err != nil {
				return []domain.Ad{}, err
			}
			ads = append(ads, dbAd.ConvertToDomain([]domain.Comment{}, []domain.Like{}, []domain.Like{}, story.Objava))
		}
	}

	return ads, nil
}

func (service *AdService) GetAdCategories(ctx context.Context) ([]domain.AdCategory, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdCategories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCategories, err := service.adRepository.GetAdCategories(ctx)
	if err != nil { return []domain.AdCategory{}, err }

	categories := []domain.AdCategory{}
	for _, category := range dbCategories{
		categories = append(categories, category.ConvertToDomain())
	}

	return categories, nil
}

func (service *AdService) GetAdCategory(ctx context.Context, id string) (domain.AdCategory, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdCategory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCategory, err := service.adRepository.GetAdCategory(ctx, id)
	if err != nil { return domain.AdCategory{}, err }

	return dbCategory.ConvertToDomain(), nil
}

func (service *AdService) CreateAdCategory(ctx context.Context, category domain.AdCategory) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateAdCategory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.adRepository.CreateAdCategory(ctx, category)
	if err != nil { return err }

	return nil
}

func (service *AdService) GetUserAdCategories(ctx context.Context, userId string) ([]domain.AdCategory, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserAdCategories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCategories, err := service.adRepository.GetUserAdCategories(ctx, userId)
	if err != nil { return []domain.AdCategory{}, err }

	categories := []domain.AdCategory{}
	for _, dbCategory := range dbCategories{
		categories = append(categories, dbCategory.ConvertToDomain())
	}

	return categories, nil
}

func (service *AdService) CreateUserAdCategories(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserAdCategories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.adRepository.CreateUserAdCategories(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *AdService) IncrementLinkClicks(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "IncrementLinkClicks")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.adRepository.IncrementLinkClicks(ctx, id)
	if err != nil { return err }

	return nil
}

func (service *AdService) GetUsersAdCategories(ctx context.Context, id string) ([]domain.AdCategory, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsersAdCategories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCategories, err := service.adRepository.GetUsersAdCategories(ctx, id)
	if err != nil { return nil, err }

	categories := []domain.AdCategory{}
	for _, dbCategory := range dbCategories{
		categories = append(categories, dbCategory.ConvertToDomain())
	}

	return categories, nil
}

func (service *AdService) UpdateUsersAdCategories(ctx context.Context, id string, categories []domain.AdCategory) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUsersAdCategories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if len(categories) < 2 { return errors.New("cannot have less that 2 ad categories") }

	err := service.adRepository.UpdateUsersAdCategories(ctx, id, categories)
	if err != nil { return err }

	return nil
}