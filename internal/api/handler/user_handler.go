package handler

import (
	"LinkCabinet_Backend/internal/api/model"
	"LinkCabinet_Backend/internal/api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)


type IUserHandler interface {
	Login(c echo.Context) error
	SignUp(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userHandler struct {
	uu usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) IUserHandler {
	return &userHandler{uu}
}

func (uh *userHandler) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := uh.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteNoneMode
	// cookie.Secure = true
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}


func (uh *userHandler) SignUp(c echo.Context) error {

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := uh.uu.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteNoneMode
	// cookie.Secure = true
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uh *userHandler) Logout(c echo.Context) error {
	cookie :=new(http.Cookie)
	cookie.Name="token"
	cookie.Value=""
	cookie.Expires=time.Now()
	cookie.HttpOnly=true
	cookie.Domain=os.Getenv("API_DOMAIN")
	cookie.SameSite=http.SameSiteNoneMode
	// cookie.Secure=true
	cookie.Path="/"
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uh *userHandler) CsrfToken(c echo.Context) error {
	token:=c.Get("csrf").(string)
	return c.JSON(http.StatusOK,
	echo.Map{
		"csrf_token":token,
	})
}

