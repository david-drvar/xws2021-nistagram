package persistence

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/util"
	"github.com/david-drvar/xws2021-nistagram/content_service/util/images"
	uuid "github.com/satori/go.uuid"
	"time"
)

func (p Post) ConvertToDomain(comments []domain.Comment, likes []domain.Like, dislikes []domain.Like, media []domain.Media, hashtags []domain.Hashtag) domain.Post{
	return domain.Post{
		Objava:   domain.Objava{
			Id:          p.Id,
			UserId:      p.UserId,
			IsAd:        p.IsAd,
			Type:        p.Type,
			Description: p.Description,
			Location:    p.Location,
			CreatedAt:   p.CreatedAt,
			Media:       media,
			Hashtags:    hashtags,
		},
		Comments: comments,
		Likes:    likes,
		Dislikes: dislikes,
	}
}

func (p Post) ConvertToDomainReduced(commentsNum int, likesNum int, dislikesNum int, media []domain.Media) domain.ReducedPost{
	return domain.ReducedPost{
		Objava:   domain.Objava{
			Id:          p.Id,
			UserId:      p.UserId,
			IsAd:        p.IsAd,
			Type:        p.Type,
			Description: p.Description,
			Location:    p.Location,
			CreatedAt:   p.CreatedAt,
			Media:       media,
		},
		CommentsNum: 	int32(commentsNum),
		LikesNum:    	int32(likesNum),
		DislikesNum: 	int32(dislikesNum),
	}
}

func (p Post) ConvertToPersistence(post domain.Post) Post {
	newPost := Post{
		Id:          uuid.NewV4().String(),
		UserId:      post.UserId,
		IsAd:        post.IsAd,
		Type:        post.Type,
		Description: post.Description,
		Location:    post.Location,
	}

	if post.CreatedAt.Equal(time.Time{}) {
		newPost.CreatedAt = time.Now()
	}else{
		newPost.CreatedAt = post.CreatedAt
	}

	return newPost
}

func (s Story) ConvertToPersistence(story domain.Story) Story {
	newStory := Story{
		Post: Post {
			Id:          uuid.NewV4().String(),
			UserId:      story.UserId,
			IsAd:        story.IsAd,
			Type:        story.Type,
			Description: story.Description,
			Location:    story.Location,
			CreatedAt:   time.Now(),
		},
		IsCloseFriends: story.IsCloseFriends,
	}

	if story.CreatedAt.Equal(time.Time{}){
		newStory.CreatedAt = time.Now()
	}else{
		newStory.CreatedAt = story.CreatedAt
	}

	return newStory
}

func (s Story) ConvertToDomain(media []domain.Media, hashtags []domain.Hashtag) domain.Story {
	return domain.Story{
		Objava: domain.Objava {
			Id:          s.Id,
			UserId:      s.UserId,
			IsAd:        s.IsAd,
			Type:        s.Type,
			Description: s.Description,
			Location:    s.Location,
			CreatedAt:   s.CreatedAt,
			Media: 		 media,
			Hashtags:    hashtags,
		},
		IsCloseFriends: s.IsCloseFriends,
	}
}

