package config

import (
	"log"

	"levelearn-backend/internal/entity"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	log.Println("Running database migration...")

	err := db.AutoMigrate(
		&entity.User{},
		&entity.Course{},
		&entity.Chapter{},
		&entity.Material{},
		&entity.Assessment{},
		&entity.Assignment{},

		&entity.UserCourse{},
		&entity.UserChapter{},

		&entity.Badge{},
		&entity.UserBadge{},

		&entity.Trade{},
		&entity.UserTrade{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully")
}
