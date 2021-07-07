package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/security"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"

	//"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"github.com/david-drvar/xws2021-nistagram/user_service/util"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/encryption"
	//"github.com/david-drvar/xws2021-nistagram/user_service/util/setup"

	//"github.com/david-drvar/xws2021-nistagram/user_service/util/setup"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	userRepository    repositories.UserRepository
	privacyRepository repositories.PrivacyRepository
}

func NewUserService(db *gorm.DB, redis *saga.RedisServer) (*UserService, error) {
	userRepository, err := repositories.NewUserRepo(db, redis)
	privacyRepository, err := repositories.NewPrivacyRepo(db)

	return &UserService{
		userRepository, privacyRepository,
	}, err
}

func (service *UserService) GetUsername(ctx context.Context, userId string) (string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsername")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbUser, err := service.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return "", err
	}

	return dbUser.Username, nil
}

func (service *UserService) GetUser(ctx context.Context, requestedUserId string) (domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	res, err := service.CheckIsActive(ctx, requestedUserId)
	if err != nil {
		return domain.User{}, err
	}else if res == false {
		return domain.User{}, errors.New("User is not active!")
	}

	dbUser, err := service.userRepository.GetUserById(ctx, requestedUserId)
	if err != nil {
		return domain.User{}, err
	}

	additionalInfo, err := service.userRepository.GetUserAdditionalInfoById(ctx, requestedUserId)
	if err != nil {
		return domain.User{}, err
	}

	//TODO Get user's additional info
	var converted *domain.User
	converted = converted.GenerateUserDTO(dbUser, additionalInfo)

	return *converted, nil
}

func (service *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.userRepository.GetAllUsers(ctx)
}

func (service *UserService) CreateUser(ctx context.Context, user *persistence.User) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if security.CheckBlacklistedPassword(user.Password) {
		return errors.New("password is among blacklisted passwords")
	}

	user.Password = encryption.HashAndSalt([]byte(user.Password))
	return service.userRepository.CreateUser(ctx, user)
}

func (service *UserService) CreateUserWithAdditionalInfo(ctx context.Context, user *persistence.User, userAdditionalInfo *persistence.UserAdditionalInfo) (*domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if security.CheckBlacklistedPassword(user.Password) && user.Password != "" {
		return nil, errors.New("password is among blacklisted passwords")
	}

	user.Password = encryption.HashAndSalt([]byte(user.Password))
	userResult, err := service.userRepository.CreateUserWithAdditionalInfo(ctx, user, userAdditionalInfo)
	if err != nil {
		return nil, errors.New("cannot create user")
	}

	//create user node in graph database

	//todo ovde pozivas SAGA-u
	//todo popravi userId
	//m := saga.Message{Service: saga.ServiceRecommendation, SenderService: saga.ServiceUser, Action: saga.ActionStart, UserId: user.Id}
	//service.redisServer.Orchestrator.Next(saga.RecommendationChannel, saga.ServiceRecommendation, m)

	//var conn *grpc.ClientConn
	//conn, err = grpc.Dial(":8095", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %s", err)
	//}
	//defer conn.Close()
	//
	//c := protopb.NewFollowersClient(conn)
	//
	//createUserRequest := protopb.CreateUserRequestFollowers{
	//	User: &protopb.UserFollowers{
	//		UserId: user.Id,
	//	},
	//}
	//
	//_, err = c.CreateUser(context.Background(), &createUserRequest)
	//if err != nil {
	//	log.Fatalf("could not create node user: %s", err)
	//}

	return userResult, nil
}

func (service *UserService) LoginUser(ctx context.Context, request domain.LoginRequest) (persistence.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "LoginUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if request.Email == "" || request.Password == "" {
		return persistence.User{}, errors.New("empty login request")
	}

	return service.userRepository.LoginUser(ctx, request)
}

func (service *UserService) UpdateUserProfile(ctx context.Context, userDTO domain.User) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserProfile")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if userDTO.Username == "" || userDTO.Email == "" {
		return false, errors.New("username or email can not be empty string")
	}

	return service.userRepository.UpdateUserProfile(ctx, userDTO)
}

