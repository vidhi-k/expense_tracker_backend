package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func InitEcho() *echo.Echo {
	ech := echo.New()

	ech.Use(middleware.Logger(), middleware.Recover())

	return ech
}

func StartServer(ech *echo.Echo, port int) {
	err := ech.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println(err)
	}
}
