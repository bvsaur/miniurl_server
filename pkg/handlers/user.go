package handlers

import (
	"net/http"

	"github.com/bveranoc/mu_server/pkg/database"
	"github.com/bveranoc/mu_server/pkg/libs"
	"github.com/bveranoc/mu_server/pkg/models"
	"github.com/labstack/echo/v4"
)

func GetMe(c echo.Context) error {
	userID, err := libs.GetUserID(c)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err.Error()}
	}
	var user models.User
	database.DB.Select("nickname").First(&user, userID)

	return c.JSON(http.StatusOK, echo.Map{"nickname": user.Nickname})
}

func UpdateNickname(c echo.Context) error {
	userID, err := libs.GetUserID(c)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err.Error()}
	}

	var user models.User
	if err = c.Bind(&user); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Something went wrong. Try again!"}
	}

	if user.Nickname == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid parameters. Try Again!"}
	}

	database.DB.Model(&models.User{}).Where("ID = ?", userID).Update("nickname", user.Nickname)

	return c.JSON(http.StatusOK, echo.Map{"nickname": user.Nickname})
}
