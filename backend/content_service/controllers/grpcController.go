package controllers

import (
	"context"
	"io"

	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	otgo "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type Server struct {
	protopb.UnimplementedContentServer
	postController      *PostGrpcController
	commentController   *CommentGrpcController
	likeController      *LikeGrpcController
	favoritesController *FavoritesGrpcController
	hashtagController   *HashtagGrpcController
	storyController     *StoryGrpcController
	highlightController *HighlightGrpcController
	complaintController *ComplaintGrpcController
	tracer              otgo.Tracer
	closer              io.Closer
}

func NewServer(db *gorm.DB, manager *common.JWTManager, logger *logger.Logger) (*Server, error) {
	postController, _ := NewPostController(db, manager, logger)
	storyController, _ := NewStoryController(db, manager, logger)
	commentController, _ := NewCommentController(db, manager)
	likeController, _ := NewLikeController(db, manager)
	favoritesController, _ := NewFavoritesController(db, manager)
	hashtagController, _ := NewHashtagController(db)
	highlightController, _ := NewHighlightController(db, manager)
	complaintController, _ := NewComplaintController(db, manager)
	tracer, closer := tracer.Init("global_ContentGrpcController")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		postController:      postController,
		commentController:   commentController,
		likeController:      likeController,
		favoritesController: favoritesController,
		hashtagController:   hashtagController,
		storyController:     storyController,
		highlightController: highlightController,
		complaintController: complaintController,
		tracer:              tracer,
		closer:              closer,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

/*   Posts   */
func (s *Server) CreatePost(ctx context.Context, in *protopb.Post) (*protopb.EmptyResponseContent, error) {
	return s.postController.CreatePost(ctx, in)
}

func (s *Server) GetAllPostsReduced(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.ReducedPostArray, error) {
	return s.postController.GetAllPostsReduced(ctx, in)
}

func (s *Server) GetAllPosts(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.PostArray, error) {
	return s.postController.GetAllPosts(ctx, in)
}

func (s *Server) GetPostsForUser(ctx context.Context, in *protopb.RequestId) (*protopb.PostArray, error) {
	return s.postController.GetPostsForUser(ctx, in)
}

func (s *Server) GetPostById(ctx context.Context, in *protopb.RequestId) (*protopb.Post, error) {
	return s.postController.GetPostById(ctx, in.Id)
}

func (s *Server) RemovePost(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	return s.postController.RemovePost(ctx, in.Id)
}

/*   Stories   */
func (s *Server) CreateStory(ctx context.Context, in *protopb.Story) (*protopb.EmptyResponseContent, error) {
	return s.storyController.CreateStory(ctx, in)
}

func (s *Server) GetAllStories(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.StoriesHome, error) {
	return s.storyController.GetAllStories(ctx, in)
}

func (s *Server) GetStoriesForUser(ctx context.Context, in *protopb.RequestId) (*protopb.StoriesArray, error) {
	return s.storyController.GetStoriesForUser(ctx, in)
}

func (s *Server) GetMyStories(ctx context.Context, in *protopb.RequestId) (*protopb.StoriesArray, error) {
	return s.storyController.GetMyStories(ctx, in)
}

func (s *Server) GetStoryById(ctx context.Context, in *protopb.RequestId) (*protopb.Story, error) {
	return s.storyController.GetStoryById(ctx, in)
}

func (s *Server) RemoveStory(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	return s.storyController.RemoveStory(ctx, in)
}

/*   Comments   */
func (s *Server) CreateComment(ctx context.Context, in *protopb.Comment) (*protopb.EmptyResponseContent, error) {
	return s.commentController.CreateComment(ctx, in)
}

func (s *Server) GetCommentsForPost(ctx context.Context, in *protopb.RequestId) (*protopb.CommentsArray, error) {
	return s.commentController.GetCommentsForPost(ctx, in.Id)
}

/* Likes & Dislikes */
func (s *Server) CreateLike(ctx context.Context, in *protopb.Like) (*protopb.EmptyResponseContent, error) {
	return s.likeController.CreateLike(ctx, in)
}

func (s *Server) GetLikesForPost(ctx context.Context, in *protopb.RequestId) (*protopb.LikesArray, error) {
	return s.likeController.GetLikesForPost(ctx, in.Id, true)
}

func (s *Server) GetUserLikedOrDislikedPosts(ctx context.Context, in *protopb.Like) (*protopb.PostArray, error) {
	return s.likeController.GetUserLikedOrDislikedPosts(ctx, in)
}

/* Collections & Favorites */
func (s *Server) GetAllCollections(ctx context.Context, in *protopb.RequestId) (*protopb.CollectionsArray, error) {
	return s.favoritesController.GetAllCollections(ctx, in)
}

func (s *Server) GetUserFavoritesOptimized(ctx context.Context, in *protopb.RequestId) (*protopb.Favorites, error) {
	return s.favoritesController.GetUserFavoritesOptimized(ctx, in)
}

func (s *Server) GetCollection(ctx context.Context, in *protopb.RequestId) (*protopb.Collection, error) {
	return s.favoritesController.GetCollection(ctx, in)
}

func (s *Server) CreateCollection(ctx context.Context, in *protopb.Collection) (*protopb.Collection, error) {
	return s.favoritesController.CreateCollection(ctx, in)
}

func (s *Server) RemoveCollection(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	return s.favoritesController.RemoveCollection(ctx, in)
}

func (s *Server) GetUserFavorites(ctx context.Context, in *protopb.RequestId) (*protopb.Favorites, error) {
	return s.favoritesController.GetUserFavorites(ctx, in)
}

func (s *Server) CreateFavorite(ctx context.Context, in *protopb.FavoritesRequest) (*protopb.EmptyResponseContent, error) {
	return s.favoritesController.CreateFavorite(ctx, in)
}

func (s *Server) RemoveFavorite(ctx context.Context, in *protopb.FavoritesRequest) (*protopb.EmptyResponseContent, error) {
	return s.favoritesController.RemoveFavorite(ctx, in)
}

func (s *Server) SearchContentByLocation(ctx context.Context, in *protopb.SearchLocationRequest) (*protopb.PostArray, error) {
	return s.postController.SearchContentByLocation(ctx, in)
}

/*   Hashtags   */
func (s *Server) CreateHashtag(ctx context.Context, in *protopb.Hashtag) (*protopb.Hashtag, error) {
	return s.hashtagController.CreateHashtag(ctx, in)
}

func (s *Server) GetAllHashtags(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.Hashtags, error) {
	return s.hashtagController.GetAllHashtags(ctx, in)
}

func (s *Server) GetPostsByHashtag(ctx context.Context, in *protopb.Hashtag) (*protopb.PostArray, error) {
	return s.postController.GetPostsByHashtag(ctx, in)
}

/*   Highlights   */
func (s *Server) GetAllHighlights(ctx context.Context, in *protopb.RequestId) (*protopb.HighlightsArray, error) {
	return s.highlightController.GetAllHighlights(ctx, in)
}

func (s *Server) GetHighlight(ctx context.Context, in *protopb.RequestId) (*protopb.Highlight, error) {
	return s.highlightController.GetHighlight(ctx, in)
}

func (s *Server) CreateHighlight(ctx context.Context, in *protopb.Highlight) (*protopb.Highlight, error) {
	return s.highlightController.CreateHighlight(ctx, in)
}

func (s *Server) RemoveHighlight(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	return s.highlightController.RemoveHighlight(ctx, in)
}

func (s *Server) CreateHighlightStory(ctx context.Context, in *protopb.HighlightRequest) (*protopb.EmptyResponseContent, error) {
	return s.highlightController.CreateHighlightStory(ctx, in)
}

func (s *Server) RemoveHighlightStory(ctx context.Context, in *protopb.HighlightRequest) (*protopb.EmptyResponseContent, error) {
	return s.highlightController.RemoveHighlightStory(ctx, in)
}

/*   Content Complaint   */
func (s *Server) CreateContentComplaint(ctx context.Context, in *protopb.ContentComplaint) (*protopb.EmptyResponseContent, error) {
	return s.complaintController.CreateContentComplaint(ctx, in)
}

func (s *Server) GetAllContentComplaints(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.ContentComplaintArray, error) {
	return s.complaintController.GetAllContentComplaints(ctx, in)
}

func (c *Server) RejectById (ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent,error) {
	return c.complaintController.RejectById(ctx, in)
}

func (c *Server) DeleteComplaintByUserId(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	return c.complaintController.DeleteComplaintByUserId(ctx, in)
}