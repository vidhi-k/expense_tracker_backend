package auth

import (
	"context"

	"github.com/vidhi-k/expense_tracker_backend/types"
)

type service struct {
	userService types.UserService
}

func (s service) Login(ctx context.Context, request *types.LoginRequest) (*types.LoginResponse, error) {
	user, err := s.userService.GetUserByEmail(ctx, request.Email)
	if err != nil {
		if err == types.ErrUserNotFound {
			return nil, types.ErrInvalidCredentials.SetError(types.ErrUserNotFound)
		}
	}

	if user.Password != request.Password {
		return nil, types.ErrInvalidCredentials
	}

	return &types.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s service) Signup(ctx context.Context, request *types.SignupRequest) (*types.SignupResponse, error) {
	usr, err := s.userService.GetUserByEmail(ctx, request.Email)
	if err != nil && err != types.ErrUserNotFound {
		return nil, err
	}

	// if user is present return error that user with email already exists
	if usr != nil {
		return nil, types.ErrUserAlreadyExists
	}

	usr, err = s.userService.CreateUser(ctx, &types.CreateUserRequest{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})

	return &types.SignupResponse{
		ID:    usr.ID,
		Name:  usr.Name,
		Email: usr.Email,
	}, nil
}

func InitService(userService types.UserService) types.AuthService {
	return &service{userService: userService}
}
