package expenses

import (
	"github.com/vidhi-k/expense_tracker_backend/types"
	"gorm.io/gorm"
)

type DB interface {
	GetExpenses(db *gorm.DB, request *types.GetExpensesRequest) ([]*Expense, error)
	CreateExpense(db *gorm.DB, expense *Expense) error
}

type Expense struct {
	gorm.Model
	UserID   uint
	Category types.Category `gorm:"default:other"`
	Note     *string
	Amount   float64
}

func (*Expense) TableName() string {
	return "expense.expenses"
}
