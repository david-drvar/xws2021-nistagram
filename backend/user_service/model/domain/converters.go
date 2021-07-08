package domain

import (
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (l LoginRequest) ConvertFromGrpc(request *protopb.LoginRequest) LoginRequest {
	return LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}
}

func (p Password) ConvertFromGrpc(pass *protopb.Password) Password {
	newPass := Password{
		OldPassword:      pass.OldPassword,
		NewPassword:      pass.NewPassword,
		RepeatedPassword: pass.NewPassword,
		Id:               pass.Id,
	}

	check, err := newPass.CheckAllFields()
	if !check || err != nil {
		return Password{}
	}

	return newPass
}

func (result InfluencerSearchResult) ConvertToGrpc() *protopb.InfluencerSearch {
	return &protopb.InfluencerSearch{
		Id:              result.Id,
		FirstName:       result.FirstName,
		LastName:        result.LastName,
		Username:        result.Username,
		ProfilePhoto:    result.ProfilePhoto,
		IsProfilePublic: result.IsProfilePublic,
	}
}

func (u User) ConvertToGrpc() *protopb.UsersDTO {
	return &protopb.UsersDTO{
		Id:              u.Id,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Email:           u.Email,
		Username:        u.Username,
		Role:            u.Role.String(),
		Birthdate:       timestamppb.New(u.BirthDate),
		ProfilePhoto:    u.ProfilePhoto,
		PhoneNumber:     u.PhoneNumber,
		Sex:             u.Sex,
		IsActive:        u.IsActive,
		Website:         u.Website,
		Biography:       u.Biography,
		Category:        model.GetUserCategoriesString(u.Category),
		ResetCode:       u.ResetCode,
		ApprovedAccount: u.ApprovedAccount,
		TokenEnd:        timestamppb.New(u.TokenEnd),
	}
}

func (req *RegistrationRequest) ConvertToGrpc() *protopb.RegistrationRequest {
	return &protopb.RegistrationRequest{
		CreatedAt: timestamppb.New(req.CreatedAt),
		LastName:  req.LastName,
		Email:     req.Email,
		Username:  req.Username,
		FirstName: req.FirstName,
		Id:        req.Id,
		UserId:    req.UserId,
		Website:   req.Website,
		Status:    model.ToStringRequestStatus(req.Status),
	}
}

func (u User) ConvertFromGrpc(user *protopb.UsersDTO) User {
	newUser := User{
		Id:              user.Id,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Username:        user.Username,
		Role:            model.GetUserRole(user.Role),
		BirthDate:       user.Birthdate.AsTime(),
		ProfilePhoto:    user.ProfilePhoto,
		PhoneNumber:     user.PhoneNumber,
		Sex:             user.Sex,
		IsActive:        user.IsActive,
		Website:         user.Website,
		Biography:       user.Biography,
		Category:        model.GetUserCategories(user.Category),
		ResetCode:       user.ResetCode,
		ApprovedAccount: user.ApprovedAccount,
		TokenEnd:        user.TokenEnd.AsTime(),
	}

	//check, err := newUser.CheckAllFields()
	//if !check || err != nil {
	//	return User{}
	//}

	return newUser
}

func (u *User) GenerateUserDTO(user persistence.User, userAdditionalInfo persistence.UserAdditionalInfo) *User {
	if u == nil {
		u = &User{}
	}
	return &User{
		Id:              user.Id,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Username:        user.Username,
		Role:            user.Role,
		BirthDate:       user.BirthDate,
		ProfilePhoto:    user.ProfilePhoto,
		PhoneNumber:     user.PhoneNumber,
		IsActive:        user.IsActive,
		Sex:             user.Sex,
		Biography:       userAdditionalInfo.Biography,
		Category:        userAdditionalInfo.Category,
		Website:         userAdditionalInfo.Website,
		ResetCode:       user.ResetCode,
		ApprovedAccount: user.ApprovedAccount,
		TokenEnd:        user.TokenEnd,
	}
}

func (u VerificationRequest) ConvertFromGrpc(verificationRequest *protopb.VerificationRequest) VerificationRequest {
	newVerificationRequest := VerificationRequest{
		Id:            verificationRequest.Id,
		UserId:        verificationRequest.UserId,
		FirstName:     verificationRequest.FirstName,
		LastName:      verificationRequest.LastName,
		Category:      model.GetCategory(verificationRequest.Category),
		CreatedAt:     verificationRequest.CreatedAt.AsTime(),
		DocumentPhoto: verificationRequest.DocumentPhoto,
		Status:        model.GetRequestStatus(verificationRequest.Status),
	}

	return newVerificationRequest
}

func (verificationRequest VerificationRequest) ConvertToGrpc() *protopb.VerificationRequest {
	return &protopb.VerificationRequest{
		Id:            verificationRequest.Id,
		UserId:        verificationRequest.UserId,
		DocumentPhoto: verificationRequest.DocumentPhoto,
		Category:      model.GetUserCategoriesString(verificationRequest.Category),
		FirstName:     verificationRequest.FirstName,
		LastName:      verificationRequest.LastName,
		Status:        model.ToStringRequestStatus(verificationRequest.Status),
		CreatedAt:     timestamppb.New(verificationRequest.CreatedAt),
	}
}
func (n *UserNotification) ConvertFromGrpc(notification protopb.CreateNotificationRequest) *UserNotification {
	if n == nil {
		n = &UserNotification{}
	}
	return &UserNotification{
		UserId:           notification.UserId,
		CreatorId:        notification.CreatorId,
		NotificationType: notification.Type,
		ContentId:        notification.ContentId,
	}
}
