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

type CampaignRepository interface {
	GetCampaign(context.Context, string) (persistence.Campaign, error)
	GetCampaigns(context.Context, string) ([]persistence.Campaign, error)
	CreateCampaign(context.Context, domain.Campaign) error
	UpdateCampaign(context.Context, domain.Campaign) error
	DeleteCampaign(context.Context, string) error
	checkCampaignChanges(context.Context, string) error
	ChangePlacementsNum(context.Context, string, int) error
	GetOngoingCampaigns(context.Context) ([]persistence.Campaign, error)
	GetCampaignInfluencers(context.Context, string, string) ([]string, error)
	UpdateCampaignRequest(ctx context.Context, request *domain.CampaignInfluencerRequest) error
	GetCampaignRequestsByAgent(ctx context.Context, id string) ([]domain.CampaignInfluencerRequest, error)
	CreateCampaignRequest(ctx context.Context, request *domain.CampaignInfluencerRequest) (*domain.CampaignInfluencerRequest, error)
}

type campaignRepository struct {
	DB           *gorm.DB
	adRepository AdRepository
}

func NewCampaignRepo(db *gorm.DB) (*campaignRepository, error) {
	if db == nil {
		panic("CampaignRepository not created, gorm.DB is nil")
	}

	adRepository, _ := NewAdRepo(db)

	return &campaignRepository{
		DB:           db,
		adRepository: adRepository,
	}, nil
}

const (
	campaignUpdateTimer = 24 * time.Hour
)

func (repository *campaignRepository) GetCampaign(ctx context.Context, campaignId string) (persistence.Campaign, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var campaign persistence.Campaign
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		// Check latest campaign changes
		err := repository.checkCampaignChanges(ctx, campaignId)
		if err != nil {
			return err
		}

		// Retrieve the campaign
		result := repository.DB.Where("id = ?", campaignId).Find(&campaign)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return persistence.Campaign{}, err
	}

	return campaign, nil
}

func (repository *campaignRepository) GetCampaigns(ctx context.Context, userId string) ([]persistence.Campaign, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaigns")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var campaigns []persistence.Campaign
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		result := repository.DB.Where("agent_id = ?", userId).Find(&campaigns)
		if result.Error != nil {
			return result.Error
		}
		for _, campaign := range campaigns {
			// Check latest campaign changes
			err := repository.checkCampaignChanges(ctx, campaign.Id)
			if err != nil {
				return err
			}
		}

		return nil
	})

	result := repository.DB.Where("agent_id = ?", userId).Find(&campaigns)
	if result.Error != nil {
		return []persistence.Campaign{}, result.Error
	}

	return campaigns, err
}

