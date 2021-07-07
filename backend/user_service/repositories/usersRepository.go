package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"github.com/david-drvar/xws2021-nistagram/user_service/util"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/encryption"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/images"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	GetAllUsers(context.Context) ([]domain.User, error)
	CreateUser(context.Context, *persistence.User) error
	CreateUserWithAdditionalInfo(context.Context, *persistence.User, *persistence.UserAdditionalInfo) (*domain.User, error)
	UpdateUserProfile(ctx context.Context, dto domain.User) (bool, error)
	UpdateUserPassword(ctx context.Context, password domain.Password) (bool, error)
	SearchUsersByUsernameAndName(ctx context.Context, user *domain.User) ([]domain.User, error)
	LoginUser(context.Context, domain.LoginRequest) (persistence.User, error)
	SaveUserProfilePhoto(ctx context.Context, user *persistence.User) (bool, error)
	GetUserAdditionalInfoById(ctx context.Context, id string) (persistence.UserAdditionalInfo, error)
	GetUserByEmail(email string) (domain.User, error)
	ChangeForgottenPass(ctx context.Context, password domain.Password) (bool, error)
	ApproveAccount(ctx context.Context, password domain.Password) (bool, error)
	GetUserById(context.Context, string) (persistence.User, error)
	DoesUserExists(context.Context, string) (bool, error)
	UpdateUserPhoto(ctx context.Context, userId string, photo string) error

	CheckIsApproved(ctx context.Context, id string) (bool, error)
	GetUserByUsername(username string) (domain.User, error)
	GetUserPhoto(context.Context, string) (string, error)
	CheckIsActive(context.Context, string) (bool, error)
	ChangeUserActiveStatus(context.Context, string) error
}

type userRepository struct {
	DB                *gorm.DB
	privacyRepository PrivacyRepository
	redisServer       saga.RedisServer
}

func NewUserRepo(db *gorm.DB, redisServer *saga.RedisServer) (*userRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}
	privacyRepository, _ := NewPrivacyRepo(db)

	return &userRepository{DB: db, privacyRepository: privacyRepository, redisServer: *redisServer}, nil
}

func (repository *userRepository) UpdateUserPassword(ctx context.Context, password domain.Password) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user *persistence.User

	db := repository.DB.Select("password").Where("id = ?", password.Id).Find(&user)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	err := encryption.CompareHashAndPassword([]byte(user.Password), []byte(password.OldPassword))
	if err != nil {
		return false, err
	}

	db = repository.DB.Model(&user).Where("id = ?", password.Id).Updates(persistence.User{Password: encryption.HashAndSalt([]byte(password.NewPassword))})
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *userRepository) UpdateUserProfile(ctx context.Context, userDTO domain.User) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserProfile")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user persistence.User
	var userAdditionalInfo persistence.UserAdditionalInfo

	db := repository.DB.Model(&user).Where("id = ?", userDTO.Id).Updates(persistence.User{FirstName: userDTO.FirstName, LastName: userDTO.LastName, Email: userDTO.Email, Username: userDTO.Username, BirthDate: userDTO.BirthDate,
		PhoneNumber: userDTO.PhoneNumber, Sex: userDTO.Sex, ResetCode: userDTO.ResetCode, ApprovedAccount: userDTO.ApprovedAccount, TokenEnd: userDTO.TokenEnd})

	fmt.Println(db.RowsAffected)

	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	if userDTO.Role != model.UserRole("Admin") {
		userAdditionalInfoUpdate := persistence.UserAdditionalInfo{Website: userDTO.Website, Category: userDTO.Category, Biography: userDTO.Biography}
		if userAdditionalInfoUpdate.Website == "" {
			userAdditionalInfoUpdate.Website = " "
		}
		if userAdditionalInfoUpdate.Biography == "" {
			userAdditionalInfoUpdate.Biography = " "
		}
		db = repository.DB.Model(&userAdditionalInfo).Where("id = ?", userDTO.Id).Updates(userAdditionalInfoUpdate)
	}

	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *userRepository) GetUserByUsername(username string) (domain.User, error) {
	var dbUser persistence.User
	var dbUserAdditionalInfo persistence.UserAdditionalInfo

	db := repository.DB.Where("is_active = ?", true).Where("username = ?", username).Find(&dbUser)
	if db.Error != nil {
		return domain.User{}, db.Error
	}

	db = repository.DB.Where("id = ?", dbUser.Id).Find(&dbUserAdditionalInfo)
	if db.Error != nil {
		return domain.User{}, db.Error
	}
	user := &domain.User{}

	user = user.GenerateUserDTO(dbUser, dbUserAdditionalInfo)

	if user.ProfilePhoto != "" {
		filename, err := images.LoadImageToBase64(user.ProfilePhoto)
		if err != nil {
			return domain.User{}, err
		}
		user.ProfilePhoto = filename
	}

	return *user, nil
}

