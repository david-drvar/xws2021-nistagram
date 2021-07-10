package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
	"time"
)

type CampaignService struct {
	campaignRepository repositories.CampaignRepository
	adService          *AdService
}

func NewCampaignService(db *gorm.DB) (*CampaignService, error) {
	campaignRepository, err := repositories.NewCampaignRepo(db)
	if err != nil {
		return nil, err
	}

	adService, err := NewAdService(db)
	if err != nil {
		return nil, err
	}

	return &CampaignService{
		campaignRepository,
		adService,
	}, err
}

// Retrieve only a list, not all the posts from the campaign
func (service *CampaignService) GetCampaigns(ctx context.Context, userId string) ([]domain.Campaign, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaigns")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCampaigns, err := service.campaignRepository.GetCampaigns(ctx, userId)
	if err != nil {
		return []domain.Campaign{}, err
	}

	campaigns := []domain.Campaign{}
	for _, dbCampaign := range dbCampaigns {
		category, err := service.adService.GetAdCategory(ctx, dbCampaign.AdCategoryId)
		if err != nil {
			return []domain.Campaign{}, err
		}

		campaigns = append(campaigns, dbCampaign.ConvertToDomain([]domain.Ad{}, category)) //Ads will be retrieved upon click
	}

	return campaigns, nil
}

// Only accessible by Agent, who gets all ads from the campaign, including influencer's ads
func (service *CampaignService) GetCampaign(ctx context.Context, campaignId string) (domain.Campaign, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCampaign, err := service.campaignRepository.GetCampaign(ctx, campaignId)
	if err != nil {
		return domain.Campaign{}, err
	}

	ads, err := service.adService.GetAdsFromCampaign(ctx, campaignId)
	if err != nil {
		return domain.Campaign{}, err
	}

	category, err := service.adService.GetAdCategory(ctx, dbCampaign.AdCategoryId)
	if err != nil {
		return domain.Campaign{}, err
	}

	return dbCampaign.ConvertToDomain(ads, category), nil
}

func (service *CampaignService) CreateCampaign(ctx context.Context, campaign domain.Campaign) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if campaign.Name == "" { return errors.New("no name provided") }
	if campaign.StartDate.Equal(time.Time{}) { return errors.New("no start date provided") }
	if campaign.EndDate.Equal(time.Time{}) { return errors.New("no start date provided") }
	if campaign.Category.Id == "" { return errors.New("no ad category") }
	if campaign.StartDate.After(campaign.EndDate) { return errors.New("start date cannot be after end date") }
	if campaign.IsOneTime && !campaign.StartDate.Equal(campaign.EndDate){ campaign.EndDate = campaign.StartDate.Add(24 * time.Hour) }
	if campaign.StartDate.Before(time.Now()) { return errors.New("you cannot create campaigns in past") }
	if campaign.EndDate.Before(time.Now()) { return errors.New("you cannot create campaigns in past") }
	if len(campaign.Ads) == 0 { return errors.New("no ads provided") }
	if campaign.StartTime < 0 || campaign.StartTime > 23 ||  campaign.StartTime >= campaign.EndTime{ return errors.New("invalid start time") }
	if campaign.EndTime < 0 || campaign.EndTime > 23 ||  campaign.EndTime <= campaign.StartTime{ return errors.New("invalid end time") }
	if campaign.IsOneTime { campaign.EndDate = campaign.StartDate.Add(24 * time.Hour) }

	return service.campaignRepository.CreateCampaign(ctx, campaign)
}

