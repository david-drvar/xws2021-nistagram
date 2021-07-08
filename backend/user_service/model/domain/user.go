package domain

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"regexp"
	"time"
)

type Password struct {
	OldPassword      string
	NewPassword      string
	RepeatedPassword string
	Id               string
}

type User struct {
	Id              string
	FirstName       string
	LastName        string
	Email           string
	Username        string
	Role            model.UserRole
	BirthDate       time.Time
	ProfilePhoto    string
	PhoneNumber     string
	Sex             string
	IsActive        bool
	Biography       string
	Website         string
	Category        model.UserCategory
	ResetCode       string
	ApprovedAccount bool
	TokenEnd        time.Time
}

type InfluencerSearchResult struct {
	Id              string
	FirstName       string
	LastName        string
	Username        string
	Role            model.UserRole
	ProfilePhoto    string
	IsProfilePublic bool
}

type LoginRequest struct {
	Email    string
	Password string
}

type RegistrationRequest struct {
	Id        string
	UserId    string
	CreatedAt time.Time
	Status    model.RequestStatus
	Username  string
	Email     string
	Website   string
	FirstName string
	LastName  string
}

type UserNotification struct {
	UserId           string
	CreatorId        string
	NotificationType string
	IsRead           bool
	ContentId        string
}

func (user *User) CheckValidation() (bool, error) {

	match, err := regexp.MatchString("^[a-zA-Z ,.'-]+$", user.FirstName)
	if !match {
		return false, err
	}

	return true, nil
}

type VerificationRequest struct {
	Id            string
	UserId        string
	DocumentPhoto string
	Status        model.RequestStatus
	CreatedAt     time.Time
	Category      model.UserCategory
	FirstName     string
	LastName      string
}

type CampaignRequest struct {
	Id           string
	AgentId      string
	InfluencerId string
	CampaignId   string
	Status       model.RequestStatus
	PostAt       time.Time
}
