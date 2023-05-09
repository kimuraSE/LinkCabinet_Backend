package server

import (
	"LinkCabinet_Backend/internal/api/handler"

	"github.com/labstack/echo/v4"
)


func NewServer(uh handler.IUserHandler) *echo.Echo {

	e := echo.New()

	e.POST("/login", uh.Login)
	e.POST("/signup", uh.SignUp)
	e.POST("/logout", uh.Logout)

	return e

}