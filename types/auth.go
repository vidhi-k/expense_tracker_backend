package types

import (
	"context"

	"github.com/vidhi-k/expense_tracker_backend/utl/fault"
)

type AuthService interface {
	Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error)
	Signup(ctx context.Context, request *SignupRequest) (*SignupResponse, error)
}

// auth service errors.
var (
	ErrInvalidCredentials = &fault.AppError{Message: "invalid credentials", Status: 401}
	ErrUserAlreadyExists  = &fault.AppError{Message: "user with email already exists", Status: 400}
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
