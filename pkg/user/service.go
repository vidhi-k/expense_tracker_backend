package user

import "gorm.io/gorm"

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
