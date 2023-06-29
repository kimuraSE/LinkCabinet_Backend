package handler

import (
	"LinkCabinet_Backend/internal/api/model"
	"LinkCabinet_Backend/internal/api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ILinksHandler interface {
	AllGetLinks(c echo.Context) error
	GetLinkById(c echo.Context) error
	CreateLink(c echo.Context) error
	UpdateLink(c echo.Context) error
	DeleteLink(c echo.Context) error
}

type linksHandler struct {
	lu usecase.ILinksUsecase
}

func NewLinksHandler(lu usecase.ILinksUsecase) ILinksHandler {
	return &linksHandler{lu}
}

func (lh *linksHandler) AllGetLinks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	linksRes, err := lh.lu.AllGetLinks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, linksRes)
}

func (lh *linksHandler) GetLinkById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("linkId")

	linkId, _ := strconv.Atoi(id)

	linkRes, err := lh.lu.GetLinksByUserID(uint(userId.(float64)), uint(linkId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, linkRes)
}

func (lh *linksHandler) CreateLink(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	link := model.Link{}
	if err := c.Bind(&link); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	link.UserID = uint(userId.(float64))
	linkRes, err := lh.lu.CreateLink(link)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, linkRes)
}

func (lh *linksHandler) UpdateLink(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	link := model.Link{}
	if err := c.Bind(&link); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := c.Param("linkId")

	linkId, _ := strconv.Atoi(id)

	linkRes, err := lh.lu.UpdateLink(link, uint(userId.(float64)), uint(linkId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, linkRes)
}

func (lh *linksHandler) DeleteLink(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("linkId")

	linkId, _ := strconv.Atoi(id)

	err := lh.lu.DeleteLink(uint(userId.(float64)), uint(linkId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Link deleted")
}
