package types

import (
	"context"

	"github.com/vidhi-k/expense_tracker_backend/utl/fault"
)

type UserService interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
}

// errors.
var (
	ErrUserNotFound = &fault.AppError{Message: "user not found"}
)

type (
	CreateUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateUserResponse struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	User struct {
		ID       uint
		Name     string
		Email    string
		Password string
	}
)
