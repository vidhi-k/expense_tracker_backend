package repository

import (
	"gorm.io/gorm"

	"github.com/vidhi-k/expense_tracker_backend/pkg/expenses"
	"github.com/vidhi-k/expense_tracker_backend/types"
)

type repo struct{}

func (r repo) CreateExpense(db *gorm.DB, expense *expenses.Expense) error {
	err := db.Save(expense).Error

	return err
}

func (r repo) GetExpenses(db *gorm.DB, request *types.GetExpensesRequest) ([]*expenses.Expense, error) {
	resp := make([]*expenses.Expense, 0)

	db = db.Model(&expenses.Expense{}).Where("user_id = ?", request.UserID)

	if request.Category != "" {
		db = db.Where("category = ?", request.Category)
	}

	if !request.StartDate.IsZero() {
		db = db.Where("created_at >= ?", request.StartDate)
	}

	if !request.EndDate.IsZero() {
		db = db.Where("created_at <= ?", request.EndDate)
	}

	err := db.Find(&resp).Error

	return resp, err
}

func NewPostgresRepo() expenses.DB {
	return &repo{}
}
