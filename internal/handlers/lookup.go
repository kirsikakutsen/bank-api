package handlers

import (
	"errors"
	"go-bank/internal/database"
	"go-bank/internal/dto/response"
	"go-bank/internal/models"
	"net/http"

	"github.com/labstack/echo/v5"
)

func LookupUser(c *echo.Context) error {
	username := c.QueryParam("username")
	
	if username == "" {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(errors.New("Username is required")))
	}

	var acc models.Account
	if err := database.DB.Where("username = ?", username).First(&acc).Error; err != nil {
		return c.JSON(http.StatusNotFound, response.NewBasicErrorDto(errors.New("User not found")))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"exists":   true,
		"username": username,
	})
}