func (repository *userRepository) GetUserByEmail(email string) (domain.User, error) {
	var dbUser persistence.User
	var dbUserAdditionalInfo persistence.UserAdditionalInfo

	db := repository.DB.Where("email = ?", email).Where("is_active = ?", true).Find(&dbUser)
	if db.Error != nil {
		return domain.User{}, db.Error
	}

	db = repository.DB.Where("id = ?", dbUser.Id).Find(&dbUserAdditionalInfo)
	if db.Error != nil {
		return domain.User{}, db.Error
	}

	var user *domain.User
	user = user.GenerateUserDTO(dbUser, dbUserAdditionalInfo)

	if user.ProfilePhoto != "" {
		filename, err := images.LoadImageToBase64(user.ProfilePhoto)
		if err != nil {
			return domain.User{}, err
		}
		user.ProfilePhoto = filename
	}

	return *user, nil
}

func (repository *userRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var users []persistence.User
	var dbUserAdditionalInfo persistence.UserAdditionalInfo
	var usersDomain []domain.User

	result := repository.DB.Where("is_active = ?", true).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, user := range users {
		result = repository.DB.Where("id = ?", user.Id).Find(&dbUserAdditionalInfo)
		if result.Error != nil {
			return nil, result.Error
		}

		userDomain := &domain.User{}
		userDomain = userDomain.GenerateUserDTO(user, dbUserAdditionalInfo)
		imageBase64, err := images.LoadImageToBase64(userDomain.ProfilePhoto)
		if err != nil {
			continue
		}
		userDomain.ProfilePhoto = imageBase64

		usersDomain = append(usersDomain, *userDomain)
	}

	return usersDomain, nil
}

func (repository *userRepository) LoginUser(ctx context.Context, request domain.LoginRequest) (persistence.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "LoginUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var dbUser persistence.User
	result := repository.DB.Where("email = ?", request.Email).Where("is_active = ?", true).First(&dbUser)
	if result.Error != nil || result.RowsAffected != 1 {
		return persistence.User{}, result.Error
	}

	if dbUser.Role == "Agent" {
		var request *persistence.RegistrationRequest
		repository.DB.Where("user_id = ?", dbUser.Id).Find(&request)
		if request.Status != "Accepted" {
			return persistence.User{}, errors.New("Request is not Accepted!")
		}
	}

	err := encryption.CompareHashAndPassword([]byte(dbUser.Password), []byte(request.Password))
	if err != nil {
		return persistence.User{}, errors.New("passwords do not match")
	}

	return dbUser, nil
}

func (repository *userRepository) CreateUser(ctx context.Context, user *persistence.User) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	user.Id = uuid.New().String()
	result := repository.DB.Create(&user)

	return result.Error
}

func (repository *userRepository) GetUserById(ctx context.Context, id string) (persistence.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user persistence.User
	result := repository.DB.Where("id = ?", id).Find(&user)
	if result.Error != nil || result.RowsAffected != 1 {
		return persistence.User{}, errors.New("cannot retrieve this user")
	}

	if user.ProfilePhoto != "" {
		filename, err := images.LoadImageToBase64(user.ProfilePhoto)
		if err != nil {
			return persistence.User{}, err
		}
		user.ProfilePhoto = filename
	}

	return user, nil
}

func (repository *userRepository) CheckEmailExists(ctx context.Context, email string) bool {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckEmailExists")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	count := int64(0)
	repository.DB.Model(persistence.User{}).Where("email = ?", email).Count(&count)
	return count > 0

}