func (repository *campaignRepository) CreateCampaign(ctx context.Context, campaign domain.Campaign) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaigns")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		var dbCampaign persistence.Campaign
		dbCampaign = dbCampaign.ConvertToPersistence(campaign)
		dbCampaign.Id = uuid.NewV4().String()
		dbCampaign.LastUpdated = time.Now()
		if dbCampaign.EndTime < dbCampaign.StartTime {
			dbCampaign.EndTime = campaign.StartTime
		}
		if dbCampaign.StartTime < 0 || dbCampaign.StartTime > 23 {
			dbCampaign.StartTime = 0
		}
		if dbCampaign.EndTime < 0 || dbCampaign.EndTime > 23 {
			dbCampaign.EndTime = 0
		}
		result := repository.DB.Save(dbCampaign)
		if result.Error != nil {
			return result.Error
		}

		for _, ad := range campaign.Ads {
			ad.CampaignId = dbCampaign.Id
			ad.Post.CreatedAt = dbCampaign.StartDate
			ad.Post.UserId = campaign.AgentId
			ad.Post.IsAd = true
			err := repository.adRepository.CreateAd(ctx, ad)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

// Create new campaign change if there is no non-applied campaign change
// If there is non-applied change, update it and reset the 24h timer
func (repository *campaignRepository) UpdateCampaign(ctx context.Context, campaign domain.Campaign) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	tomorrow := time.Now().Add(campaignUpdateTimer)
	if campaign.IsOneTime && campaign.StartDate.After(tomorrow) {
		// Do not allow updates on OneTime Campaigns before they start
		return errors.New("cannot update one time campaigns before they start")
	}

	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		// Apply non-applied changes if there are any
		err := repository.checkCampaignChanges(ctx, campaign.Id)
		if err != nil {
			return err
		}

		campaignChanges := persistence.CampaignChanges{}
		result := repository.DB.Model(&campaignChanges).Where("campaign_id = ? AND applied = false", campaign.Id).First(&campaignChanges)
		if result.Error != nil && result.Error.Error() != "record not found" {
			return result.Error
		}

		checkedCampaignChanges := persistence.CampaignChanges{}

		// There is an existing CampaignChanges, update it
		if campaignChanges.CampaignId != "" {
			// Creating update struct
			willUpdate := false
			if campaign.StartDate.After(tomorrow) {
				checkedCampaignChanges.StartDate = campaign.StartDate
				willUpdate = true
			}
			if campaign.EndDate.After(tomorrow) {
				checkedCampaignChanges.EndDate = campaign.EndDate
				willUpdate = true
			}
			if campaign.Category.Id != "" {
				checkedCampaignChanges.AdCategoryId = campaign.Category.Id
				willUpdate = true
			}
			if campaign.Name != "" {
				checkedCampaignChanges.Name = campaign.Name
				willUpdate = true
			}
			if campaign.StartTime != 0 {
				checkedCampaignChanges.StartTime = campaign.StartTime
				willUpdate = true
			}
			if campaign.EndTime != 0 {
				checkedCampaignChanges.EndTime = campaign.EndTime
				willUpdate = true
			}
			if campaign.EndTime < campaign.StartTime {
				checkedCampaignChanges.EndTime = campaign.StartTime
			}
			if campaign.StartTime < 0 || campaign.StartTime > 23 {
				checkedCampaignChanges.StartTime = 0
			}
			if campaign.EndTime < 0 || campaign.EndTime > 23 {
				checkedCampaignChanges.EndTime = 0
			}

			checkedCampaignChanges.ValidFrom = time.Now().Add(campaignUpdateTimer)

			if campaign.IsOneTime && !campaign.StartDate.Equal(campaign.EndDate) {
				checkedCampaignChanges.EndDate = checkedCampaignChanges.StartDate
			}
			if willUpdate {
				result = repository.DB.Model(&checkedCampaignChanges).Where("campaign_id = ?", campaign.Id).Updates(checkedCampaignChanges)
				if result.Error != nil {
					return result.Error
				}
			}
		} else {
			// There is no existing CampaignChanges, create a new one
			if campaign.StartDate.After(tomorrow) {
				checkedCampaignChanges.StartDate = campaign.StartDate
			}
			if campaign.EndDate.After(tomorrow) {
				checkedCampaignChanges.EndDate = campaign.EndDate
			}
			if campaign.Category.Id != "" {
				checkedCampaignChanges.AdCategoryId = campaign.Category.Id
			}
			if campaign.IsOneTime && !campaign.StartDate.Equal(campaign.EndDate) {
				checkedCampaignChanges.EndDate = checkedCampaignChanges.StartDate
			}
			if campaign.Name != "" {
				checkedCampaignChanges.Name = campaign.Name
			}
			if campaign.StartTime != 0 {
				checkedCampaignChanges.StartTime = campaign.StartTime
			}
			if campaign.EndTime != 0 {
				checkedCampaignChanges.EndTime = campaign.EndTime
			}
			if campaign.EndTime < campaign.StartTime {
				checkedCampaignChanges.EndTime = campaign.StartTime
			}
			if campaign.StartTime < 0 || campaign.StartTime > 23 {
				checkedCampaignChanges.StartTime = 0
			}
			if campaign.EndTime < 0 || campaign.EndTime > 23 {
				checkedCampaignChanges.EndTime = 0
			}
			checkedCampaignChanges.CampaignId = campaign.Id
			checkedCampaignChanges.Applied = false
			checkedCampaignChanges.ValidFrom = time.Now().Add(campaignUpdateTimer)
			checkedCampaignChanges.Id = uuid.NewV4().String()

			result = repository.DB.Save(checkedCampaignChanges)
			if result.Error != nil {
				return result.Error
			}
		}

		return nil
	})

	return err
}

