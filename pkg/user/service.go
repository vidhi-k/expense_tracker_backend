package user

import "gorm.io/gorm"

type DB interface {
	CreateUser(db *gorm.DB, user *User) error
}

// User is user entity
type User struct {
	gorm.Model
	Name     *string
	Email    *string
	Password *string
}

func (*User) TableName() string {
	return "usr.users"
}