func (repository *userRepository) CreateUserWithAdditionalInfo(ctx context.Context, user *persistence.User, userAdditionalInfo *persistence.UserAdditionalInfo) (*domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserWithAdditionalInfo")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var userDomain domain.User
	userDomain, _ = repository.GetUserByUsername(user.Username)
	if userDomain.Username == user.Username {
		return nil, errors.New("username already exists")
	}

	if repository.CheckEmailExists(ctx, user.Email) {
		return nil, errors.New("email already exists")
	}

	user.Id = uuid.New().String()
	resultUser := repository.DB.Create(&user)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}

	if user.ProfilePhoto != "" {
		_, err := repository.SaveUserProfilePhoto(ctx, user)
		if err != nil {
			return nil, err
		}
	}

	userAdditionalInfo.Id = user.Id
	if userAdditionalInfo.Website == "" {
		userAdditionalInfo.Website = " "
	}
	if userAdditionalInfo.Biography == "" {
		userAdditionalInfo.Biography = " "
	}
	err2 := repository.DB.Create(&userAdditionalInfo)
	print(err2)

	var privacy = persistence.Privacy{}
	privacy.UserId = user.Id
	privacy.IsProfilePublic = true
	privacy.IsDMPublic = true
	privacy.IsTagEnabled = true
	_, err := repository.privacyRepository.CreatePrivacy(ctx, &privacy)
	if err != nil {
		return nil, err
	}

	var userReturn *domain.User
	userReturn = userReturn.GenerateUserDTO(*user, *userAdditionalInfo)

	//todo SAGA
	m := saga.Message{Service: saga.ServiceRecommendation, SenderService: saga.ServiceUser, Action: saga.ActionStart, UserId: user.Id}
	repository.redisServer.Orchestrator.Next(saga.RecommendationChannel, saga.ServiceRecommendation, m)

	return userReturn, nil
}

func (repository *userRepository) GetUserAdditionalInfoById(ctx context.Context, id string) (persistence.UserAdditionalInfo, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserAdditionalInfoById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var dbUserAdditionalInfo persistence.UserAdditionalInfo
	db := repository.DB.Where("id = ?", id).Find(&dbUserAdditionalInfo)
	if db.Error != nil {
		return persistence.UserAdditionalInfo{}, db.Error
	}
	return dbUserAdditionalInfo, nil
}

func (repository *userRepository) SearchUsersByUsernameAndName(ctx context.Context, user *domain.User) ([]domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SearchUsersByUsernameAndName")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var users []persistence.User

	if user.Username != "" && user.FirstName != "" && user.LastName != "" {
		repository.DB.Where("is_active = ?", true).Where("username ILIKE ? AND first_name ILIKE ? AND last_name ILIKE ? AND role != 'Admin'", "%"+user.Username+"%", "%"+user.FirstName+"%", "%"+user.LastName+"%").Find(&users)
	} else if user.Username != "" && user.FirstName != "" && user.LastName == "" {
		repository.DB.Where("is_active = ?", true).Where("username ILIKE ? AND first_name ILIKE ? AND role != 'Admin'", "%"+user.Username+"%", "%"+user.FirstName+"%").Find(&users)
	} else if user.Username != "" && user.FirstName == "" && user.LastName != "" {
		repository.DB.Where("is_active = ?", true).Where("username ILIKE ? AND last_name ILIKE ? AND role != 'Admin'", "%"+user.Username+"%", "%"+user.LastName+"%").Find(&users)
	} else if user.Username == "" && user.FirstName != "" && user.LastName != "" {
		repository.DB.Where("is_active = ?", true).Where("first_name ILIKE ? AND last_name ILIKE ? AND role != 'Admin'", "%"+user.FirstName+"%", "%"+user.LastName+"%").Find(&users)
	} else if user.Username != "" && user.FirstName == "" && user.LastName == "" {
		repository.DB.Where("is_active = ?", true).Where("username ILIKE ? AND role != 'Admin'", "%"+user.Username+"%").Find(&users)
	} else if user.Username == "" && user.FirstName != "" && user.LastName == "" {
		repository.DB.Where("is_active = ?", true).Where("first_name ILIKE ? AND role != 'Admin'", "%"+user.FirstName+"%").Find(&users)
	} else if user.Username == "" && user.FirstName == "" && user.LastName != "" {
		repository.DB.Where("is_active = ?", true).Where("last_name ILIKE ? AND role != 'Admin'", "%"+user.LastName+"%").Find(&users)
	}

	var usersDomain []domain.User

	for _, v := range users { //i - index, v - user
		user := &domain.User{}
		dbUserAdditionalInfo, err := repository.GetUserAdditionalInfoById(ctx, v.Id)
		if err != nil {
			return nil, err
		}
		user = user.GenerateUserDTO(v, dbUserAdditionalInfo)

		if v.ProfilePhoto != "" {
			filename, err := images.LoadImageToBase64(v.ProfilePhoto)
			if err != nil {
				return nil, err
			}
			user.ProfilePhoto = filename
		}
		usersDomain = append(usersDomain, *user)
	}

	return usersDomain, nil
}

