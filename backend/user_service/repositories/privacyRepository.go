package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"gorm.io/gorm"
)

type PrivacyRepository interface {
	CreatePrivacy(context.Context, *persistence.Privacy) 	(persistence.Privacy, error)
	UpdatePrivacy(context.Context, *persistence.Privacy) 	(bool, error)
	BlockUser(context.Context, *persistence.BlockedUsers) 	(bool, error)
	UnBlockUser(context.Context, *persistence.BlockedUsers) (bool, error)
	CheckIfBlocked(context.Context, string, string) 		(bool, error)
	GetUserPrivacy(context.Context, string) 				(*persistence.Privacy, error)
	GetAlLPublicUsers(context.Context) 						([]persistence.Privacy, error)
	GetBlockedUsers(context.Context, string)     ([]persistence.BlockedUsers, error)
}

type privacyRepository struct {
	DB *gorm.DB
}

func NewPrivacyRepo(db *gorm.DB) (PrivacyRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}

	return &privacyRepository{DB: db}, nil
}

func (repository *privacyRepository) BlockUser(ctx context.Context, b *persistence.BlockedUsers) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "BlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	db := repository.DB.Create(&b)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *privacyRepository) GetBlockedUsers(ctx context.Context, userId string) ([]persistence.BlockedUsers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetBlockedUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var blockedUsers []persistence.BlockedUsers

	db := repository.DB.Where("user_id = ?", userId).Find(&blockedUsers)
	if db.Error != nil {
		return nil, errors.New("Could not find blocked users")
	}

	return blockedUsers, nil
}

func (repository *privacyRepository) UnBlockUser(ctx context.Context, b *persistence.BlockedUsers) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UnBlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	db := repository.DB.Delete(&b)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *privacyRepository) CheckIfBlocked(ctx context.Context, requestedUserId string, requestingUserId string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIfBlocked")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var blocked persistence.BlockedUsers
	result := repository.DB.
		Where(
			repository.DB.Where("(user_id = ? AND blocked_user_id = ?)", requestedUserId, requestingUserId)).
		Or(
			repository.DB.Where("(user_id = ? AND blocked_user_id = ?)", requestingUserId, requestedUserId)).
		Find(&blocked)

	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected != 0, nil
}

func (repository *privacyRepository) CreatePrivacy(ctx context.Context, p *persistence.Privacy) (persistence.Privacy, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Create(&p)
	return *p, err.Error
}

func (repository *privacyRepository) UpdatePrivacy(ctx context.Context, p *persistence.Privacy) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdatePrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err := repository.DB.Model(&p).Updates(map[string]interface{}{
		"is_tag_enabled":          p.IsTagEnabled,
		"is_profile_public": p.IsProfilePublic,
		"is_dm_public": p.IsDMPublic,
	}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repository *privacyRepository) GetUserPrivacy(ctx context.Context, userId string) (*persistence.Privacy, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserPrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var privacy persistence.Privacy

	db := repository.DB.Where("user_id = ?", userId).Find(&privacy)
	if db.Error != nil {
		return nil, nil
	}

	return &privacy, nil
}

func (repository *privacyRepository) GetAlLPublicUsers(ctx context.Context) ([]persistence.Privacy, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAlLPublicUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var privacies []persistence.Privacy

	db := repository.DB.Where("is_profile_public = true").Find(&privacies)
	if db.Error != nil {
		return nil, nil
	}



	return privacies, nil
}
