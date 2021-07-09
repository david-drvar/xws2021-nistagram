package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	otgo "github.com/opentracing/opentracing-go"
	"io"
)

type Server struct {
	closer io.Closer
	tracer otgo.Tracer
	protopb.UnimplementedFollowersServer
	followerController *FollowersGrpcController
	recommendationController *RecommendationGrpcController
}

func NewServer(driver neo4j.Driver, logger *logger.Logger) (*Server, error) {
	followerController, _ := NewFollowersController(driver, logger)
	tracer, closer := tracer.Init("recommendationService")
	recommendationController, _ := NewRecommendationGrpcController(driver, logger)
	otgo.SetGlobalTracer(tracer)
	return &Server{
		tracer:             tracer,
		closer:             closer,
		followerController: followerController,
		recommendationController: &recommendationController,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

func (s *Server) CreateUserConnection(ctx context.Context, in *protopb.CreateFollowerRequest) (*protopb.EmptyResponseFollowers, error) {
	return s.followerController.CreateUserConnection(ctx, in)
}

func (s *Server) GetAllFollowers(ctx context.Context, in *protopb.CreateUserRequestFollowers) (*protopb.CreateUserResponse, error) {
	return s.followerController.GetAllFollowers(ctx, in)
}

func (s *Server) GetAllFollowing(ctx context.Context, in *protopb.CreateUserRequestFollowers) (*protopb.CreateUserResponse, error) {
	return s.followerController.GetAllFollowing(ctx, in)
}

func (s *Server) GetAllFollowingsForHomepage(ctx context.Context, in *protopb.CreateUserRequestFollowers) (*protopb.CreateUserResponse, error) {
	return s.followerController.GetAllFollowingsForHomepage(ctx, in)
}

func (s *Server) DeleteBiDirectedConnection(ctx context.Context, in *protopb.CreateFollowerRequest) (*protopb.EmptyResponseFollowers, error) {
	return s.followerController.DeleteBiDirectedConnection(ctx, in)
}

func (s *Server) DeleteDirectedConnection(ctx context.Context, in *protopb.CreateFollowerRequest) (*protopb.EmptyResponseFollowers, error) {
	return s.followerController.DeleteDirectedConnection(ctx, in)
}

func (s *Server) CreateUser(ctx context.Context, in *protopb.CreateUserRequestFollowers) (*protopb.EmptyResponseFollowers, error) {
	return s.followerController.CreateUser(ctx, in)
}

func (s *Server) UpdateUserConnection(ctx context.Context, in *protopb.CreateFollowerRequest) (*protopb.CreateFollowerResponse, error) {
	return s.followerController.UpdateUserConnection(ctx, in)
}

func (s *Server) GetCloseFriends(ctx context.Context, in *protopb.RequestIdFollowers) (*protopb.CreateUserResponse, error) {
	return s.followerController.GetCloseFriends(ctx, in)
}

func (s *Server) GetFollowersConnection(ctx context.Context, in *protopb.Follower) (*protopb.Follower, error) {
	return s.followerController.GetFollowersConnection(ctx, in)
}

func (s *Server)  AcceptFollowRequest(ctx context.Context,in *protopb.CreateFollowerRequest) (*protopb.CreateFollowerResponse, error) {
	return s.followerController.AcceptFollowRequest(ctx, in)
}

func (s *Server) GetUsersForNotificationEnabled(ctx context.Context, in *protopb.RequestForNotification) (*protopb.CreateUserResponse, error) {
	return s.followerController.GetUsersForNotificationEnabled(ctx, in)
}

func (s *Server) RecommendationPattern(ctx context.Context, in *protopb.RequestIdFollowers) (*protopb.RecommendationResponse, error) {
	return s.recommendationController.RecommendationPattern(ctx, in)
}