func (repository *campaignRepository) DeleteCampaign(ctx context.Context, campaignId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		result := repository.DB.Where("id = ?", campaignId).Delete(&persistence.Campaign{})
		if result.Error != nil {
			return result.Error
		}

		result = repository.DB.Where("campaign_id = ?", campaignId).Delete(&persistence.CampaignChanges{})
		if result.Error != nil {
			return result.Error
		}

		err := repository.adRepository.DeleteAdsFromCampaign(ctx, campaignId)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (repository *campaignRepository) checkCampaignChanges(ctx context.Context, campaignId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "checkCampaignChanges")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	// Get the latest CampaignChanges
	var campaignChanges persistence.CampaignChanges
	result := repository.DB.Model(&campaignChanges).Where("campaign_id = ? AND applied = false", campaignId).First(&campaignChanges)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return result.Error
	}

	// End early because there are no campaign changes
	if campaignChanges.Id == "" {
		return nil
	}

	// If there have been any changes that are ready to be applied, apply them to main table
	if campaignChanges.ValidFrom.Before(time.Now()) { // If the ValidFrom time has passed, apply changes
		err := repository.DB.Transaction(func(tx *gorm.DB) error {
			tomorrow := time.Now().Add(campaignUpdateTimer)
			updateChanges := persistence.Campaign{
				Id:           campaignChanges.CampaignId,
				AdCategoryId: campaignChanges.AdCategoryId,
				StartDate:    time.Time{}, // Zero values won't be updated
				EndDate:      time.Time{},
				LastUpdated:  time.Now(),
			}

			if campaignChanges.Name != "" {
				updateChanges.Name = campaignChanges.Name
			}
			if campaignChanges.AdCategoryId != "" {
				updateChanges.AdCategoryId = campaignChanges.AdCategoryId
			}

			// Allow dates update up to 24hrs before campaign starts
			if campaignChanges.StartDate.After(tomorrow) {
				updateChanges.StartDate = campaignChanges.StartDate
			}
			if campaignChanges.EndDate.After(tomorrow) {
				updateChanges.EndDate = campaignChanges.EndDate
			}
			if campaignChanges.StartTime != 0 {
				updateChanges.StartTime = campaignChanges.StartTime
			}
			if campaignChanges.EndTime != 0 {
				updateChanges.EndTime = campaignChanges.EndTime
			}
			if campaignChanges.StartTime < 0 || campaignChanges.StartTime > 23 {
				updateChanges.StartTime = 0
			}
			if campaignChanges.EndTime < 0 || campaignChanges.EndTime > 23 {
				updateChanges.EndTime = 0
			}

			// Update main table with new values
			result = repository.DB.Model(&persistence.Campaign{}).Where("id = ?", campaignId).Updates(updateChanges)
			if result.Error != nil {
				return result.Error
			}

			// Set Applied to true and confirm that the update has been made
			result = repository.DB.Model(&campaignChanges).Where("id = ?", campaignChanges.Id).Update("applied", true)
			if result.Error != nil {
				return result.Error
			}

			// Update post/story date as well
			if campaignChanges.StartDate.After(tomorrow) {
				var campaign persistence.Campaign
				result = repository.DB.Model(&campaign).Where("id = ?", campaignId).Find(&campaign)
				if result.Error != nil {
					return result.Error
				}

				err := repository.adRepository.UpdateCampaignAdDate(ctx, campaignId, campaign.Type, updateChanges.StartDate)
				if err != nil {
					return err
				}
			}

			return nil
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (repository *campaignRepository) ChangePlacementsNum(ctx context.Context, campaignId string, number int) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "ChangePlacementsNum")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var campaign persistence.Campaign
	result := repository.DB.Where("id = ?", campaignId).First(&campaign)
	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	}

	campaign.Placements += number

	result = repository.DB.Model(&campaign).Update("placements", campaign.Placements).Where("id = ?", campaignId)

	return result.Error
}

