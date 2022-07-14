package migrator

import (
	"github.com/vidhi-k/expense_tracker_backend/pkg/user"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.Exec("CREATE SCHEMA IF NOT EXISTS usr;").Error
	CheckError(err)

	err = db.AutoMigrate(&user.User{})
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatalf("error running migration: %v", err)
	}
}
