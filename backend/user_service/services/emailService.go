package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"gorm.io/gorm"
	"math/rand"
	"net/smtp"
	"time"
)

type EmailService struct {
	repository repositories.UserRepository
}

func NewEmailService(db *gorm.DB, redis *saga.RedisServer) (*EmailService, error) {
	repository, err := repositories.NewUserRepo(db, redis)

	return &EmailService{
		repository: repository,
	}, err
}

func (service *EmailService) SendEmail(ctx context.Context, in string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserProfile")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var userDomain domain.User
	userDomain, _ = service.repository.GetUserByEmail(in)
	if userDomain.Id == "" {
		return false, errors.New("email not exist")
	}

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, 8)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Println(string(s))

	userDomain.ResetCode = string(s)
	userDomain.TokenEnd = time.Now().AddDate(0, 0, 1)
	service.repository.UpdateUserProfile(ctx, userDomain)

	from := "bsep2021@gmail.com"
	password := "BSEP2021"
	to := []string{
		"t.kovacevic98@gmail.com",
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Hi,\nwe received a request to reset your password.Your old password has been locked for security reasons.\nTo unlock your profile you must verify your identity.\n\nYour password reset code is:" + string(s))
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println("Email Sent Successfully!")

	return true, nil
}
