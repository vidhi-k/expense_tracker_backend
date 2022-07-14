package user

import (
	"context"
	"github.com/vidhi-k/expense_tracker_backend/types"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func (s service) CreateUser(ctx context.Context, req *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	return nil, nil
}

// InitService initializes user service.
func InitService(db *gorm.DB) types.UserService {
	return &service{db: db}
}
