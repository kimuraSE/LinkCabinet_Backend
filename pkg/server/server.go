package server

import (
	"LinkCabinet_Backend/internal/api/handler"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)


func NewServer(uh handler.IUserHandler,lh handler.ILinksHandler) *echo.Echo {

	e := echo.New()

	e.POST("/login", uh.Login)
	e.POST("/signup", uh.SignUp)
	e.POST("/logout", uh.Logout)

	linkApi := e.Group("/links")
	linkApi.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey : []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:jwt_token",
	}))

	linkApi.GET("", lh.AllGetLinks)
	linkApi.GET("/:linkId", lh.GetLinkById)
	linkApi.POST("", lh.CreateLink)
	linkApi.PUT("/:linkId", lh.UpdateLink)
	linkApi.DELETE("/:linkId", lh.DeleteLink)
	
	return e

}