package repository

import (
	"github.com/vidhi-k/expense_tracker_backend/pkg/user"
	"github.com/vidhi-k/expense_tracker_backend/types"
	"gorm.io/gorm"
)

type repo struct{}

func (r repo) GetUserByID(db *gorm.DB, id uint) (*user.User, error) {
	usr := new(user.User)
	err := db.Model(&user.User{}).Where("id = ?", id).First(usr).Error

	return usr, err
}

func (r repo) GetUserByEmail(db *gorm.DB, email string) (*user.User, error) {
	usr := new(user.User)

	err := db.Model(&user.User{}).Where("email = ?", email).First(usr).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, types.ErrUserNotFound.SetError(err)
	}

	return usr, err
}

func (r repo) CreateUser(db *gorm.DB, user *user.User) error {
	err := db.Save(user).Error

	return err
}

func NewPostgresRepo() user.DB {
	return &repo{}
}
