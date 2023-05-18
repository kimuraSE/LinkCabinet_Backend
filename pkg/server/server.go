package server

import (
	"LinkCabinet_Backend/internal/api/handler"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func NewServer(uh handler.IUserHandler,lh handler.ILinksHandler) *echo.Echo {

	e := echo.New()

	
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:[]string{"http://localhost:3000",os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin,echo.HeaderContentType,echo.HeaderAccept,
			echo.HeaderAccessControlAllowCredentials,echo.HeaderXCSRFToken},
			AllowMethods: []string{"GET","POST","PUT","DELETE"},
			AllowCredentials:true,
		}))
		
		//修正
		e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{	
			CookiePath: "/",
			CookieDomain: os.Getenv("API_DOMAIN"),
			CookieHTTPOnly: true,
			CookieSameSite: http.SameSiteNoneMode,
			// CookieSameSite: http.SameSiteDefaultMode,
			// CookieMaxAge: 60,
		}))
		
		e.POST("/login", uh.Login)
		e.POST("/signup", uh.SignUp)
		e.POST("/logout", uh.Logout)
		e.GET("/csrf", uh.CsrfToken)
	
	top := e.Group("/top")
	top.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey : []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:jwt_token",
	}))

	top.DELETE("", uh.DeleteUser)
	top.PUT("/settings/name", uh.UpdateUserName)
	top.PUT("/settings/email", uh.UpdateUserEmail)
	top.PUT("/settings/password", uh.UpdateUserPassword)

	top.GET("", lh.AllGetLinks)
	top.GET("/:linkId", lh.GetLinkById)
	top.POST("", lh.CreateLink)
	top.PUT("/:linkId", lh.UpdateLink)
	top.DELETE("/:linkId", lh.DeleteLink)
	
	return e

}