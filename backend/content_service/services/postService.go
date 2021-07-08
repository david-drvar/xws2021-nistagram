package services

import (
	"context"
	"errors"
	"log"

	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type PostService struct {
	postRepository    repositories.PostRepository
	commentRepository repositories.CommentRepository
	likeRepository    repositories.LikeRepository
	mediaRepository   repositories.MediaRepository
	tagRepository     repositories.TagRepository
	hashtagRepository repositories.HashtagRepository
}

func NewPostService(db *gorm.DB) (*PostService, error) {
	postRepository, err := repositories.NewPostRepo(db)
	if err != nil {
		return nil, err
	}

	commentRepository, err := repositories.NewCommentRepo(db)
	if err != nil {
		return nil, err
	}

	likeRepository, err := repositories.NewLikeRepo(db)
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

	return &PostService{
		postRepository,
		commentRepository,
		likeRepository,
		mediaRepository,
		tagRepository,
		hashtagRepository,
	}, err
}

func (service *PostService) GetAllPostsReduced(ctx context.Context, followings []string) ([]domain.ReducedPost, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPostsReduced")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []domain.ReducedPost{}

	dbPosts, err := service.postRepository.GetAllPosts(ctx, followings)
	if err != nil {
		return posts, err
	}

	for _, post := range dbPosts {
		converted, err := service.GetReducedPostData(ctx, post.Id)
		if err != nil {
			return []domain.ReducedPost{}, err
		}

		posts = append(posts, converted)
	}

	return posts, nil
}

func (service *PostService) GetAllPosts(ctx context.Context, followings []string) ([]domain.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []domain.Post{}

	dbPosts, err := service.postRepository.GetAllPosts(ctx, followings)
	if err != nil {
		return posts, err
	}

	for _, post := range dbPosts {
		converted, err := service.GetPostById(ctx, post.Id)
		if err != nil {
			return []domain.Post{}, err
		}

		posts = append(posts, converted)
	}

	return posts, nil
}

func (service *PostService) CreatePost(ctx context.Context, post *domain.Post) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if len(post.Media) == 0 {
		return errors.New("cannot create empty post")
	}

	_, err :=  service.postRepository.CreatePost(ctx, post)
	if err != nil {
		return err
	}
	users, err := grpc_common.GetUsersForNotificationEnabled(ctx, post.UserId, "IsPostNotificationEnabled")
	if err != nil {
		return errors.New("Could not create notification")
	}
	for _, u := range users.Users {
		grpc_common.CreateNotification(ctx, u.UserId, post.UserId, "Post", post.Id)
	}
	return nil
}

func (service *PostService) RemovePost(ctx context.Context, id string, userId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemovePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.postRepository.RemovePost(ctx, id, userId)
}

func (service *PostService) GetPostById(ctx context.Context, id string) (domain.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbPost, err := service.postRepository.GetPostById(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}

	dbComments, err := service.commentRepository.GetCommentsForPost(ctx, dbPost.Id)
	if err != nil {
		return domain.Post{}, err
	}

	dbHashtags, err := service.hashtagRepository.GetPostHashtags(ctx, dbPost.Id)
	if err != nil {
		return domain.Post{}, err
	}

	hashtags := []domain.Hashtag{}
	for _, hashtag := range dbHashtags {
		hashtags = append(hashtags, domain.Hashtag{
			Id:   hashtag.Id,
			Text: hashtag.Text,
		})
	}

	comments := []domain.Comment{}
	for _, comment := range dbComments {
		username, err := grpc_common.GetUsernameById(ctx, comment.UserId)
		if err == nil {
			comments = append(comments, comment.ConvertToDomain(username))
		}
	}

	dbLikes, err := service.likeRepository.GetLikesForPost(ctx, dbPost.Id, true)
	if err != nil {
		return domain.Post{}, err
	}
	likes := []domain.Like{}
	for _, like := range dbLikes {
		username, err := grpc_common.GetUsernameById(ctx, like.UserId)
		if err == nil {
			likes = append(likes, like.ConvertToDomain(username))
		}
	}

	dbDislikes, err := service.likeRepository.GetLikesForPost(ctx, dbPost.Id, false)
	if err != nil {
		return domain.Post{}, err
	}
	dislikes := []domain.Like{}
	for _, dislike := range dbDislikes {
		username, err := grpc_common.GetUsernameById(ctx, dislike.UserId)
		if err == nil {
			dislikes = append(dislikes, dislike.ConvertToDomain(username))
		}
	}

	hashtags, err = service.hashtagRepository.GetPostHashtags(ctx, dbPost.Id)
	if err != nil {
		return domain.Post{}, err
	}


	dbMedia, err := service.mediaRepository.GetMediaForPost(ctx, dbPost.Id)
	if err != nil {
		return domain.Post{}, err
	}
	media := []domain.Media{}
	for _, single := range dbMedia {
		tags, err := service.tagRepository.GetTagsForMedia(ctx, single.Id)
		if err != nil {
			return domain.Post{}, err
		}

		for index, tag := range tags {
			username, err := grpc_common.GetUsernameById(ctx, tag.UserId)
			if username == "" || err != nil {
				return domain.Post{}, errors.New("cannot retrieve tags")
			}
			tags[index].Username = username
		}

		converted, err := single.ConvertToDomain(tags)
		if err != nil {
			return domain.Post{}, err
		}

		media = append(media, converted)
	}

	post := dbPost.ConvertToDomain(comments, likes, dislikes, media, hashtags)

	return post, nil
}
func (service *PostService) GetReducedPostData(ctx context.Context, postId string) (domain.ReducedPost, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetReducedPostData")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	commentsNum, err := service.commentRepository.GetCommentsNumForPost(ctx, postId)
	if err != nil {
		return domain.ReducedPost{}, errors.New("unable to retrieve posts comments")
	}

	likes, err := service.likeRepository.GetLikesNumForPost(ctx, postId, true)
	if err != nil {
		return domain.ReducedPost{}, errors.New("unable to retrieve posts likes")
	}

	dislikes, err := service.likeRepository.GetLikesNumForPost(ctx, postId, false)
	if err != nil {
		return domain.ReducedPost{}, errors.New("unable to retrieve posts dislikes")
	}

	media, err := service.mediaRepository.GetMediaForPost(ctx, postId)
	if err != nil {
		return domain.ReducedPost{}, errors.New("unable to retrieve posts media")
	}

	convertedMedia := []domain.Media{}
	for _, single := range media {
		tags, err := service.tagRepository.GetTagsForMedia(ctx, single.Id)
		if err != nil {
			return domain.ReducedPost{}, errors.New("unable to retrieve media tags")
		}

		for index, tag := range tags {
			username, err := grpc_common.GetUsernameById(ctx, tag.UserId)
			if username == "" || err != nil {
				return domain.ReducedPost{}, errors.New("cannot retrieve tags")
			}
			tags[index].Username = username
		}

		converted, err := single.ConvertToDomain(tags)
		if err != nil {
			return domain.ReducedPost{}, errors.New("unable to convert media")
		}

		convertedMedia = append(convertedMedia, converted)
	}

	if err != nil {
		return domain.ReducedPost{}, errors.New("unable to convert posts media")
	}

	post, err := service.postRepository.GetPostById(ctx, postId)
	if err != nil {
		return domain.ReducedPost{}, err
	}

	retVal := post.ConvertToDomainReduced(commentsNum, likes, dislikes, convertedMedia)
	return retVal, nil
}

