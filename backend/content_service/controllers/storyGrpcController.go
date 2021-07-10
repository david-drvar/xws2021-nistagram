package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type StoryGrpcController struct {
	campaignService *services.CampaignService
	service    		*services.StoryService
	jwtManager 		*common.JWTManager
	logger	   		*logger.Logger
}

func NewStoryController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger) (*StoryGrpcController, error) {
	service, err := services.NewStoryService(db)
	if err != nil {
		return nil, err
	}

	campaignService, err := services.NewCampaignService(db)
	if err != nil { return nil, err }

	return &StoryGrpcController{
		campaignService,
		service,
		jwtManager,
		logger,
	}, nil
}

func (c *StoryGrpcController) GetStoriesForUser(ctx context.Context, in *protopb.RequestId) (*protopb.StoriesArray, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoriesForUser")
	defer span.Finish()
	claims, _ := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	isCloseFriends := false
	if claims.UserId == "" {
		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
		if err != nil {
			return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
		}
		if !isPublic{
			return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, "this user is private")
		}
	}else{
		if in.Id != claims.UserId{
			followConnection, err := grpc_common.CheckFollowInteraction(ctx, in.Id, claims.UserId)
			if err != nil {
				return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
			}

			isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
			if err != nil {
				return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
			}

			isBlocked, err := grpc_common.CheckIfBlocked(ctx, in.Id, claims.UserId)
			if err != nil {
				return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
			}

			// If used is blocked or his profile is private and did not approve your request
			if isBlocked || (!isPublic && !followConnection.IsApprovedRequest ) {
				return &protopb.StoriesArray{}, nil
			}

			isCloseFriends = followConnection.IsCloseFriends
		}else{
			isCloseFriends = true
		}
	}
	stories, err := c.service.GetStoriesForUser(ctx, in.Id, isCloseFriends)
	if err != nil{
		return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responseStories := domain.ConvertMultipleStoriesToGrpc(stories)

	return &protopb.StoriesArray{
		Stories: responseStories,
	}, nil
}

func (c *StoryGrpcController) GetMyStories(ctx context.Context, in *protopb.RequestId) (*protopb.StoriesArray, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetMyStories")
	defer span.Finish()
	claims, _ := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if claims.UserId == "" {
		return &protopb.StoriesArray{}, status.Errorf(codes.Unauthenticated, "you do not have your stories")
	}else if in.Id != claims.UserId{
		return &protopb.StoriesArray{}, status.Errorf(codes.Unauthenticated, "cannot access other person's archive")
	}

	stories, err := c.service.GetMyStories(ctx, in.Id)
	if err != nil{
		return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responseStories := domain.ConvertMultipleStoriesToGrpc(stories)

	return &protopb.StoriesArray{
		Stories: responseStories,
	}, nil
}

func (c *StoryGrpcController) GetAllStories(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.StoriesHome, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllStories")
	defer span.Finish()
	claims, _ := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	allStories := domain.StoriesHome{}
	userIds := []string{}	// all user ids that user can access, regardless if they are close friends or not, used for ads
	if claims.UserId == "" {
		publicUserIds, err := grpc_common.GetPublicUsers(ctx)
		if err != nil { return &protopb.StoriesHome{}, err }

		stories, err := c.service.GetAllHomeStories(ctx, publicUserIds, false)
		if err != nil { return &protopb.StoriesHome{}, err }

		allStories.Stories = stories.Stories
		userIds = publicUserIds
	}else {
		publicUserIds, err := grpc_common.GetHomepageUsers(ctx, claims.UserId)
		if err != nil {
			return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
		}

		closeFriends, err := grpc_common.GetCloseFriendsReversed(ctx, claims.UserId)
		nonCloseFriends := []string{}
		for _, userId := range publicUserIds {
			found := false
			for _, closeFriends := range closeFriends {
				if closeFriends == userId {
					found = true
					break
				}
			}
			// My stories will be counted as closefriends+nonclosefriends in the loop below
			if !found && userId != claims.UserId { nonCloseFriends = append(nonCloseFriends, userId) }
			userIds = append(userIds, userId)
		}

		nonCloseFriendsStories, err := c.service.GetAllHomeStories(ctx, nonCloseFriends, false)
		if err != nil {
			return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
		}

		allStories.Stories = nonCloseFriendsStories.Stories

		closeFriends = append(closeFriends, claims.UserId) // Get my close friends stories too

		if len(closeFriends) > 0 {
			closeFriendsStories, err := c.service.GetAllHomeStories(ctx, closeFriends, true)
			if err != nil {
				return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
			}
			for _, storyHome := range closeFriendsStories.Stories {
				allStories.Stories = append(allStories.Stories, storyHome)
			}
		}
	}

	ads, err := c.campaignService.GetOngoingCampaignsAds(ctx, userIds, claims.UserId, model.TypeStory)
	if err != nil { return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error()) }

	responseAds := []*protopb.StoryAdHome{}
	// Get usernames
	for index, story := range allStories.Stories {
		username, err := grpc_common.GetUsernameById(ctx, story.UserId)
		if err != nil {
			return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
		}
		allStories.Stories[index].Username = username

		photo, err := grpc_common.GetPhotoById(ctx, story.UserId)
		if err != nil {
			return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
		}
		allStories.Stories[index].UserPhoto = photo

		// If owner of the ad doesnt have any other stories, we will have to retrieve his info
		// just to display that ad as a standalone story. With this, we can know which story ads will have
		// their owners, so that we do not retrieve user data for them.
		for _, ad := range ads {
			if story.UserId == ad.Post.UserId{
				storyAdHome := &protopb.StoryAdHome{Ad: ad.ConvertToGrpc()}
				storyAdHome.OwnerHasStories = true
				responseAds = append(responseAds, storyAdHome)
				break
			}
		}
	}

	// Fill responseAds with remaining stories that do not have their users
	// and which we will process on frontend
	for _, ad := range ads{
		found := false
		for _, responseAd := range responseAds {
			if responseAd.Ad.Id == ad.Id {
				found = true
				break
			}
		}
		if found { continue }
		responseAds = append(responseAds, &protopb.StoryAdHome{
			OwnerHasStories: false,
			Ad: ad.ConvertToGrpc(),
		})
	}

	responseStories := allStories.ConvertToGrpc(responseAds)
	return responseStories, nil
}

