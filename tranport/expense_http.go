package tranport

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/vidhi-k/expense_tracker_backend/types"
)

type expenseHTTPHandler struct {
	service types.ExpenseService
}

func InitExpenseHTTPHandlers(expenseService types.ExpenseService, v1 *echo.Group) {
	h := expenseHTTPHandler{service: expenseService}

	userGroup := v1.Group("/expenses")

	userGroup.GET("", h.getExpenses)
	userGroup.POST("", h.createExpense)
}

func (h expenseHTTPHandler) getExpenses(c echo.Context) error {
	req := new(types.GetExpensesRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := h.service.GetExpenses(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h expenseHTTPHandler) createExpense(c echo.Context) error {
	req := new(types.CreateExpenseRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := h.service.CreateExpense(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
