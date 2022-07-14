package user

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/vidhi-k/expense_tracker_backend/types"
	"gorm.io/gorm"
)

type service struct {
	db   *gorm.DB
	repo DB
}

func (s service) CreateUser(ctx context.Context, req *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	user := new(User)

	err := copier.Copy(user, req)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(s.db, user)
	if err != nil {
		return nil, err
	}

	res := new(types.CreateUserResponse)
	err = copier.Copy(res, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// InitService initializes user service.
func InitService(db *gorm.DB, repo DB) types.UserService {
	return &service{db: db, repo: repo}
}