// Updates on !isOneTime campaigns need to be taken in consideration after 24hrs
func (service *CampaignService) UpdateCampaign(ctx context.Context, campaign domain.Campaign) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if campaign.Name == "" { return errors.New("no name provided") }
	if campaign.StartDate.Equal(time.Time{}) { return errors.New("no start date provided") }
	if campaign.EndDate.Equal(time.Time{}) { return errors.New("no end date provided") }
	if campaign.Category.Id == "" { return errors.New("no ad category") }
	if campaign.StartDate.After(campaign.EndDate) { return errors.New("start date cannot be after end date") }
	if campaign.IsOneTime && campaign.StartDate.Before(time.Now()) { return errors.New("you cannot update already started one time campaign") }
	if campaign.EndDate.Before(time.Now()) { return errors.New("you cannot update past campaigns") }
	if campaign.StartDate.Before(time.Now()) { return errors.New("you cannot update ongoing campaigns") }
	if campaign.StartTime < 0 || campaign.StartTime > 23 ||  campaign.StartTime >= campaign.EndTime{ return errors.New("invalid start time") }
	if campaign.EndTime < 0 || campaign.EndTime > 23 ||  campaign.EndTime <= campaign.StartTime{ return errors.New("invalid end time") }
	if campaign.IsOneTime { campaign.EndDate = campaign.StartDate.Add(24 * time.Hour) }

	return service.campaignRepository.UpdateCampaign(ctx, campaign)
}
func (service *CampaignService) DeleteCampaign(ctx context.Context, campaignId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.campaignRepository.DeleteCampaign(ctx, campaignId)
}
func (service *CampaignService) GetOngoingCampaignsAds(ctx context.Context, userIds []string, userId string, campaignType model.PostType) ([]domain.Ad, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetOngoingCampaignsAds")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	// Take category, exposure dates and user into consideration
	dbCampaigns, err := service.campaignRepository.GetOngoingCampaigns(ctx)
	if err != nil {
		return nil, err
	}

	userAdCategories, err := service.adService.GetUserAdCategories(ctx, userId)
	if err != nil {
		return nil, err
	}

	ads := []domain.Ad{}
	for _, dbCampaign := range dbCampaigns {
		if campaignType.String() != dbCampaign.Type {
			continue
		}

		appliesToUser := false
		if len(userAdCategories) == 0 {
			appliesToUser = true
		} // non-registered users will get all the ads
		for _, category := range userAdCategories {
			if category.Id == dbCampaign.AdCategoryId {
				appliesToUser = true
				break
			}
		}
		if !appliesToUser {
			continue
		}

		campaignAds, err := service.adService.GetAdsFromCampaign(ctx, dbCampaign.Id)
		if err != nil {
			return []domain.Ad{}, err
		}

		for _, ad := range campaignAds {
			canSee := false
			// Checking if user can see posts from this ad's creator
			for _, id := range userIds {
				if id == ad.Post.UserId {
					canSee = true
					break
				}
			}
			if canSee {
				ads = append(ads, ad)
			}
		}

		err = service.campaignRepository.ChangePlacementsNum(ctx, dbCampaign.Id, len(campaignAds))
		if err != nil {
			return []domain.Ad{}, err
		}
	}

	return ads, nil
}
func (service *CampaignService) GetCampaignStatistics(ctx context.Context, agentId string, campaignId string) (domain.CampaignStats, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaignStatistics")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	campaign, err := service.campaignRepository.GetCampaign(ctx, campaignId)
	if err != nil {
		return domain.CampaignStats{}, err
	}

	if campaign.AgentId != agentId {
		return domain.CampaignStats{}, errors.New("you cannot preview stats for other agent's campaign")
	}

	ads, err := service.adService.GetAdsFromCampaign(ctx, campaignId)
	if err != nil {
		return domain.CampaignStats{}, err
	}

	influencers, err := service.campaignRepository.GetCampaignInfluencers(ctx, campaignId, campaign.Type)
	if err != nil {
		return domain.CampaignStats{}, err
	}

	category, err := service.adService.GetAdCategory(ctx, campaign.AdCategoryId)
	if err != nil {
		return domain.CampaignStats{}, err
	}

	stats := domain.CampaignStats{
		Id:          campaignId,
		Name:        campaign.Name,
		IsOneTime:   campaign.IsOneTime,
		StartDate:   campaign.StartDate,
		EndDate:     campaign.EndDate,
		StartTime:   campaign.StartTime,
		EndTime:     campaign.EndTime,
		Placements:  campaign.Placements,
		Category:    category.Name,
		Type:        campaign.Type,
		Influencers: []domain.InfluencerStats{},
	}

	for _, influencerId := range influencers {
		username, err := grpc_common.GetUsernameById(ctx, influencerId)
		if err != nil {
			return domain.CampaignStats{}, err
		}

		influencerStats := domain.InfluencerStats{Id: influencerId, Username: username, Ads: []domain.AdStats{}}
		// Calculate stats for all influencer's ads
		for _, ad := range ads {
			if ad.Post.UserId != influencerId {
				continue
			}

			mediaContent := []string{}
			for _, media := range ad.Post.Media {
				mediaContent = append(mediaContent, media.Content)
			}

			hashtags := []string{}
			for _, hashtag := range ad.Post.Hashtags {
				hashtags = append(hashtags, hashtag.Text)
			}

			influencerStats.Ads = append(influencerStats.Ads, domain.AdStats{
				Id:       ad.Id,
				Media:    mediaContent,
				Type:     ad.Post.Type.String(),
				Hashtags: hashtags,
				Location: ad.Post.Location,
				Likes:    len(ad.Post.Likes),
				Dislikes: len(ad.Post.Dislikes),
				Comments: len(ad.Post.Comments),
				Clicks:   ad.LinkClicks,
			})
		}

		// Calculate influencer's global stats (e.g. total number of likes, dislikes etc)
		for _, ad := range influencerStats.Ads {
			influencerStats.TotalLikes += ad.Likes
			influencerStats.TotalDislikes += ad.Dislikes
			influencerStats.TotalComments += ad.Comments
			influencerStats.TotalClicks += ad.Clicks
		}

		stats.Influencers = append(stats.Influencers, influencerStats)
		stats.Likes += influencerStats.TotalLikes
		stats.Dislikes += influencerStats.TotalDislikes
		stats.Comments += influencerStats.TotalComments
		stats.Clicks += influencerStats.TotalClicks
	}

	return stats, nil

}
func (service *CampaignService) UpdateCampaignRequest(ctx context.Context, request *domain.CampaignInfluencerRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateCampaignRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.campaignRepository.UpdateCampaignRequest(ctx, request)
	if err != nil {
		return err
	}

	if request.Status == "Accepted" {
		campaign, err := service.GetCampaign(ctx, request.CampaignId)
		if err != nil {
			return err
		}
		for _, ad := range campaign.Ads {
			ad.CampaignId = campaign.Id
			ad.Post.CreatedAt = request.PostAt
			ad.Post.UserId = request.InfluencerId
			ad.Post.IsAd = true
			err := service.adService.CreateAd(ctx, ad)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
func (service *CampaignService) GetCampaignRequestsByAgent(ctx context.Context, agentId string) ([]domain.CampaignInfluencerRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaignRequestsByAgent")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.campaignRepository.GetCampaignRequestsByAgent(ctx, agentId)
}
func (service *CampaignService) CreateCampaignRequest(ctx context.Context, request *domain.CampaignInfluencerRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCampaignRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	campaign, err := service.GetCampaign(ctx, request.CampaignId)
	if err != nil {
		return err
	}
	if !campaign.EndDate.Equal(campaign.StartDate) && (campaign.StartDate.After(request.PostAt) || campaign.EndDate.Before(request.PostAt)) {
		return errors.New("can only hire influencer for the campaingn during campaign duration")
	}
	if campaign.EndDate.Equal(campaign.StartDate) && campaign.StartDate.Before(request.PostAt) {
		return errors.New("for one-time campaign, influencer can only post before the campaign date")

	}

	request, err = service.campaignRepository.CreateCampaignRequest(ctx, request)
	if err != nil {
		return err
	}

	return grpc_common.CreateNotification(ctx, request.InfluencerId, request.AgentId, "Campaign", request.CampaignId)
}
