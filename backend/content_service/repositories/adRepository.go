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
	GetAdsFromInfluencer(context.Context, string) ([]persistence.Ad, error)
	UpdateCampaignAdDate(context.Context, string, string, time.Time) error
	IncrementLinkClicks(context.Context, string) error

	GetAdCategories(context.Context) ([]persistence.AdCategory, error)
	GetUserAdCategories(context.Context, string) ([]persistence.AdCategory, error)
	CreateAdCategory(context.Context, domain.AdCategory) error
	GetAdCategory(context.Context, string) (persistence.AdCategory, error)
	CreateUserAdCategories(context.Context, string) error
	GetUsersAdCategories(context.Context, string) ([]persistence.AdCategory, error)
	UpdateUsersAdCategories(context.Context, string, []domain.AdCategory) error
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
	}

	return ads, nil
}

func (repository *adRepository) GetAdsFromInfluencer(ctx context.Context, userId string) ([]persistence.Ad, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAds")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	now := time.Now()
	var postAds []persistence.Ad
	result := repository.DB.Model(&persistence.Ad{}).
		Joins("left join posts on ads.post_id = posts.id").
		Where("posts.user_id = ? AND posts.created_at <= ?", userId, now).
		Find(&postAds)
	if result.Error != nil { return []persistence.Ad{}, result.Error }

	var storyAds []persistence.Ad
	result = repository.DB.Model(&persistence.Ad{}).
		Joins("left join stories on ads.post_id = stories.id").
		Where("stories.user_id = ? AND stories.created_at <= ?", userId, now).
		Find(&storyAds)
	if result.Error != nil { return []persistence.Ad{}, result.Error }

	ads := []persistence.Ad{}
	for _, storyAd := range storyAds{ ads = append(ads, storyAd) }
	for _, postAd := range postAds{ ads = append(ads, postAd) }

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

func (repository *adRepository) GetUserAdCategories(ctx context.Context, userId string) ([]persistence.AdCategory, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdCategories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var categories []persistence.AdCategory
	result := repository.DB.Model(&persistence.AdCategory{}).
		Joins("left join user_ad_categories on user_ad_categories.id_ad_category = ad_categories.id").
		Where("user_ad_categories.user_id = ?", userId).
		Find(&categories)
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

func (repository *adRepository) CreateUserAdCategories(ctx context.Context, id string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserAdCategories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	adCategories := []persistence.AdCategory{}
	result := repository.DB.Find(&adCategories)
	if result.Error != nil { return result.Error }

	repository.DB.Transaction(func (tx *gorm.DB) error {
		for _, category := range adCategories{
			result := repository.DB.Create(persistence.UserAdCategories{
				UserId:       id,
				IdAdCategory: category.Id,
			})
			if result.Error != nil {
				return result.Error
			}
		}

		return nil
	})

	return nil
}

func (repository *adRepository) IncrementLinkClicks(ctx context.Context, id string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "IncrementLinkClicks")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	ad := persistence.Ad{}
	result := repository.DB.Where("id = ? OR post_id = ?", id, id).Find(&ad)
	if result.Error != nil { return result.Error }

	ad.LinkClicks += 1
	// 'OR post_id' part is a quick hotfix, it is not clean.
	result = repository.DB.Model(&persistence.Ad{}).Where("id = ? OR post_id = ?", ad.Id, ad.Id).Update("link_clicks", ad.LinkClicks)
	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	}

	return nil
}

func (repository *adRepository) GetUsersAdCategories(ctx context.Context, id string) ([]persistence.AdCategory, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdCategory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var categories []persistence.AdCategory
	result := repository.DB.Model(&persistence.AdCategory{}).
		Joins("left join user_ad_categories ON user_ad_categories.id_ad_category = ad_categories.id").
		Where("user_ad_categories.user_id = ?", id).Find(&categories)
	if result.Error != nil { return categories, result.Error }

	return categories, nil
}

func (repository *adRepository) UpdateUsersAdCategories(ctx context.Context, id string, categories []domain.AdCategory) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdCategory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var usersCategories []persistence.AdCategory
	// Get user's ad categories
	result := repository.DB.Model(&persistence.AdCategory{}).
		Joins("left join user_ad_categories ON user_ad_categories.id_ad_category = ad_categories.id").
		Where("user_ad_categories.user_id = ?", id).Find(&usersCategories)
	if result.Error != nil { return result.Error }

	err := repository.DB.Transaction(func (tx *gorm.DB) error {
		// First, determine which categories we need to remove from the database
		for _, dbCategory := range usersCategories{
			found := false
			for _, category := range categories{
				if category.Id == dbCategory.Id {
					found = true
					break
				}
			}
			if found { continue } // category is in both in database and in sent categories, skip it

			// If it's not found, we need to remove it
			result := repository.DB.Delete(&persistence.UserAdCategories{UserId: id, IdAdCategory: dbCategory.Id})
			if result.Error != nil { return result.Error }
		}

		// Second, determine which categories we need to add to the database
		for _, category := range categories{
			found := false
			for _, dbCategory := range usersCategories{
				if category.Id == dbCategory.Id {
					found = true
					break
				}
			}
			if found { continue } // category is in both in database and in sent categories, skip it

			// If it's not found, we need to add it
			result := repository.DB.Save(&persistence.UserAdCategories{UserId: id, IdAdCategory: category.Id})
			if result.Error != nil { return result.Error }
		}

		return nil
	})


	return err
}