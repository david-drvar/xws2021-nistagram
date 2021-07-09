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


type RegistrationRequest struct {
	Id        string
	UserId    string
	CreatedAt time.Time
	Status    model.RequestStatus
	Username string
}

type Ad struct {
	Id 			 string
	Link 		 string
	CampaignId   string
	Post 		 Post
	LinkClicks   int
}

type Campaign struct {
	Id 			 string
	Name		 string
	IsOneTime 	 bool
	StartDate 	 time.Time
	EndDate 	 time.Time
	StartTime	 int
	EndTime		 int
	Placements   int
	AgentId 	 string
	Category	 AdCategory
	LastUpdated  time.Time
	Ads			 []Ad
	Type		 model.PostType
}

type CampaignInfluencerRequest struct {
	Id           string `gorm:"primaryKey"`
	AgentId      string
	InfluencerId string
	CampaignId   string
	Status       model.RequestStatus
	PostAt       time.Time
}

/*
type ContentComplaint struct {
	Id       string `gorm:"primaryKey"`
	Category model.ComplaintCategory
	PostId   string
	Status   model.RequestStatus
}*/

type AdCategory struct {
	Id   string
	Name string
}

type CampaignChanges struct {
	CampaignId 	 string
	Name		 string
	AdCategoryId string
	StartDate 	 time.Time
	EndDate		 time.Time
	PlacementNum int
}

type CampaignStats struct {
	Id 			 string
	Name		 string
	IsOneTime 	 bool
	StartDate 	 time.Time
	EndDate 	 time.Time
	StartTime	 int
	EndTime		 int
	Placements   int
	Category	 string
	Type		 string
	Influencers	 []InfluencerStats
	Likes	 	 int
	Dislikes 	 int
	Comments 	 int
	Clicks   	 int
}

type AdStats struct {
	Id 		 string
	Media    []string
	Type	 string
	Hashtags []string
	Location string
	Likes	 int
	Dislikes int
	Comments int
	Clicks   int
}

type InfluencerStats struct {
	Id				string
	Username		string
	Ads			    []AdStats
	TotalLikes		int
	TotalDislikes	int
	TotalComments	int
	TotalClicks		int
}