package types

import "context"

type UserService interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
}

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
)
