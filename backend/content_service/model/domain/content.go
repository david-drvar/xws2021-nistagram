package domain

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	"time"
)

type Objava struct {
	Id          string
	UserId      string
	IsAd        bool
	Type        model.PostType
	Description string
	Location    string
	CreatedAt   time.Time
	Media       []Media
	Hashtags 	[]Hashtag
}

type Post struct {
	Objava
	Comments []Comment
	Likes    []Like
	Dislikes []Like
}

type ReducedPost struct {
	Objava
	CommentsNum int32
	LikesNum    int32
	DislikesNum int32
}

type Story struct {
	Objava
	IsCloseFriends bool
}

type StoryHome struct {
	UserId    string
	Username  string
	UserPhoto string
	Stories   []Story
}

type StoriesHome struct {
	Stories []StoryHome
}

type Media struct {
	Id       string
	Type     model.MediaType
	PostId   string
	Content  string
	OrderNum int32
	Tags     []Tag
}

type Tag struct {
	MediaId  string
	UserId   string
	Username string
}

type Hashtag struct {
	Id   string
	Text string
}

type ContentComplaint struct {
	Id       string
	Category model.ComplaintCategory
	PostId   string
	Status   model.RequestStatus
	IsPost   bool
	UserId   string
}

type Collection struct {
	Id     string
	Name   string
	UserId string
	Posts  []Post
}

type Favorites struct {
	UserId       string
	Collections  []Collection
	Unclassified []Post
}

type FavoritesRequest struct {
	UserId       string
	PostId       string
	CollectionId string
}

type Like struct {
	PostId   string
	UserId   string
	IsLike   bool
	Username string
}

type Comment struct {
	Id        string
	PostId    string
	UserId    string
	Username  string
	Content   string
	CreatedAt time.Time
}

type Highlight struct {
	Id      string
	Name    string
	UserId  string
	Stories []Story
}

type HighlightRequest struct {
	UserId      string
	HighlightId string
	StoryId     string
}

/*
type RegistrationRequest struct {
	Id        string
	UserId    string
	CreatedAt time.Time
	Status    model.RequestStatus
	Username string
}

type Ad struct {
	Id string `gorm:"primaryKey"`
	Link string
	CampaignId string
	PostId string
	LinkClickNum int
}

type Campaign struct {
	Id string `gorm:"primaryKey"`
	IsOneTime bool
	StartDate time.Time //TODO
	EndDate time.Time //TODO
	PlacementNum int
	AgentId string
	IdAdCategory string
	LastUpdated time.Time //TODO
}

type CampaignInfluencerRequest struct {
	CampaignId   string `gorm:"primaryKey"`
	InfluencerId string `gorm:"primaryKey"`
	Status       model.RequestStatus
}

type ContentComplaint struct {
	Id       string `gorm:"primaryKey"`
	Category model.ComplaintCategory
	PostId   string
	Status   model.RequestStatus
}

type AdCategory struct {
	Id string `gorm:"primaryKey"`
	Name string
}

type UserAdCategories struct {
	UserId string `gorm:"primaryKey"`
	IdAdCategory string `gorm:"primaryKey"`
}

type CampaignChanges struct {
	CampaignId string `gorm:"primaryKey"`
	AdCategoryId string
	StartDate time.Time
	EndDate time.Time
	PlacementNum int
}
*/
