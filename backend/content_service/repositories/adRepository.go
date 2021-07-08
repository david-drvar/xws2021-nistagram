package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type AdRepository interface {
	GetAd(context.Context, string) 	 	 (persistence.Ad, error)
	GetAds(context.Context, string)      ([]persistence.Ad, error)
	CreateAd(context.Context, domain.Ad) error
	UpdateAd(context.Context, domain.Ad) error
	DeleteAdsFromCampaign(context.Context, string) error
	GetAdsFromCampaign(context.Context, string) ([]persistence.Ad, error)

	GetAdCategories(context.Context) ([]persistence.AdCategory, error)
	CreateAdCategory(context.Context, domain.AdCategory) error
	GetAdCategory(context.Context, string) (persistence.AdCategory, error)

	UpdateCampaignAdDate(context.Context, string, string, time.Time) error
}

type adRepository struct {
	DB 				*gorm.DB
	postRepository  PostRepository
	storyRepository StoryRepository
}

func NewAdRepo(db *gorm.DB) (*adRepository, error) {
	if db == nil {
		panic("AdRepository not created, gorm.DB is nil")
	}

	postRepository, _ := NewPostRepo(db)
	storyRepository, _ := NewStoryRepo(db)

	return &adRepository{
		DB: db,
		postRepository: postRepository,
		storyRepository: storyRepository,
	}, nil
}

func (repository *adRepository) GetAd(ctx context.Context, adId string) (persistence.Ad, error){

	return persistence.Ad{}, nil
}

func (repository *adRepository) GetAds(ctx context.Context, userId string) ([]persistence.Ad, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAds")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var ads []persistence.Ad
	if userId == "" {
		result := repository.DB.Raw("SELECT * FROM ads ORDER BY random() LIMIT 3").Scan(&ads)
		if result.Error != nil { return []persistence.Ad{}, result.Error }
	}else{
		// TODO Determine retrieving mechanism
	}

	return ads, nil
}

func (repository *adRepository) CreateAd(ctx context.Context, ad domain.Ad) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateAd")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func (tx *gorm.DB) error {
		var dbAd persistence.Ad
		dbAd = dbAd.ConvertToPersistence(ad)
		dbAd.Id = uuid.NewV4().String()

		if ad.Post.Type == model.TypePost{
			post, err := repository.postRepository.CreatePost(ctx, &ad.Post)
			if err != nil { return err }
			dbAd.PostId = post.Id
		}else if ad.Post.Type == model.TypeStory {
			story, err := repository.storyRepository.CreateStory(ctx, ad.CreateStoryAd())
			if err != nil { return err }
			dbAd.PostId = story.Id
		}

		result := repository.DB.Create(dbAd)
		if result.Error != nil || result.RowsAffected != 1 {
			return result.Error
		}

		return nil
	})

	return err
}

// Used for updating link clicks only
func (repository *adRepository) UpdateAd(ctx context.Context, ad domain.Ad) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateAd")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repository.DB.Model(&ad).Where("id = ?", ad.Id).Update("link_clicks", ad.LinkClicks)
	if result.Error != nil || result.RowsAffected != 1 { return result.Error }

	return nil
}

func (repository *adRepository) DeleteAdsFromCampaign(ctx context.Context, campaignId string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteAdsFromCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func (tx *gorm.DB) error {
		ads := []persistence.Ad{}
		result := repository.DB.Model(&persistence.Ad{}).Where("campaign_id = ?", campaignId).Find(&ads)
		if result.Error != nil { return result.Error }

		err := repository.DB.Transaction(func (tx *gorm.DB) error {
			for _, ad := range ads {
				result = repository.DB.Where("id = ?", ad.Id).Delete(&ad)
				if result.Error != nil || result.RowsAffected != 1 { return result.Error }

				if ad.Type == model.TypePost.String() {
					err := repository.postRepository.RemovePost(ctx, ad.PostId, "")
					if err != nil { return err }
				}else if ad.Type == model.TypeStory.String(){
					err := repository.storyRepository.RemoveStory(ctx, ad.PostId, "")
					if err != nil { return err }
				}
			}

			return nil
		})

		return err
	})

	return err
}

func (repository *adRepository) GetAdsFromCampaign(ctx context.Context, id string) ([]persistence.Ad, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdsFromCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var ads []persistence.Ad
	result := repository.DB.Model(&persistence.Ad{}).Where("campaign_id = ?", id).Find(&ads)
	if result.Error != nil { return []persistence.Ad{}, result.Error }

	return ads, nil
}

func (repository *adRepository) GetAdCategories(ctx context.Context) ([]persistence.AdCategory, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdCategories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var categories []persistence.AdCategory
	result := repository.DB.Model(&persistence.AdCategory{}).Find(&categories)
	if result.Error != nil { return categories, result.Error }

	return categories, nil
}

func (repository *adRepository) CreateAdCategory(ctx context.Context, category domain.AdCategory) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateAdCategory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repository.DB.Save(&category)
	if result.Error != nil { return result.Error }

	return nil
}

func (repository *adRepository) GetAdCategory(ctx context.Context, id string) (persistence.AdCategory, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdCategory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var category persistence.AdCategory
	result := repository.DB.Model(&persistence.AdCategory{}).Where("id = ?", id).Find(&category)
	if result.Error != nil { return category, result.Error }

	return category, nil
}

func (repository *adRepository) UpdateCampaignAdDate(ctx context.Context, campaignId string, postType string, startDate time.Time) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateAdDates")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	ads, err := repository.GetAdsFromCampaign(ctx, campaignId)
	if err != nil { return err }

	for _, ad := range ads{
		if postType == model.TypePost.String(){
			err := repository.postRepository.UpdateCreatedAt(ctx, ad.PostId, startDate)
			if err != nil { return err }
		}else if postType == model.TypeStory.String(){
			err := repository.storyRepository.UpdateCreatedAt(ctx, ad.PostId, startDate)
			if err != nil { return err }
		}else{
			return errors.New("unknown post type")
		}
	}

	return nil
}