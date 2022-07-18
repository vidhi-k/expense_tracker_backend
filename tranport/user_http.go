package tranport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vidhi-k/expense_tracker_backend/types"
)

type userHTTPHandler struct {
	service types.UserService
}

func InitUserHTTPHandlers(userService types.UserService, v1 *echo.Group) {
	h := userHTTPHandler{service: userService}

	userGroup := v1.Group("/users")

	userGroup.GET("/:id", h.getUser)
}

func (h userHTTPHandler) getUser(ctx echo.Context) error {
	res, err := h.service.GetUserByEmail(ctx.Request().Context(), ctx.Param("id"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}
