package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"

	"go-bank/internal/database"
	"go-bank/internal/dto/response"
	"go-bank/internal/models"
)

func GetAccount(c *echo.Context) error {
	accID := c.Get("accID").(uint)

	var acc models.Account
	res := database.DB.Where("id=?", accID).First(&acc)

	if res.Error != nil {
		fmt.Println(res.Error)
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(errors.New("Account not found")))
	}

	resp := response.NewAccountDto(acc)
	return c.JSON(http.StatusOK, response.NewBasicSuccessDto(resp))
}