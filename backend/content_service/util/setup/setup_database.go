package setup

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	// dropTables(db)

	db.Migrator().DropTable(&persistence.AdCategory{})
	err := db.AutoMigrate(&persistence.Post{},
		&persistence.Story{},
		&persistence.Media{},
		&persistence.Tag{},
		&persistence.Collection{},
		&persistence.Favorites{},
		&persistence.Like{},
		&persistence.Comment{},
		&persistence.Highlight{},
		&persistence.HighlightStory{},
		&persistence.RegistrationRequest{},
		&persistence.Ad{},
		&persistence.Campaign{},
		&persistence.CampaignInfluencerRequest{},
		&persistence.ContentComplaint{},
		&persistence.AdCategory{},
		&persistence.UserAdCategories{},
		&persistence.CampaignChanges{},
		&persistence.Hashtag{},
		&persistence.HashtagObjava{},
	)
	result := db.Create([]*persistence.AdCategory{
		{ Id: "6696fd8b-cbe5-4fe6-b1bc-ad523b9d4346", Name: "Sports"},
		{ Id: "e5cf9990-2a13-4fe5-9bb5-0a6a50c568c1", Name: "Fashion"},
		{ Id: "e48f8c1c-c788-400e-a07a-8c9b3a757f13", Name: "Drinks"},
		{ Id: "cd804355-f4c4-475f-9869-844d5dfe882d", Name: "Technology"},
		{ Id: "e6d6b6be-7910-417c-9ca7-e9dffecf310f", Name: "Events"},
	})
	if result.Error != nil { return result.Error }

	return err
}

func dropTables(db *gorm.DB){
	if db.Migrator().HasTable(&persistence.Post{}) {
		db.Migrator().DropTable(&persistence.Post{},
			&persistence.Story{},
			&persistence.Media{},
			&persistence.Tag{},
			&persistence.Collection{},
			&persistence.Favorites{},
			&persistence.Like{},
			&persistence.Comment{},
			&persistence.Highlight{},
			&persistence.HighlightStory{},
			&persistence.RegistrationRequest{},
			&persistence.Ad{},
			&persistence.Campaign{},
			&persistence.CampaignInfluencerRequest{},
			&persistence.ContentComplaint{},
			&persistence.AdCategory{},
			&persistence.UserAdCategories{},
			&persistence.CampaignChanges{},
			&persistence.Hashtag{},
			&persistence.HashtagObjava{},
		)
	}
}
