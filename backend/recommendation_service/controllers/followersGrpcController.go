package controllers

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/model"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/services"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowersGrpcController struct {
	service *services.FollowersService
	logger  *logger.Logger
}

func NewFollowersController(driver neo4j.Driver, logger *logger.Logger) (*FollowersGrpcController, error) {
	service, err := services.NewFollowersService(driver)
	if err != nil {
		return nil, err
	}

	return &FollowersGrpcController{
		service,
		logger,
	}, nil
}

func (controller *FollowersGrpcController) CreateUserConnection(ctx context.Context, in *protopb.CreateFollowerRequest) (*protopb.EmptyResponseFollowers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	err := controller.service.CreateUserConnection(ctx, follower)

	if err != nil {
		return &protopb.EmptyResponseFollowers{}, errors.New("Could not make follow!")
	}
	return &protopb.EmptyResponseFollowers{}, nil
}

func (controller *FollowersGrpcController) GetAllFollowing(ctx context.Context, in *protopb.CreateUserRequestFollowers) (*protopb.CreateUserResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowing")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user = model.User{}
	user = *user.ConvertFromGrpc(in.User)
	users, err := controller.service.GetAllFollowing(ctx, user.UserId)
	if err != nil {
		return nil, errors.New("Could not get all followings")
	}

	return &protopb.CreateUserResponse{Users: user.ConvertAllToGrpc(users)}, err
}

func (controller *FollowersGrpcController) GetAllFollowingsForHomepage(ctx context.Context, in *protopb.CreateUserRequestFollowers) (*protopb.CreateUserResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowingsForHomepage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user = model.User{}
	user = *user.ConvertFromGrpc(in.User)
	users, err := controller.service.GetAllFollowingsForHomepage(ctx, user.UserId)
	if err != nil {
		return nil, errors.New("Could not get all followings")
	}

	return &protopb.CreateUserResponse{Users: user.ConvertAllToGrpc(users)}, err
}

func (controller *FollowersGrpcController) CheckIfMuted(ctx context.Context, in *protopb.Follower) (*protopb.BooleanResponseFollowers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIfMuted")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	isMuted, err := controller.service.CheckIfMuted(ctx, in.FollowerId, in.UserId)
	if err != nil {
		return nil, errors.New("could not get close friends")
	}

	return &protopb.BooleanResponseFollowers{ Response: isMuted }, nil
}

func (controller *FollowersGrpcController) GetCloseFriends(ctx context.Context, in *protopb.RequestIdFollowers) (*protopb.CreateUserResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCloseFriends")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user model.User
	users, err := controller.service.GetCloseFriends(ctx, in.Id)
	if err != nil {
		return nil, errors.New("could not get close friends")
	}

	return &protopb.CreateUserResponse{Users: user.ConvertAllToGrpc(users)}, err
}

func (controller *FollowersGrpcController) GetCloseFriendsReversed(ctx context.Context, in *protopb.RequestIdFollowers) (*protopb.CreateUserResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCloseFriendsReversed")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user model.User
	users, err := controller.service.GetCloseFriendsReversed(ctx, in.Id)
	if err != nil {
		return nil, errors.New("could not get close friends")
	}

	return &protopb.CreateUserResponse{Users: user.ConvertAllToGrpc(users)}, err
}

func (controller *FollowersGrpcController) GetAllFollowers(ctx context.Context, in *protopb.CreateUserRequestFollowers) (*protopb.CreateUserResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user model.User
	user = *user.ConvertFromGrpc(in.User)
	users, err := controller.service.GetAllFollowers(ctx, user.UserId)
	if err != nil {
		return nil, errors.New("Could not get all followers!")
	}

	return &protopb.CreateUserResponse{Users: user.ConvertAllToGrpc(users)}, err
}

func (controller *FollowersGrpcController) GetFollowersConnection(ctx context.Context, in *protopb.Follower) (*protopb.Follower, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var follower *model.Follower
	follower = follower.ConvertFromGrpc(in)
	follower, err := controller.service.GetFollowersConnection(ctx, *follower)
	if err != nil {
		return nil, errors.New("couldn't get connection between users")
	}

	return follower.ConvertToGrpc(), err
}

func (controller *FollowersGrpcController) DeleteBiDirectedConnection(ctx context.Context, in *protopb.CreateFollowerRequest) (*protopb.EmptyResponseFollowers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteBiDirectedConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	_, err := controller.service.DeleteBiDirectedConnection(ctx, follower)

	if err != nil {
		return &protopb.EmptyResponseFollowers{}, errors.New("Could not delete bidirected connection!")
	}
	return &protopb.EmptyResponseFollowers{}, nil
}

func (controller *FollowersGrpcController) DeleteDirectedConnection(ctx context.Context, in *protopb.CreateFollowerRequest) (*protopb.EmptyResponseFollowers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteDirectedConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	_, err := controller.service.DeleteDirectedConnection(ctx, follower)

	if err != nil {
		return &protopb.EmptyResponseFollowers{}, errors.New("Could not delete directed connection!")
	}
	return &protopb.EmptyResponseFollowers{}, nil
}

func (controller *FollowersGrpcController) CreateUser(ctx context.Context, in *protopb.CreateUserRequestFollowers) (*protopb.EmptyResponseFollowers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	controller.logger.ToStdoutAndFile("CreateUser", "Create user node attempt for " + in.User.UserId, logger.Info)

	var user = model.User{}
	user = *user.ConvertFromGrpc(in.User)
	_, err := controller.service.CreateUser(ctx, user)

	if err != nil {
		controller.logger.ToStdoutAndFile("CreateUser", "Create user node attempt failed for " + in.User.UserId, logger.Error)
		return &protopb.EmptyResponseFollowers{}, errors.New("could not create User")
	}

	controller.logger.ToStdoutAndFile("CreateUser", "Create user node attempt successful for " + in.User.UserId, logger.Info)
	return &protopb.EmptyResponseFollowers{}, nil
}

func (controller *FollowersGrpcController) UpdateUserConnection(ctx context.Context, in *protopb.CreateFollowerRequest) (*protopb.CreateFollowerResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	result, err := controller.service.UpdateUserConnection(ctx, follower)

	if err != nil {
		return nil, errors.New("Could not update follower info!")
	}

	return &protopb.CreateFollowerResponse{Follower: result.ConvertToGrpc()}, nil
}

func (controller *FollowersGrpcController)  AcceptFollowRequest(ctx context.Context,in *protopb.CreateFollowerRequest) (*protopb.CreateFollowerResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "AcceptFollowRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	result, err := controller.service.AcceptFollowRequest(ctx, follower)

	if err != nil {
		return nil, errors.New("Could not accept follow request!")
	}

	return &protopb.CreateFollowerResponse{Follower: result.ConvertToGrpc()}, nil

}

func (c *FollowersGrpcController) GetUsersForNotificationEnabled(ctx context.Context, in *protopb.RequestForNotification) (*protopb.CreateUserResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsersForNotificationEnabled")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	var user model.User
	users, err := c.service.GetUsersForNotificationEnabled(ctx, in.UserId, in.NotificationType)
	if err != nil {
		return nil, err
	}

	return &protopb.CreateUserResponse{Users: user.ConvertAllToGrpc(users)}, nil

}
