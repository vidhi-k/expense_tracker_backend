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

func (s service) GetUserByID(ctx context.Context, id uint) (*types.User, error) {
	user, err := s.repo.GetUserByID(s.db, id)
	if err != nil {
		return nil, err
	}

	usr := new(types.User)

	err = copier.Copy(usr, user)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (s service) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	user, err := s.repo.GetUserByEmail(s.db, email)
	if err != nil {
		return nil, err
	}

	usr := new(types.User)

	err = copier.Copy(usr, user)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (s service) CreateUser(ctx context.Context, req *types.CreateUserRequest) (*types.User, error) {
	user := new(User)

	err := copier.Copy(user, req)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(s.db, user)
	if err != nil {
		return nil, err
	}

	return &types.User{
		ID:    user.ID,
		Name:  *user.Name,
		Email: *user.Email,
	}, nil
}

// InitService initializes user service.
func InitService(db *gorm.DB, repo DB) types.UserService {
	return &service{db: db, repo: repo}
}