func (c *StoryGrpcController) CreateStory(ctx context.Context, in *protopb.Story) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateStory")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt by " + claims.UserId, logger.Info)

	if err != nil {
		c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt failed by " + claims.UserId + ", invalid JWT", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == "" {
		c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt failed by " + claims.UserId + ", invalid JWT", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id is provided")
	}else if in.UserId != claims.UserId {
		c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt failed by " + claims.UserId + ", tried to create story for " + in.UserId, logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot create stories for other people")
	}

	var story *domain.Story
	story = story.ConvertFromGrpc(in)

	for _, media := range story.Media {
		for _, tag := range media.Tags {
			following, err := grpc_common.CheckFollowInteraction(ctx, tag.UserId, story.UserId)
			if err != nil {
				c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt failed by " + claims.UserId + ", cannot tag "+tag.UserId, logger.Error)
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
			}

			isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.UserId)
			if err != nil {
				c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt failed by " + claims.UserId + ", cannot tag "+tag.UserId, logger.Error)
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
			}

			isBlocked, err := grpc_common.CheckIfBlocked(ctx, in.UserId, claims.UserId)
			if err != nil {
				c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt failed by " + claims.UserId + ", cannot tag "+tag.UserId, logger.Error)
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
			}

			// If used is blocked or his profile is private and did not approve your request
			if isBlocked || (!isPublic && !following.IsApprovedRequest) {
				c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt failed by " + claims.UserId + ", cannot tag "+tag.UserId, logger.Error)
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
			}

			username, err := grpc_common.GetUsernameById(ctx, tag.UserId)
			if username == "" || err != nil {
				c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt failed by " + claims.UserId+ ", cannot tag "+tag.UserId, logger.Error)
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
			}
		}
	}

	err = c.service.CreateStory(ctx, story)
	if err != nil {
		c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt failed by " + claims.UserId + ", server error", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create story")
	}

	c.logger.ToStdoutAndFile("CreateStory", "Story creation attempt successful by " + claims.UserId, logger.Info)
	return &protopb.EmptyResponseContent{}, nil
}

func (c *StoryGrpcController) GetStoryById(ctx context.Context, in *protopb.RequestId) (*protopb.Story, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoryById")
	defer span.Finish()
	claims, _ := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if in.Id == "" {
		return &protopb.Story{}, status.Errorf(codes.Unknown, "cannot retrieve non-existing stories")
	}

	story, err := c.service.GetStoryById(ctx, in.Id)
	if err != nil { return &protopb.Story{}, status.Errorf(codes.Unknown, err.Error()) }

	if claims.UserId == "" {
		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, story.UserId)
		if err != nil { return &protopb.Story{}, status.Errorf(codes.Unknown, err.Error()) }

		if !isPublic { return &protopb.Story{}, status.Errorf(codes.Unknown, "this story is not public") }
	}else if story.UserId != claims.UserId && claims.Role!="Admin" {
		following, err := grpc_common.CheckFollowInteraction(ctx, story.UserId, claims.UserId)
		if err != nil { return &protopb.Story{}, status.Errorf(codes.Unknown, err.Error()) }

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, story.UserId)
		if err != nil { return &protopb.Story{}, status.Errorf(codes.Unknown, err.Error()) }

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, story.UserId, claims.UserId)
		if err != nil { return &protopb.Story{}, status.Errorf(codes.Unknown, err.Error()) }

		if (!following.IsApprovedRequest && !isPublic) || isBlocked || (story.IsCloseFriends && !following.IsCloseFriends) {
			return &protopb.Story{}, status.Errorf(codes.PermissionDenied, "cannot retrieve this story")
		}
	}

	grpcStory := story.ConvertToGrpc()

	return grpcStory, nil
}

func (c *StoryGrpcController) RemoveStory(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveStory")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	c.logger.ToStdoutAndFile("RemoveStory", "Story removal attempt by " + claims.UserId, logger.Info)

	if err != nil {
		c.logger.ToStdoutAndFile("RemoveStory", "Story removal attempt failed by " + claims.UserId + ", invalid JWT", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == ""{
		c.logger.ToStdoutAndFile("RemoveStory", "Story removal attempt failed by " + claims.UserId + ", invalid JWT", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove other people's posts")
	}else if in.Id == "" {
		c.logger.ToStdoutAndFile("RemoveStory", "Story removal attempt failed by " + claims.UserId + ", no story id provided", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove non-existing posts")
	}

	err = c.service.RemoveStory(ctx, in.Id, claims.UserId)

	if err != nil{
		c.logger.ToStdoutAndFile("RemoveStory", "Story removal attempt failed by " + claims.UserId + ", server error", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}

	c.logger.ToStdoutAndFile("RemoveStory", "Story removal attempt successful by " + claims.UserId, logger.Info)
	return &protopb.EmptyResponseContent{}, nil
}