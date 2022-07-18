package types

import (
	"context"
	"time"
)

type ExpenseService interface {
	GetExpenses(ctx context.Context, request *GetExpensesRequest) (*GetExpensesResponse, error)
	CreateExpense(ctx context.Context, request *CreateExpenseRequest) (*Expense, error)
}

type GetExpensesRequest struct {
	UserID    uint      `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Category  Category  `json:"category"`
}

type GetExpensesResponse struct {
	Expenses []*Expense `json:"expenses"`
}

type CreateExpenseRequest struct {
	UserID   uint     `json:"user_id"`
	Amount   float64  `json:"amount"`
	Category Category `json:"category"`
	Note     *string  `json:"note"`
}

type Expense struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Category  Category  `json:"category"`
	Amount    float64   `json:"amount"`
	Note      *string   `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

type Category string

const CategoryFood Category = "food"
const CategoryTravel Category = "travel"
const CategoryEntertainment Category = "entertainment"
const CategoryShopping Category = "shopping"
const CategoryOther Category = "other"
