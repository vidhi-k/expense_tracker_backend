package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/vidhi-k/expense_tracker_backend/utl/fault"
)

func InitEcho() *echo.Echo {
	ech := echo.New()

	ech.Use(middleware.Logger(), middleware.Recover(), middleware.CORS())

	ech.HTTPErrorHandler = errorHandler

	return ech
}

func errorHandler(err error, c echo.Context) {
	respErr := new(fault.HTTPErrorResp)

	respErr.Message = "something went wrong, try again later."
	respErr.Status = http.StatusInternalServerError
	respErr.Err = err.Error()

	if herr, ok := err.(*echo.HTTPError); ok {
		respErr.Status = herr.Code
		if msg, ok := herr.Message.(string); ok {
			respErr.Message = msg
		}

		if herr.Internal != nil {
			respErr.Message = herr.Internal.Error()
		}
	}

	if herr, ok := err.(*fault.AppError); ok {
		respErr.Status = herr.Status
		respErr.Message = herr.Message
		respErr.Err = ""
		if herr.Err != nil {
			respErr.Err = herr.Err.Error()
		}
	}

	err = c.JSON(respErr.Status, respErr)
	if err != nil {
		log.Println(err)
	}
}

func StartServer(ech *echo.Echo, port int) {
	err := ech.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println(err)
	}
}