func (c Comment) ConvertToDomain(username string) domain.Comment {
	return domain.Comment{
		Id:		   c.Id,
		PostId:    c.PostId,
		UserId:    c.UserId,
		Username:  username,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
}

func (c *Comment) ConvertToPersistence(comment domain.Comment) *Comment{
	if c == nil { c = &Comment{} }
	return &Comment{
		Id:		   uuid.NewV4().String(),
		PostId:    comment.PostId,
		UserId:    comment.UserId,
		Content:   comment.Content,
		CreatedAt: time.Now(),
	}
}

func (l Like) ConvertToDomain(username string) domain.Like {
	return domain.Like{
		PostId: l.PostId,
		UserId: l.UserId,
		IsLike: l.IsLike,
		Username: username,
	}
}

func (l *Like) ConvertToPersistence(like domain.Like) *Like{
	if l == nil { l = &Like{} }
	return &Like{
		PostId:    	like.PostId,
		UserId:    	like.UserId,
		IsLike:		like.IsLike,
	}
}

func (m *Media) ConvertToDomain(tags []domain.Tag) (domain.Media, error){
	if m == nil { m = &Media{} }
	filename, err := images.LoadImageToBase64(util.GetContentLocation(m.Filename))
	if err != nil {
		return domain.Media{}, err
	}
	return domain.Media{
		Id:       m.Id,
		Type:     m.Type,
		PostId:   m.PostId,
		Content:  filename,
		OrderNum: int32(m.OrderNum),
		Tags:	  tags,
	}, nil
}

func (m *Media) ConvertToPersistence(media domain.Media, filename string) *Media{
	if m == nil { m = &Media{} }
	return &Media{
		Id:       uuid.NewV4().String(),
		Type:     media.Type,
		PostId:   media.PostId,
		Filename: filename,
		OrderNum: int(media.OrderNum),
	}
}

func (t Tag) ConvertToDomain(username string) domain.Tag {
	return domain.Tag{
		MediaId:   t.MediaId,
		UserId:    t.UserId,
		Username:  username,
	}
}

func (t *Tag) ConvertToPersistence(tag domain.Tag) *Tag{
	if t == nil { t = &Tag{} }
	return &Tag{
		MediaId:    tag.MediaId,
		UserId:    	tag.UserId,
	}
}

func (f Favorites) ConvertToDomain(collections []domain.Collection, unclassified []domain.Post) domain.Favorites {
	return domain.Favorites{
		UserId:       f.UserId,
		Collections:  collections,
		Unclassified: unclassified,
	}
}

func (f *Favorites) ConvertToPersistence(favorites domain.FavoritesRequest) *Favorites{
	if f == nil { f = &Favorites{} }
	return &Favorites{
		PostId:       favorites.PostId,
		CollectionId: favorites.CollectionId,
		UserId: 	  favorites.UserId,
	}
}

func (c Collection) ConvertToDomain(posts []domain.Post) domain.Collection {
	return domain.Collection{
		Id:     c.Id,
		Name:   c.Name,
		UserId: c.UserId,
		Posts:  posts,
	}
}

func (c *Collection) ConvertToPersistence(collection domain.Collection) *Collection{
	if c == nil { c = &Collection{} }
	return &Collection{
		Id:     uuid.NewV4().String(),
		Name:   collection.Name,
		UserId: collection.UserId,
	}
}

/* Highlights */
func (h Highlight) ConvertToDomain(stories []domain.Story) domain.Highlight {
	return domain.Highlight{
		Id:      h.Id,
		Name:    h.Name,
		UserId:  h.UserId,
		Stories: stories,
	}
}

func (h *Highlight) ConvertToPersistence(highlight domain.Highlight) *Highlight{
	if h == nil { h = &Highlight{} }
	return &Highlight{
		Id:     uuid.NewV4().String(),
		UserId: highlight.UserId,
		Name:   highlight.Name,
	}
}

func (c *HighlightStory) ConvertToPersistence(request domain.HighlightRequest) *HighlightStory{
	if c == nil { c = &HighlightStory{} }
	return &HighlightStory{
		HighlightId: request.HighlightId,
		StoryId:     request.StoryId,
	}
}

func (a Ad) ConvertToPersistence(ad domain.Ad) Ad{
	return Ad{
		Id:         ad.Id,
		Link:       ad.Link,
		CampaignId: ad.CampaignId,
		PostId:     ad.Post.Id,
		LinkClicks: ad.LinkClicks,
		Type:       ad.Post.Type.String(),
	}
}

func (a Ad) ConvertToDomain(comments []domain.Comment, likes []domain.Like, dislikes []domain.Like, content domain.Objava) domain.Ad {
	return domain.Ad{
		Id:         a.Id,
		Link:       a.Link,
		CampaignId: a.CampaignId,
		Post:       domain.Post{
			Objava:   content,
			Comments: comments,
			Likes:    likes,
			Dislikes: dislikes,
		},
		LinkClicks: a.LinkClicks,
	}
}

func (c Campaign) ConvertToDomain(ads []domain.Ad, category domain.AdCategory) domain.Campaign{
	return domain.Campaign{
		Id:           c.Id,
		Name:		  c.Name,
		IsOneTime:    c.IsOneTime,
		StartDate:    c.StartDate,
		EndDate:      c.EndDate,
		Placements:   c.Placements,
		AgentId:      c.AgentId,
		Category: 	  category,
		LastUpdated:  time.Now(),
		Ads: 		  ads,
		Type: 		  model.PostType(c.Type),
	}
}

func (c Campaign) ConvertToPersistence(camp domain.Campaign) Campaign{
	return Campaign{
		Id:           camp.Id,
		Name:		  camp.Name,
		IsOneTime:    camp.IsOneTime,
		StartDate:    camp.StartDate,
		EndDate:      camp.EndDate,
		Placements:   camp.Placements,
		AgentId:      camp.AgentId,
		AdCategoryId: camp.Category.Id,
		LastUpdated:  camp.LastUpdated,
		Type: 		  camp.Type.String(),
	}
}

func (ac AdCategory) ConvertToDomain() domain.AdCategory{
	return domain.AdCategory{
		Id:   ac.Id,
		Name: ac.Name,
	}
}

func (ac AdCategory) ConvertToPersistence(category domain.AdCategory) AdCategory{
	return AdCategory{
		Id:   category.Id,
		Name: category.Name,
	}
}