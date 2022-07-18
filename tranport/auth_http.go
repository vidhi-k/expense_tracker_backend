package tranport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vidhi-k/expense_tracker_backend/types"
)

type authHTTPHandler struct {
	service types.AuthService
}

func InitAuthHTTPHandlers(authService types.AuthService, v1 *echo.Group) {
	h := authHTTPHandler{service: authService}

	userGroup := v1.Group("/auth")

	userGroup.POST("/login", h.login)
	userGroup.POST("/signup", h.signup)
}

func (h authHTTPHandler) login(ctx echo.Context) error {
	req := new(types.LoginRequest)
	if err := ctx.Bind(req); err != nil {
		return err
	}

	res, err := h.service.Login(ctx.Request().Context(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h authHTTPHandler) signup(ctx echo.Context) error {
	req := new(types.SignupRequest)
	if err := ctx.Bind(req); err != nil {
		return err
	}

	res, err := h.service.Signup(ctx.Request().Context(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}
