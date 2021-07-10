package persistence

import (
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type User struct {
	Id              string `gorm:"primaryKey"`
	FirstName       string
	LastName        string
	Email           string `gorm:"unique"`
	Username        string `gorm:"unique"`
	Password        string
	Role            model.UserRole
	BirthDate       time.Time
	ProfilePhoto    string
	PhoneNumber     string
	Sex             string
	IsActive        bool
	ResetCode       string
	ApprovedAccount bool
	TokenEnd        time.Time
}

func (u User) ConvertToGrpc() *protopb.User {
	return &protopb.User{
		Id:              u.Id,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Email:           u.Email,
		Username:        u.Username,
		Password:        u.Password,
		Role:            u.Role.String(),
		Birthdate:       timestamppb.New(u.BirthDate),
		ProfilePhoto:    u.ProfilePhoto,
		PhoneNumber:     u.PhoneNumber,
		Sex:             u.Sex,
		IsActive:        u.IsActive,
		ResetCode:       u.ResetCode,
		ApprovedAccount: u.ApprovedAccount,
		TokenEnd:        timestamppb.New(u.TokenEnd),
	}
}

func (u *User) ConvertFromGrpc(user *protopb.User) *User {
	newUser := &User{
		Id:              user.Id,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Username:        user.Username,
		Password:        user.Password,
		Role:            model.GetUserRole(user.Role),
		BirthDate:       user.Birthdate.AsTime(),
		ProfilePhoto:    user.ProfilePhoto,
		PhoneNumber:     user.PhoneNumber,
		Sex:             user.Sex,
		IsActive:        user.IsActive,
		ResetCode:       user.ResetCode,
		ApprovedAccount: user.ApprovedAccount,
		TokenEnd:        user.TokenEnd.AsTime(),
	}

	//check, err := newUser.CheckAllFields()
	//if !check || err != nil {
	//	return nil
	//}

	return newUser
}

type UserAdditionalInfo struct {
	Id        string `gorm:"primaryKey"`
	Biography string
	Website   string
	Category  model.UserCategory
}

func (u UserAdditionalInfo) ConvertFromGrpc(user *protopb.User) *UserAdditionalInfo {
	return &UserAdditionalInfo{
		Id:        user.Id,
		Biography: user.Biography,
		Website:   user.Website,
	}
}

type Privacy struct {
	UserId          string `gorm:"primaryKey"`
	IsProfilePublic bool   `gorm:"type:boolean"`
	IsDMPublic      bool   `gorm:"type:boolean"`
	IsTagEnabled    bool   `gorm:"type:boolean"`
}

func (privacy *Privacy) ConvertFromGrpc(p *protopb.PrivacyMessage) *Privacy {
	return &Privacy{
		UserId:          p.Id,
		IsDMPublic:      p.IsDmPublic,
		IsProfilePublic: p.IsProfilePublic,
		IsTagEnabled:    p.IsTagEnabled,
	}
}

type BlockedUsers struct {
	UserId        string `gorm:"primaryKey"`
	BlockedUserId string `gorm:"primaryKey"`
}

func (block *BlockedUsers) ConvertFromGrpc(b *protopb.Block) *BlockedUsers {
	return &BlockedUsers{
		UserId:        b.UserId,
		BlockedUserId: b.BlockedUserId,
	}
}

type Followers struct {
	UsedId                string `gorm:"primaryKey"`
	FollowerId            string `gorm:"primaryKey"`
	IsMuted               bool
	IsCloseFriend         bool
	IsApprovedRequest     bool
	IsNotificationEnabled bool
}

type VerificationRequest struct {
	Id            string `gorm:"primaryKey"`
	UserId        string
	DocumentPhoto string
	Status        model.RequestStatus
	CreatedAt     time.Time
	Category      model.UserCategory
}

type CampaignRequest struct {
	Id           string `gorm:"primaryKey"`
	AgentId      string
	InfluencerId string
	CampaignId   string
	Status       model.RequestStatus
	PostAt       time.Time
}

type APIKeys struct {
	UserId string `gorm:"primaryKey"`
	APIKey string
}

type RegistrationRequest struct {
	Id        string `gorm:"primaryKey"`
	UserId    string
	CreatedAt time.Time
	Status    model.RequestStatus
}

type UserNotification struct {
	NotificationId string `gorm:"primaryKey"`
	UserId         string
	CreatorId      string
	Text           string
	Type           string
	IsRead         bool
	ContentId      string
	CreatedAt      time.Time
}

func (notification *UserNotification) ConvertFromGrpc(n *protopb.Notification) *UserNotification {
	return &UserNotification{
		NotificationId: n.Id,
		UserId:         n.UserId,
		CreatorId:      n.CreatorId,
		Text:           n.Text,
		IsRead:         n.IsRead,
	}
}
