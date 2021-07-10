package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/model"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/repositories"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowersService struct {
	repository repositories.FollowersRepository
}

func NewFollowersService(driver neo4j.Driver) (*FollowersService, error) {
	repository, err := repositories.NewFollowersRepository(driver)

	return &FollowersService{
		repository: repository,
	}, err
}

func (service *FollowersService) CreateUserConnection(ctx context.Context, follower model.Follower)  error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	follower.IsCommentNotificationEnabled = true
	follower.IsMessageNotificationEnabled = true
	follower.IsPostNotificationEnabled = true
	follower.IsStoryNotificationEnabled = true
	follower.IsCloseFriends = false
	follower.IsMuted = false

	privacy, err := grpc_common.CheckUserProfilePublic(ctx, follower.FollowerId)
	if err != nil {
		return err
	}
	if !privacy {
		follower.IsApprovedRequest = false
		follower.RequestIsPending=true
		res, err := service.repository.CreateUserConnection(ctx, follower)
		if err != nil || res == false {
			return err
		}
		return grpc_common.CreateNotification(ctx, follower.FollowerId, follower.UserId,  "FollowPrivate", "")
	}else {
		follower.IsApprovedRequest = true
		follower.RequestIsPending=false
		res, err := service.repository.CreateUserConnection(ctx, follower)
		if err != nil || res == false {
			return err
		}
		return grpc_common.CreateNotification(ctx, follower.FollowerId ,follower.UserId, "FollowPublic", "")
	}

}

func (service *FollowersService) GetAllFollowers(ctx context.Context, userId string) ([]model.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetAllFollowers(ctx, userId)
}

func (service *FollowersService) GetAllFollowing(ctx context.Context, userId string) ([]model.User, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowing")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetAllFollowing(ctx, userId)
}

func (service *FollowersService) GetAllFollowingsForHomepage(ctx context.Context, userId string) ([]model.User, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowingsForHomepagePosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetAllFollowingsForHomepage(ctx, userId)
}

func (service *FollowersService) CheckIfMuted(ctx context.Context, userId string, requestingUserId string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIfMuted")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	nonmuteds, err := service.repository.CheckIfMuted(ctx, requestingUserId)
	if err != nil { return true, err }

	isMuted := true
	for _, nonmuted := range nonmuteds{
		if nonmuted.UserId == userId{
			isMuted = false
			break
		}
	}

	return isMuted, nil
}

func (service *FollowersService) GetCloseFriends(ctx context.Context, userId string) ([]model.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCloseFriends")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetCloseFriends(ctx, userId)
}

func (service *FollowersService) GetCloseFriendsReversed(ctx context.Context, userId string) ([]model.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCloseFriendsReversed")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetCloseFriendsReversed(ctx, userId)
}

func (service *FollowersService) CreateUser(ctx context.Context, u model.User) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.CreateUser(ctx, u);
}

func (service *FollowersService) DeleteDirectedConnection(ctx context.Context, f model.Follower) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := service.repository.DeleteDirectedConnection(ctx, f)
	if err != nil {
		return false, err
	}
	err = grpc_common.DeleteByTypeAndCreator(ctx, "FollowPrivate", f.FollowerId, f.UserId)
	if err != nil {
		return false, err
	}
	err = grpc_common.DeleteByTypeAndCreator(ctx, "FollowPublic", f.FollowerId, f.UserId)
	if err != nil {
		return false, err
	}

	return true, nil

}

func (service *FollowersService) DeleteBiDirectedConnection(ctx context.Context, f model.Follower) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	_, err := service.repository.DeleteBiDirectedConnection(ctx, f)
	if err != nil {
		return false, err
	}
	err = grpc_common.DeleteByTypeAndCreator(ctx, "FollowPublic",  f.UserId, f.FollowerId)
	if err != nil {
		return false, err
	}
	err = grpc_common.DeleteByTypeAndCreator(ctx, "FollowPublic", f.FollowerId, f.UserId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (service *FollowersService)  UpdateUserConnection(ctx context.Context, f model.Follower) (*model.Follower,error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.UpdateUserConnection(ctx, f)
}

func (service *FollowersService) GetFollowersConnection(ctx context.Context, f model.Follower) (*model.Follower, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetFollowersConnection(ctx, f)
}

func (service *FollowersService) AcceptFollowRequest(ctx context.Context, f model.Follower) (*model.Follower, error ){
	span := tracer.StartSpanFromContextMetadata(ctx, "AcceptFollowRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	follower, err := service.UpdateUserConnection(ctx, f)
	if err != nil {
		return nil, err
	}

	notification, err := grpc_common.GetByTypeAndCreator(ctx, f.FollowerId, f.UserId, "FollowPrivate")
	if err != nil {
		return nil, err
	}

	err = grpc_common.UpdateNotification(ctx, notification.Id, "FollowPublic", " started following you.")
	if err != nil {
		return nil, err
	}
	return follower, nil

}

func (service *FollowersService) GetUsersForNotificationEnabled(ctx context.Context, userId string,notification string) ([]model.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsersForNotificationEnabled")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetUsersForNotificationEnabled(ctx, userId, notification)

}



