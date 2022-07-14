package tranport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vidhi-k/expense_tracker_backend/types"
)

type httpHandler struct {
	service types.UserService
}

func InitHTTPUserHandlers(userService types.UserService, v1 *echo.Group) {
	h := httpHandler{service: userService}

	userGroup := v1.Group("/users")

	userGroup.POST("", h.createUser)
}

func (h httpHandler) createUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "created user")
}