func (service *PostService) SearchContentByLocation(ctx context.Context, location string) ([]domain.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SearchContentByLocation")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []domain.Post{}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protopb.NewPrivacyClient(conn)

	dbPosts, err := service.postRepository.GetPostsByLocation(ctx, location)
	if err != nil {
		return posts, err
	}

	for _, post := range dbPosts {
		converted, err := service.GetPostById(ctx, post.Id)
		if err != nil {
			return []domain.Post{}, err
		}

		posts = append(posts, converted)
	}

	finalPosts := []domain.Post{}

	//call user service to check if users profile is public
	for _, post := range posts {
		privacyRequest := protopb.PrivacyRequest{
			UserId: post.UserId,
		}
		response, err := c.CheckUserProfilePublic(context.Background(), &privacyRequest)
		if err != nil {
			log.Fatalf("Error when calling CheckUserProfilePublic: %s", err)
		}
		if response.Response {
			res, _ := grpc_common.CheckIsActive(ctx, post.UserId)
			if res {
				finalPosts = append(finalPosts, post)
			}

		}
	}

	//todo proveri da li su se useri blokirali medjusobno - ne moze u bilo kom smeru
	//todo ako neregistrovan pretrazuje onda nemoj proveravati

	return finalPosts, nil
}

func (service *PostService) GetPostsByHashtag(ctx context.Context, text string) ([]domain.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsByHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	hashtag, err := service.hashtagRepository.GetHashtagByText(ctx, text)
	if err != nil {
		log.Fatalf("Error when calling GetPostsByHashtag: %s", err)
	}

	var posts []domain.Post

	postIds, _ := service.hashtagRepository.GetPostIdsByHashtag(ctx, persistence.Hashtag{Id: hashtag.Id, Text: hashtag.Text})
	for _, postId := range postIds {
		post, err := service.GetPostById(ctx, postId)
		if err != nil {
			log.Fatalf("Error when calling GetPostById: %s", err)
		}
		posts = append(posts, post)
	}

	var postsWithPublicAccess []domain.Post

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(":8091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protopb.NewPrivacyClient(conn)

	// for debugging microservice communication
	//privacyRequest := protopb.PrivacyRequest{
	//	UserId: "0",
	//}
	//response, err := c.CheckUserProfilePublic(context.Background(), &privacyRequest)
	//if err != nil {
	//	log.Fatalf("Error when calling CheckUserProfilePublic: %s", err)
	//}
	//print(response.Response)

	//check if user account is public
	for _, post := range posts {
		privacyRequest := protopb.PrivacyRequest{
			UserId: post.UserId,
		}
		response, err := c.CheckUserProfilePublic(context.Background(), &privacyRequest)
		if err != nil {
			log.Fatalf("Error when calling CheckUserProfilePublic: %s", err)
		}
		if response.Response {
			res, _ := grpc_common.CheckIsActive(ctx, post.UserId)
			if res {
				postsWithPublicAccess = append(postsWithPublicAccess, post)
			}
		}
	}

	return postsWithPublicAccess, nil
}

func (service *PostService) GetPostsForUser(ctx context.Context, id string) ([]domain.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []domain.Post{}

	res, err := grpc_common.CheckIsActive(ctx, id)
	if err != nil {
		return nil, err
	}else if res == false {
		return nil, errors.New("User is not active!")
	}

	dbPosts, err := service.postRepository.GetPostsForUser(ctx, id)
	if err != nil {
		return posts, err
	}

	for _, post := range dbPosts {
		converted, err := service.GetPostById(ctx, post.Id)
		if err != nil {
			return []domain.Post{}, err
		}

		posts = append(posts, converted)
	}

	return posts, nil
}