func (repository *campaignRepository) GetOngoingCampaigns(ctx context.Context) ([]persistence.Campaign, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetOngoingCampaigns")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		var campaigns []persistence.Campaign
		now := time.Now()
		result := repository.DB.Where("start_date <= ? AND end_date >= ?", now, now).Find(&campaigns)
		if result.Error != nil {
			return result.Error
		}

		for _, campaign := range campaigns {
			// Check latest campaign changes
			err := repository.checkCampaignChanges(ctx, campaign.Id)
			if err != nil {
				return err
			}
		}

		return nil
	})

	var campaigns []persistence.Campaign
	now := time.Now()
	currentHour, _, _ := now.Clock()
	result := repository.DB.
		Where("start_date <= ? AND end_date >= ? AND start_time <= ? AND end_time >= ?", now, now, currentHour, currentHour).
		Find(&campaigns)
	if result.Error != nil {
		return []persistence.Campaign{}, result.Error
	}

	return campaigns, err
}

func (repository *campaignRepository) GetCampaignInfluencers(ctx context.Context, id string, campaignType string) ([]string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaignInfluencers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	influencers := []string{}
	if campaignType == model.TypePost.String() {
		result := repository.DB.Model(&persistence.Post{}).
			Joins("left join ads on posts.id = ads.post_id").
			Where("ads.campaign_id = ?", id).
			Pluck("posts.user_id", &influencers)
		if result.Error != nil {
			return nil, result.Error
		}
	} else if campaignType == model.TypeStory.String() {
		result := repository.DB.Model(&persistence.Story{}).
			Joins("left join ads on stories.id = ads.post_id").
			Where("ads.campaign_id = ?", id).
			Pluck("stories.user_id", &influencers)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	return influencers, nil
}
func (repository *campaignRepository) CreateCampaignRequest(ctx context.Context, request *domain.CampaignInfluencerRequest) (*domain.CampaignInfluencerRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCampaignRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var requests []persistence.CampaignInfluencerRequest

	db := repository.DB.Model(&request).Where("agent_id = ? AND influencer_id=? AND campaign_id=?", request.AgentId, request.InfluencerId, request.CampaignId).Find(&requests)

	if db.Error != nil {
		return nil, db.Error
	} else if db.RowsAffected != 0 {
		return nil, errors.New("cannot create campaign")
	}

	//if request.PostAt.Before(time.Now()) {
	//	return nil, errors.New("cannot create campaign")
	//}

	request.Id = uuid.NewV4().String()
	db = repository.DB.Create(&request)

	return request, db.Error
}

func (repository *campaignRepository) UpdateCampaignRequest(ctx context.Context, request *domain.CampaignInfluencerRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateCampaignRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var dbRequest persistence.CampaignInfluencerRequest
	db := repository.DB.Model(&request).Where("agent_id = ? AND influencer_id=? AND campaign_id=?", request.AgentId, request.InfluencerId, request.CampaignId).Find(&dbRequest)
	if db.Error != nil {
		return db.Error
	}
	if dbRequest.Status != "Pending" {
		return errors.New("request can only be updated once")
	}

	dbRequest.Status = request.Status
	return repository.DB.Model(&request).Where("id = ?", dbRequest.Id).Updates(dbRequest).Error
}

func (repository *campaignRepository) GetCampaignRequestsByAgent(ctx context.Context, agentId string) ([]domain.CampaignInfluencerRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaignRequestsByAgent")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var requests []domain.CampaignInfluencerRequest
	result := repository.DB.Where("agent_id = ?", agentId).Find(&requests)
	return requests, result.Error
}
