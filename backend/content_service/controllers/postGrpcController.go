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

type PostGrpcController struct {
	service    		*services.PostService
	campaignService *services.CampaignService
	jwtManager *common.JWTManager
	logger     *logger.Logger
}

func NewPostController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger) (*PostGrpcController, error) {
	service, err := services.NewPostService(db)
	if err != nil {
		return nil, err
	}

	campaignService, err := services.NewCampaignService(db)
	if err != nil { return nil, err }

	return &PostGrpcController{
		service,
		campaignService,
		jwtManager,
		logger,
	}, nil
}

func (c *PostGrpcController) CreatePost(ctx context.Context, in *protopb.Post) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	c.logger.ToStdoutAndFile("CreatePost", "Post creation attempt by "+in.UserId, logger.Info)

	if err != nil {
		c.logger.ToStdoutAndFile("CreatePost", "Post creation attempt failed by "+in.UserId+", invalid JWT", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	} else if claims.UserId == "" {
		c.logger.ToStdoutAndFile("CreatePost", "Post creation attempt failed by, invalid JWT format", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	} else if claims.UserId != in.UserId {
		c.logger.ToStdoutAndFile("CreatePost", "Post creation attempt failed by "+claims.UserId+" tried creating post for "+in.UserId, logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot create post for another user")
	}

	var post *domain.Post
	post = post.ConvertFromGrpc(in)

	for _, media := range post.Media {
		for _, tag := range media.Tags {
			following, err := grpc_common.CheckFollowInteraction(ctx, tag.UserId, post.UserId)
			if err != nil {
				c.logger.ToStdoutAndFile("CreatePost", "Post creation attempt failed by "+claims.UserId+", cannot tag "+tag.UserId, logger.Error)
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
			}

			isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.UserId)
			if err != nil {
				c.logger.ToStdoutAndFile("CreatePost", "Post creation attempt failed by "+claims.UserId+", cannot tag "+tag.UserId, logger.Error)
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
			}

			isBlocked, err := grpc_common.CheckIfBlocked(ctx, in.UserId, claims.UserId)
			if err != nil {
				c.logger.ToStdoutAndFile("CreatePost", "Post creation attempt failed by "+claims.UserId+", cannot tag "+tag.UserId, logger.Error)
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
			}

			// If used is blocked or his profile is private and did not approve your request
			if isBlocked || (!isPublic && !following.IsApprovedRequest) {
				c.logger.ToStdoutAndFile("CreatePost", "Post creation attempt failed by "+claims.UserId+", cannot tag "+tag.UserId, logger.Error)
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
			}

			username, err := grpc_common.GetUsernameById(ctx, tag.UserId)
			if username == "" || err != nil {
				c.logger.ToStdoutAndFile("CreatePost", "Post creation attempt failed by "+claims.UserId+", cannot tag "+tag.UserId, logger.Error)
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
			}
		}
	}

	err = c.service.CreatePost(ctx, post)
	if err != nil {
		c.logger.ToStdoutAndFile("CreatePost", "Post creation attempt failed by "+claims.UserId+", due to server error", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}

	c.logger.ToStdoutAndFile("CreatePost", "Post creation successful by "+claims.UserId, logger.Info)
	return &protopb.EmptyResponseContent{}, nil
}

func (c *PostGrpcController) GetAllPostsReduced(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPostsReduced")
	defer span.Finish()
	claims, _ := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	userIds := []string{}
	if claims.UserId == "" {
		publicIds, err := grpc_common.GetPublicUsers(ctx)
		if err != nil {
			return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
		}
		for _, id := range publicIds {
			userIds = append(userIds, id)
		}
	} else {
		homepageIds, err := grpc_common.GetHomepageUsers(ctx, claims.UserId)
		if err != nil {
			return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
		}
		for _, id := range homepageIds {
			userIds = append(userIds, id)
		}
	}

	posts, err := c.service.GetAllPostsReduced(ctx, userIds)
	if err != nil {
		return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}

func (c *PostGrpcController) GetAllPosts(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.PostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	claims, _ := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	userIds := []string{}
	if claims.UserId == "" {
		publicIds, err := grpc_common.GetPublicUsers(ctx)
		if err != nil {
			return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error())
		}
		for _, id := range publicIds {
			userIds = append(userIds, id)
		}
	} else {
		homepageIds, err := grpc_common.GetHomepageUsers(ctx, claims.UserId)
		if err != nil {
			return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error())
		}
		for _, id := range homepageIds {
			userIds = append(userIds, id)
		}
	}

	posts, err := c.service.GetAllPosts(ctx, userIds)
	if err != nil {
		return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.Post{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	ads, err := c.campaignService.GetOngoingCampaignsAds(ctx, userIds, claims.UserId, model.TypePost)
	if err != nil { return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error()) }

	responseAds := []*protopb.Ad{}
	for _, ad := range ads {
		responseAds = append(responseAds, ad.ConvertToGrpc())
	}

	return &protopb.PostArray{
		Posts: responsePosts,
		Ads: responseAds,
	}, nil
}

func (c *PostGrpcController) GetPostsForUser(ctx context.Context, in *protopb.RequestId) (*protopb.PostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsForUser")
	defer span.Finish()
	claims, _ := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if claims.UserId == "" {
		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
		if err != nil {
			return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error())
		}
		if !isPublic {
			return &protopb.PostArray{}, status.Errorf(codes.Unknown, "this user is not public")
		}
	} else if in.Id != claims.UserId {
		followConnection, err := grpc_common.CheckFollowInteraction(ctx, in.Id, claims.UserId)
		if err != nil {
			return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
		if err != nil {
			return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, in.Id, claims.UserId)
		if err != nil {
			return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		// If used is blocked or his profile is private and did not approve your request
		if isBlocked || (!isPublic && !followConnection.IsApprovedRequest) {
			return &protopb.PostArray{}, nil
		}
	}

	posts, err := c.service.GetPostsForUser(ctx, in.Id)
	if err != nil {
		return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := domain.ConvertMultiplePostsToGrpc(posts)

	return &protopb.PostArray{Posts: responsePosts}, nil
}

func (c *PostGrpcController) GetPostById(ctx context.Context, id string) (*protopb.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostById")
	defer span.Finish()
	claims, _ := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if id == "" {
		return &protopb.Post{}, status.Errorf(codes.InvalidArgument, "cannot retrieve non-existing posts")
	}

	post, err := c.service.GetPostById(ctx, id)
	if err != nil {
		return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
	}

	if claims.UserId == "" {
		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, post.UserId)
		if err != nil {
			return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
		}
		if !isPublic {
			return &protopb.Post{}, status.Errorf(codes.Unknown, "this post is not public")
		}
	} else if post.UserId != claims.UserId {
		following, err := grpc_common.CheckFollowInteraction(ctx, post.UserId, claims.UserId)
		if err != nil {
			return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
		}

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, post.UserId)
		if err != nil {
			return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
		}

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, post.UserId, claims.UserId)
		if err != nil {
			return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
		}

		if (!following.IsApprovedRequest && !isPublic) || isBlocked {
			return &protopb.Post{}, status.Errorf(codes.PermissionDenied, "cannot retrieve this post")
		}
	}

	grpcPost := post.ConvertToGrpc()
	return grpcPost, nil
}

func (c *PostGrpcController) RemovePost(ctx context.Context, id string) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemovePost")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	c.logger.ToStdoutAndFile("RemovePost", "Post removal attempt by "+claims.UserId, logger.Info)

	if err != nil {
		c.logger.ToStdoutAndFile("RemovePost", "Post removal attempt failed by "+claims.UserId+", invalid JWT", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	} else if claims.UserId == "" {
		c.logger.ToStdoutAndFile("RemovePost", "Post removal attempt failed by "+claims.UserId+", invalid JWT", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove other people's posts")
	} else if id == "" {
		c.logger.ToStdoutAndFile("RemovePost", "Post removal attempt failed by "+claims.UserId+", no post id provided", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove non-existing posts")
	}

	err = c.service.RemovePost(ctx, id, claims.UserId)
	if err != nil {
		c.logger.ToStdoutAndFile("RemovePost", "Post removal attempt failed by "+claims.UserId+", server error", logger.Error)
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}

	c.logger.ToStdoutAndFile("RemovePost", "Post removal attempt successful by "+claims.UserId, logger.Info)
	return &protopb.EmptyResponseContent{}, nil
}

func (c *PostGrpcController) SearchContentByLocation(ctx context.Context, in *protopb.SearchLocationRequest) (*protopb.PostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SearchContentByLocation")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	location := in.Location

	posts, err := c.service.SearchContentByLocation(ctx, location)
	if err != nil {
		return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.Post{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.PostArray{
		Posts: responsePosts,
	}, nil
}

func (c *PostGrpcController) GetPostsByHashtag(ctx context.Context, in *protopb.Hashtag) (*protopb.PostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsByHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts, err := c.service.GetPostsByHashtag(ctx, in.Text)
	if err != nil {
		return &protopb.PostArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.Post{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.PostArray{
		Posts: responsePosts,
	}, nil
}
