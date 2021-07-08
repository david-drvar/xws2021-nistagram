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


func (service *AdService) GetAds(ctx context.Context, userId string) ([]domain.Ad, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAds")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	// TODO Check if current user has blocked/muted this user's Ads
	dbAds, err := service.adRepository.GetAds(ctx, userId)
	if err != nil { return []domain.Ad{}, err }

	ads := []domain.Ad{}
	campaignIds := []string{}
	for _, dbAd := range dbAds {
		if dbAd.Type == model.TypePost.String(){
			post, err := service.postService.GetPostById(ctx, dbAd.PostId)
			if err != nil {
				rollback := service.rollbackPlacementNums(ctx, campaignIds)
				if rollback != nil { return []domain.Ad{}, rollback }
				return []domain.Ad{}, err
			}
			ads = append(ads, dbAd.ConvertToDomain(post.Comments, post.Likes, post.Dislikes, post.Objava))
		}else if dbAd.Type == model.TypeStory.String(){
			story, err := service.storyService.GetStoryById(ctx, dbAd.PostId)
			if err != nil {
				rollback := service.rollbackPlacementNums(ctx, campaignIds)
				if rollback != nil { return []domain.Ad{}, rollback }
				return []domain.Ad{}, err
			}
			ads = append(ads, dbAd.ConvertToDomain([]domain.Comment{}, []domain.Like{}, []domain.Like{}, story.Objava))
		}else{
			continue
		}

		campaignIds = append(campaignIds, dbAd.CampaignId)
		// Increment placement num
		err = service.campaignRepository.ChangePlacementsNum(ctx, dbAd.CampaignId, true)
		if err != nil {
			rollback := service.rollbackPlacementNums(ctx, campaignIds)
			if rollback != nil { return []domain.Ad{}, rollback }
			return []domain.Ad{}, err
		}
	}

	return ads, nil
}

func (service *AdService) CreateAd(ctx context.Context, ad domain.Ad) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateAd")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.adRepository.CreateAd(ctx, ad)
}

func (service *AdService) rollbackPlacementNums(ctx context.Context, campaignIds []string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateAd")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	for _, campaignId := range campaignIds{
		err := service.campaignRepository.ChangePlacementsNum(ctx, campaignId, false)
		if err != nil { return errors.New("could not rollback placement numbers, stopped at " + campaignId) }
	}

	return nil
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