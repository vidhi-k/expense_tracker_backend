package repository

import (
	"github.com/vidhi-k/expense_tracker_backend/pkg/user"
	"gorm.io/gorm"
)

type repo struct{}

func (r repo) CreateUser(db *gorm.DB, user *user.User) error {
	err := db.Save(user).Error
	return err
}

func NewPostgresRepo() user.DB {
	return &repo{}
}
