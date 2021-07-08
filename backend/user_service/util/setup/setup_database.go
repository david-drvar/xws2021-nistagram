package setup

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	// dropTables(db)

	err := db.AutoMigrate(&persistence.User{},
		&persistence.UserAdditionalInfo{},
		&persistence.Privacy{},
		&persistence.BlockedUsers{},
		&persistence.Followers{},
		&persistence.VerificationRequest{},
		&persistence.RegistrationRequest{},
		&persistence.APIKeys{},
		&persistence.UserNotification{},
		&persistence.CampaignRequest{},
		)

	return err
}

func dropTables(db *gorm.DB) {
	if db.Migrator().HasTable(&persistence.User{}) {
		db.Migrator().DropTable(&persistence.User{},
			&persistence.UserAdditionalInfo{},
			&persistence.Privacy{},
			&persistence.BlockedUsers{},
			&persistence.Followers{},
			&persistence.VerificationRequest{},
			&persistence.APIKeys{},
			&persistence.RegistrationRequest{},
			&persistence.UserNotification{},
			&persistence.CampaignRequest{})
	}
}