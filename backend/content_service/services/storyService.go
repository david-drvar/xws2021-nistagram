package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type StoryService struct {
	storyRepository   	repositories.StoryRepository
	mediaRepository   	repositories.MediaRepository
	tagRepository	  	repositories.TagRepository
	hashtagRepository	repositories.HashtagRepository
}

func NewStoryService(db *gorm.DB) (*StoryService, error){
	storyRepository, err := repositories.NewStoryRepo(db)
	if err != nil {
		return nil, err
	}

	mediaRepository, err := repositories.NewMediaRepo(db)
	if err != nil {
		return nil, err
	}

	tagRepository, err := repositories.NewTagRepo(db)
	if err != nil {
		return nil, err
	}

	hashtagRepository, err := repositories.NewHashtagRepo(db)
	if err != nil {
		return nil, err
	}

	return &StoryService{
		storyRepository,
		mediaRepository,
		tagRepository,
		hashtagRepository,
	}, err
}

func (service *StoryService) GetStoriesForUser(ctx context.Context, userId string, isCloseFriend bool) ([]domain.Story, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoriesForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	res, err := grpc_common.CheckIsActive(ctx, userId)
	if err != nil {
		return nil, err
	}else if res == false {
		return nil, errors.New("User is not active!")
	}
	dbStories, err := service.storyRepository.GetUsersStories(ctx, userId, false)
	if err != nil { return []domain.Story{}, err }

	if isCloseFriend{
		closeFriendsStories, err := service.storyRepository.GetUsersStories(ctx, userId, true)
		if err != nil { return []domain.Story{}, err }

		for _, closeFriendStory := range closeFriendsStories {
			dbStories = append(dbStories, closeFriendStory)
		}
	}

	stories := []domain.Story{}
	for _, dbStory := range dbStories{
		story, err := service.retrieveStoryAdditionalData(ctx, dbStory)
		if err != nil { return []domain.Story{}, err }

		stories = append(stories, story)
	}

	return stories, nil
}

func (service *StoryService) GetMyStories(ctx context.Context, userId string) ([]domain.Story, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetMyStories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	res, err := grpc_common.CheckIsActive(ctx, userId)
	if err != nil {
		return nil, err
	}else if res == false {
		return nil, errors.New("User is not active!")
	}

	dbStories, err := service.storyRepository.GetMyStories(ctx, userId)
	if err != nil { return []domain.Story{}, err }

	stories := []domain.Story{}
	for _, dbStory := range dbStories{
		story, err := service.retrieveStoryAdditionalData(ctx, dbStory)
		if err != nil { return []domain.Story{}, err }

		stories = append(stories, story)
	}

	return stories, nil
}

func (service *StoryService) GetAllHomeStories(ctx context.Context, userIds []string, isCloseFriends bool) (domain.StoriesHome, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHomeStories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var userIdsActive []string
	for _, id := range userIds {
		res, err := grpc_common.CheckIsActive(ctx, id)
		if err != nil {
			continue
		}else if res {
			userIdsActive = append(userIdsActive, id)
		}
	}

	dbStories, err := service.storyRepository.GetAllHomeStories(ctx, userIdsActive, false)
	if err != nil { return domain.StoriesHome{}, err }

	if isCloseFriends {
		closeFriendsStories, err := service.storyRepository.GetAllHomeStories(ctx, userIdsActive, true)
		if err != nil { return domain.StoriesHome{}, err }

		for _, closeFriendStory := range closeFriendsStories {
			dbStories = append(dbStories, closeFriendStory)
		}
	}

	storiesHome := domain.StoriesHome{}
	for _, dbStory := range dbStories{
		story, err := service.retrieveStoryAdditionalData(ctx, dbStory)
		if err != nil { return domain.StoriesHome{}, err }

		// If storiesHome are empty (not initialized), create a new entry with first story
		if len(storiesHome.Stories) == 0 {
			storiesHome.Stories = append(storiesHome.Stories, domain.StoryHome{
				UserId:   story.UserId,
				Username: "",
				Stories:  []domain.Story{story},
			})
		}else{
			// If storiesHome is initialized, check where current story should go
			updated := false
			for index, storyHome := range storiesHome.Stories {
				if storyHome.UserId == story.UserId { // If storiesHome already has stories from this user, append story
					storiesHome.Stories[index].Stories = append(storiesHome.Stories[index].Stories, story)
					updated = true
					break
				}
			}
			// If we couldn't find a place for this story, create a new group for stories from this user
			if !updated {
				storiesHome.Stories = append(storiesHome.Stories, domain.StoryHome{
					UserId:   story.UserId,
					Username: "",
					Stories:  []domain.Story{story},
				})
			}
		}
	}

	return storiesHome, nil
}
func (service *StoryService) CreateStory(ctx context.Context, story *domain.Story) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if len(story.Media) == 0 {
		return errors.New("cannot create empty story")
	}

	_, err := service.storyRepository.CreateStory(ctx, story)
	if err != nil {
		return err
	}

	users, err := grpc_common.GetUsersForNotificationEnabled(ctx, story.UserId, "IsStoryNotificationEnabled")
	if err != nil {
		return errors.New("Could not create notification")
	}
	for _, u := range users.Users {
			grpc_common.CreateNotification(ctx, u.UserId, story.UserId, "Story", story.Id)
	}
	return nil
}

func (service *StoryService) GetStoryById(ctx context.Context, id string) (domain.Story, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoryById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbStory, err := service.storyRepository.GetStoryById(ctx, id)
	if err != nil { return domain.Story{}, err }

	story, err := service.retrieveStoryAdditionalData(ctx, *dbStory)
	if err != nil { return domain.Story{}, err }

	return story, nil
}

func (service *StoryService) RemoveStory(ctx context.Context, id string, userId string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if id == "" {
		return errors.New("cannot remove story")
	}

	return service.storyRepository.RemoveStory(ctx, id, userId)
}

func (service *StoryService) retrieveStoryAdditionalData(ctx context.Context, story persistence.Story) (domain.Story, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "retrieveStoryAdditionalData")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbMedia, err := service.mediaRepository.GetMediaForPost(ctx, story.Id)
	if err != nil { return domain.Story{}, err }
	media := []domain.Media{}
	for _, single := range dbMedia{
		tags, err := service.tagRepository.GetTagsForMedia(ctx, single.Id)
		if err != nil { return domain.Story{}, err }

		for index, tag := range tags {
			username, err := grpc_common.GetUsernameById(ctx, tag.UserId)
			if username == "" || err != nil {
				return domain.Story{}, errors.New("cannot retrieve tags")
			}
			tags[index].Username = username
		}

		converted, err := single.ConvertToDomain(tags)
		if err != nil { return domain.Story{}, err }

		media = append(media, converted)
	}

	hashtags, err := service.hashtagRepository.GetPostHashtags(ctx, story.Id)
	if err != nil { return domain.Story{}, err }

	converted := story.ConvertToDomain(media, hashtags)

	return converted, nil
}

func (service *StoryService) GetStoriesFromHighlight(ctx context.Context, highlightId string) ([]domain.Story, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoriesFromHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbStories, err := service.storyRepository.GetHighlightsStories(ctx, highlightId)
	if err != nil { return []domain.Story{}, err }

	stories := []domain.Story{}
	for _, dbStory := range dbStories {
		story, err := service.retrieveStoryAdditionalData(ctx, dbStory)
		if err != nil { return []domain.Story{}, err }
		stories = append(stories, story)
	}

	return stories, nil
}
