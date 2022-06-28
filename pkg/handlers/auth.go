package handlers

import (
	"net/http"

	"github.com/bveranoc/mu_server/pkg/database"
	"github.com/bveranoc/mu_server/pkg/libs"
	"github.com/bveranoc/mu_server/pkg/models"
	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) (err error) {
	// Binding
	var user models.User
	if err = c.Bind(&user); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Something went wrong. Try again!"}
	}

	// Validate
	if user.Nickname == "" || user.Email == "" || user.Provider == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid parameters. Try Again!"}
	}
	database.DB.Where(models.User{Email: user.Email}).FirstOrCreate(&user)

	// Generate Token
	token := libs.GenerateToken(user.ID)
	if token == "" {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Something went wrong. Try again!"}
	}

	return c.JSON(http.StatusCreated, echo.Map{"token": token})
}
