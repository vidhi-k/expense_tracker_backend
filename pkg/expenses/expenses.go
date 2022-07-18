package expenses

import (
	"context"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	
	"github.com/vidhi-k/expense_tracker_backend/types"
)

type service struct {
	db          *gorm.DB
	repo        DB
	userService types.UserService
}

// GetExpenses returns expenses of a user.
func (s service) GetExpenses(ctx context.Context, request *types.GetExpensesRequest) (*types.GetExpensesResponse, error) { //nolint:lll
	_, err := s.userService.GetUserByID(ctx, request.UserID)
	if err != nil {
		return nil, err
	}

	expenses, err := s.repo.GetExpenses(s.db, request)
	if err != nil {
		return nil, err
	}

	exp := make([]*types.Expense, 0)

	err = copier.Copy(&exp, &expenses)
	if err != nil {
		return nil, err
	}

	return &types.GetExpensesResponse{Expenses: exp}, nil
}

func (s service) CreateExpense(ctx context.Context, request *types.CreateExpenseRequest) (*types.Expense, error) {
	_, err := s.userService.GetUserByID(ctx, request.UserID)
	if err != nil {
		return nil, err
	}

	expense := new(Expense)

	err = copier.Copy(expense, request)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateExpense(s.db, expense)
	if err != nil {
		return nil, err
	}

	resp := new(types.Expense)

	err = copier.Copy(resp, expense)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// InitService initializes expense service.
func InitService(db *gorm.DB, repo DB, userService types.UserService) types.ExpenseService {
	return &service{db: db, repo: repo, userService: userService}
}