func (service *UserService) UpdateUserPassword(ctx context.Context, password domain.Password) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if password.NewPassword != password.RepeatedPassword {
		return false, errors.New("Passwords do not match!")
	}

	if security.CheckBlacklistedPassword(password.NewPassword) {
		return false, errors.New("password is among blacklisted passwords")
	}

	_, err := service.userRepository.UpdateUserPassword(ctx, password)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (service *UserService) SearchUsersByUsernameAndName(ctx context.Context, user *domain.User, userWhoSearchedId string) ([]domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SearchUsersByUsernameAndName")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	users, err := service.userRepository.SearchUsersByUsernameAndName(ctx, user)
	if err != nil {
		return nil, err
	}

	var finalUsers []domain.User

	for _, user := range users {
		blocked, err := service.privacyRepository.CheckIfBlocked(ctx, user.Id, userWhoSearchedId)
		if err != nil {
			return nil, err
		}
		if blocked == false {
			finalUsers = append(finalUsers, user)
		}
	}

	return finalUsers, nil

}

func (service *UserService) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserByEmail")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.userRepository.GetUserByEmail(email)
}

func (service *UserService) ValidateResetCode(ctx context.Context, resetCode string, email string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user domain.User
	user, _ = service.GetUserByEmail(ctx, email)
	if user.ResetCode != resetCode {
		return false, errors.New("wrong reset code!")
	}

	today := time.Now()
	if today.After(user.TokenEnd) {
		return false, errors.New("Reset code expired!")
	}
	return true, nil
}

func (service *UserService) ChangeForgottenPass(ctx context.Context, password domain.Password) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if password.NewPassword != password.RepeatedPassword {
		return false, errors.New("Passwords do not match!")
	}

	_, err := service.userRepository.ChangeForgottenPass(ctx, password)
	if err != nil {
		return false, err
	}

	return true, nil
}
func (service *UserService) ApproveAccount(ctx context.Context, password domain.Password) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if password.NewPassword != password.RepeatedPassword {
		return false, errors.New("Passwords do not match!")
	}

	_, err := service.userRepository.ApproveAccount(ctx, password)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (service *UserService) GoogleSignIn(ctx context.Context, token string) (*domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GoogleSignIn")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	claims, err := util.ValidateGoogleJWT(token)
	if err != nil {
		return &domain.User{}, err
	}

	exists, err := service.userRepository.DoesUserExists(ctx, claims.Email)
	if err != nil {
		return &domain.User{}, err
	}

	if !exists {
		newUser := &persistence.User{
			FirstName:       claims.FirstName,
			LastName:        claims.LastName,
			Email:           claims.Email,
			Username:        claims.FirstName + "_" + claims.LastName + "_" + uuid.NewV4().String(),
			Password:        "", // Google will handle password
			Role:            model.Basic,
			IsActive:        true,
			ApprovedAccount: true,
		}
		dbUser, err := service.CreateUserWithAdditionalInfo(ctx, newUser, &persistence.UserAdditionalInfo{})
		if err != nil {
			return &domain.User{}, errors.New("unable to create user")
		}
		return dbUser, nil
	} else {
		dbUser, err := service.userRepository.GetUserByEmail(claims.Email)
		if err != nil {
			return &domain.User{}, errors.New("unable to create user")
		}
		return &dbUser, nil
	}
}

func (service *UserService) UpdateUserPhoto(ctx context.Context, userId string, photo string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if userId == "" || photo == "" {
		return errors.New("Invalid arguments")
	}

	err := service.userRepository.UpdateUserPhoto(ctx, userId, photo)
	if err != nil {
		return err
	}
	return nil
}
func (service *UserService) CheckIsApproved(ctx context.Context, id string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPublicUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.userRepository.CheckIsApproved(ctx, id)
}

func (service *UserService) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.userRepository.GetUserByUsername(username)
}

func (service *UserService) GetUserPhoto(ctx context.Context, userId string) (string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if userId == "" {
		return "", errors.New("can't find photo from non-existing user")
	}

	return service.userRepository.GetUserPhoto(ctx, userId)
}

func (service *UserService) CheckIsActive(ctx context.Context, id string) (bool , error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIsActive")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.userRepository.CheckIsActive(ctx, id)
}

func (service UserService) ChangeUserActiveStatus(ctx context.Context, id string) (error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "ChangeUserActiveStatus")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.userRepository.ChangeUserActiveStatus(ctx, id)

}