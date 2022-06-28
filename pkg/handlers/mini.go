package handlers

import (
	"fmt"
	"net/http"

	"github.com/bveranoc/mu_server/pkg/database"
	"github.com/bveranoc/mu_server/pkg/libs"
	"github.com/bveranoc/mu_server/pkg/models"
	"github.com/labstack/echo/v4"
)

func CreateMini(c echo.Context) error {
	userID, err := libs.GetUserID(c)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err.Error()}
	}

	var mini models.Mini
	if err := c.Bind(&mini); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Something went wrong. Try again!"}

	}

	if mini.Url == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid parameters. Try Again!"}
	}
	mini.UserID = uint(userID.(float64))

	var count int64
	database.DB.Model(&models.Mini{}).Where("ID = ?", userID).Count(&count)
	if count <= 10 {
		database.DB.Create(&mini)
		mini.Mini = libs.EncodeBase62(mini.ID)
		return c.JSON(http.StatusOK, echo.Map{"mini": mini})
	}

	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "You've reach your 10 minis max."}
}

func GetMinis(c echo.Context) error {
	userID, err := libs.GetUserID(c)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err.Error()}
	}

	var minis []models.Mini
	database.DB.Select("id", "created_at", "url").Where("user_id = ?", userID).Order("created_at desc").Find(&minis)

	for i, v := range minis {
		minis[i].Mini = libs.EncodeBase62(v.ID)
	}

	return c.JSON(http.StatusOK, echo.Map{"minis": minis})
}

func GetMini(c echo.Context) error {
	miniString := c.Param("mini")
	id, err := libs.DecodeBase62(miniString)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Something went wrong. Try again!"}
	}

	var mini models.Mini
	res := database.DB.Select("url").Where("ID = ?", id).First(&mini)
	if res.Error != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "Mini not found :("}
	}

	return c.JSON(http.StatusOK, echo.Map{"mini": mini})
}

func DeleteMini(c echo.Context) error {
	id := c.ParamValues()[0]
	userID, err := libs.GetUserID(c)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err.Error()}
	}

	res := database.DB.Where("user_id = ?", uint(userID.(float64))).Where("ID = ?", id).Delete(&models.Mini{})
	if res.RowsAffected == 0 {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "Mini not found :("}
	}

	fmt.Println(res)

	return c.JSON(http.StatusOK, echo.Map{})
}