func (repository *userRepository) SaveUserProfilePhoto(ctx context.Context, user *persistence.User) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveUserProfilePhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	mimeType, err := images.GetImageType(user.ProfilePhoto)
	if err != nil {
		return false, err
	}

	t := time.Now()
	formatted := fmt.Sprintf("%s%d%02d%02d%02d%02d%02d%02d", user.Id, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	name := formatted + "." + mimeType

	err = images.SaveImage(name, user.ProfilePhoto)
	if err != nil {
		return false, err
	}

	user.ProfilePhoto = util.GetContentLocation(name)
	db := repository.DB.Model(&user).Where("id = ?", user.Id).Updates(user)
	if db.Error != nil {
		return false, db.Error
	}

	return true, nil
}

func (repository *userRepository) ChangeForgottenPass(ctx context.Context, password domain.Password) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user *persistence.User

	db := repository.DB.Model(&user).Where("id = ?", password.Id).Updates(persistence.User{Password: encryption.HashAndSalt([]byte(password.NewPassword))})

	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *userRepository) ApproveAccount(ctx context.Context, password domain.Password) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user *persistence.User

	db := repository.DB.Select("password").Where("id = ?", password.Id).Find(&user)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	err := encryption.CompareHashAndPassword([]byte(user.Password), []byte(password.OldPassword))
	if err != nil {
		return false, err
	}

	db = repository.DB.Model(&user).Where("id = ?", password.Id).Updates(persistence.User{Password: encryption.HashAndSalt([]byte(password.NewPassword)), ApprovedAccount: true})
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *userRepository) DoesUserExists(ctx context.Context, email string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DoesUserExists")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user *persistence.User
	db := repository.DB.Model(&persistence.User{}).Where("email = ?", email).Find(&user)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (repository *userRepository) UpdateUserPhoto(ctx context.Context, userId string, photo string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user persistence.User
	user, _ = repository.GetUserById(ctx, userId)
	user.ProfilePhoto = photo
	_, err := repository.SaveUserProfilePhoto(ctx, &user)
	if err != nil {
		return err

	}

	return nil
}

func (repository *userRepository) CheckIsApproved(ctx context.Context, id string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return false, err
	}
	return user.ApprovedAccount, nil
}

func (repository *userRepository) GetUserPhoto(ctx context.Context, userId string) (string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserPhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user *persistence.User
	result := repository.DB.Model(&persistence.User{}).Where("id = ?", userId).Find(&user)
	if result.Error != nil || result.RowsAffected != 1 {
		return "", result.Error
	}

	if user.ProfilePhoto != "" {
		photo, err := images.LoadImageToBase64(user.ProfilePhoto)
		if err != nil {
			return "", err
		}
		return photo, nil
	}

	return "", nil
}

func (repository *userRepository) CheckIsActive(ctx context.Context, id string) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIsActive")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return false, err
	}
	return user.IsActive, nil

}

func (repository userRepository) ChangeUserActiveStatus(ctx context.Context, id string) (error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIsActive")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user persistence.User
	user, _ = repository.GetUserById(ctx, id)
	user.IsActive = !user.IsActive
	_, err := repository.SaveUserProfilePhoto(ctx, &user)
	if err != nil {
		return err

	}

	return nil

}


