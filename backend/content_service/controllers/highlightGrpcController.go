package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type HighlightGrpcController struct {
	service 	 *services.HighlightService
	storyService *services.StoryService
	jwtManager  *common.JWTManager
}

func NewHighlightController(db *gorm.DB, jwtManager *common.JWTManager) (*HighlightGrpcController, error) {
	service, err := services.NewHighlightService(db)
	if err != nil {
		return nil, err
	}

	storyService, err := services.NewStoryService(db)
	if err != nil {
		return nil, err
	}

	return &HighlightGrpcController{
		service,
		storyService,
		jwtManager,
	}, nil
}

func (c *HighlightGrpcController) GetAllHighlights (ctx context.Context, in *protopb.RequestId) (*protopb.HighlightsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHighlights")
	defer span.Finish()
	claims, _ := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)


	isCloseFriend := false
	if claims.UserId == ""{
		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
		if err != nil { return &protopb.HighlightsArray{}, status.Errorf(codes.Unknown, err.Error()) }
		if !isPublic{
			return &protopb.HighlightsArray{}, status.Errorf(codes.Unknown, "this user is private")
		}
	}else if claims.UserId != in.Id {
		followConnection, err := grpc_common.CheckFollowInteraction(ctx, in.Id, claims.UserId)
		if err != nil {
			return &protopb.HighlightsArray{}, status.Errorf(codes.Unknown, err.Error())
		}
		isCloseFriend = followConnection.IsCloseFriends

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
		if err != nil {
			return &protopb.HighlightsArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, in.Id, claims.UserId)
		if err != nil {
			return &protopb.HighlightsArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		// If used is blocked or his profile is private and did not approve your request
		if isBlocked || (!isPublic && !followConnection.IsApprovedRequest ) {
			return &protopb.HighlightsArray{}, nil
		}
	}else if claims.UserId == in.Id {
		isCloseFriend = true
	}

	highlights, err := c.service.GetAllHighlights(ctx, in.Id)
	if err != nil {
		return &protopb.HighlightsArray{}, status.Errorf(codes.Unknown, "could not retrieve highlights")
	}

	// Allow retrieving only non-close friend stories
	if !isCloseFriend {
		for index, highlight := range highlights{
			newStories := []domain.Story{}
			for _, story := range highlight.Stories {
				if !story.IsCloseFriends {
					newStories = append(newStories, story)
				}
			}
			highlights[index].Stories = newStories
		}
	}

	grpcHighlights := []*protopb.Highlight{}
	for _, highlight := range highlights{
		grpcHighlights = append(grpcHighlights, highlight.ConvertToGrpc())
	}

	return &protopb.HighlightsArray{
		Highlights: grpcHighlights,
	}, nil
}

func (c *HighlightGrpcController) GetHighlight (ctx context.Context, in *protopb.RequestId) (*protopb.Highlight, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetHighlight")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if in.Id == "" {
		return &protopb.Highlight{}, status.Errorf(codes.InvalidArgument, "no highlight id provided")
	}

	highlight, err := c.service.GetHighlight(ctx, in.Id)
	if err != nil || highlight.Id == "" {
		return &protopb.Highlight{}, status.Errorf(codes.Unknown, "could not retrieve highlight")
	}

	isCloseFriend := false
	if claims.UserId == ""{
		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, highlight.UserId)
		if err != nil {
			return &protopb.Highlight{}, status.Errorf(codes.Unknown, err.Error())
		}
		if !isPublic{
			return &protopb.Highlight{}, status.Errorf(codes.Unknown, "this highlight is not public")
		}
	}else if claims.UserId != highlight.UserId {
		followConnection, err := grpc_common.CheckFollowInteraction(ctx, highlight.UserId, claims.UserId)
		if err != nil {
			return &protopb.Highlight{}, status.Errorf(codes.Unknown, err.Error())
		}
		isCloseFriend = followConnection.IsCloseFriends

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, highlight.UserId)
		if err != nil {
			return &protopb.Highlight{}, status.Errorf(codes.Unknown, err.Error())
		}

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, highlight.UserId, claims.UserId)
		if err != nil {
			return &protopb.Highlight{}, status.Errorf(codes.Unknown, err.Error())
		}

		// If used is blocked or his profile is private and did not approve your request
		if isBlocked || (!isPublic && !followConnection.IsApprovedRequest ) {
			return &protopb.Highlight{}, nil
		}
	}

	// Allow retrieving only non-close friend stories
	if !isCloseFriend {
		newStories := []domain.Story{}
		for _, story := range highlight.Stories {
			if !story.IsCloseFriends {
				newStories = append(newStories, story)
			}
		}
		highlight.Stories = newStories
	}

	grpcHighlight := highlight.ConvertToGrpc()

	return grpcHighlight, nil
}

func (c *HighlightGrpcController) CreateHighlightStory(ctx context.Context, in *protopb.HighlightRequest) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHighlightStory")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}  else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.UserId {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot save story for another user")
	}

	story, err := c.storyService.GetStoryById(ctx, in.StoryId)
	if err != nil { return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot save story") }

	if claims.UserId != story.UserId {
		followConnection, err := grpc_common.CheckFollowInteraction(ctx, story.UserId, claims.UserId)
		if err != nil {
			return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
		}

		if !followConnection.IsCloseFriends && story.IsCloseFriends {
			return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot save non-close friends story")
		}

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, story.UserId)
		if err != nil {
			return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
		}

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, story.UserId, claims.UserId)
		if err != nil {
			return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
		}

		// If used is blocked or his profile is private and did not approve your request
		if isBlocked || (!isPublic && !followConnection.IsApprovedRequest ) {
			return &protopb.EmptyResponseContent{}, nil
		}
	}

	var highlightRequest *domain.HighlightRequest
	highlightRequest = highlightRequest.ConvertFromGrpc(in)

	err = c.service.CreateHighlightStory(ctx, *highlightRequest)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create story from highlight")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *HighlightGrpcController) RemoveHighlightStory(ctx context.Context, in *protopb.HighlightRequest) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveHighlightStory")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}  else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.UserId {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot create post for another user")
	}

	var highlightRequest *domain.HighlightRequest
	highlightRequest = highlightRequest.ConvertFromGrpc(in)

	err = c.service.RemoveHighlightStory(ctx, *highlightRequest)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not remove story from highlight")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *HighlightGrpcController) CreateHighlight (ctx context.Context, in *protopb.Highlight) (*protopb.Highlight, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHighlight")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.Highlight{}, status.Errorf(codes.Unknown, err.Error())
	}  else if claims.UserId == "" {
		return &protopb.Highlight{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.UserId {
		return &protopb.Highlight{}, status.Errorf(codes.Unknown, "cannot create post for another user")
	}

	var collection *domain.Highlight
	collection = collection.ConvertFromGrpc(in)

	highlight, err := c.service.CreateHighlight(ctx, *collection)
	if err != nil {
		return &protopb.Highlight{}, status.Errorf(codes.Unknown, "could not create highlight")
	}

	highlightResponse := highlight.ConvertToGrpc()
	return highlightResponse, nil
}

func (c *HighlightGrpcController) RemoveHighlight (ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveHighlight")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}  else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}

	err = c.service.RemoveHighlight(ctx, in.Id, claims.UserId)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not remove highlight")
	}

	return &protopb.EmptyResponseContent{}, nil
}