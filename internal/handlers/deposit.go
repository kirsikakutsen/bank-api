package handlers

import (
	"errors"
	"go-bank/internal/database"
	"go-bank/internal/dto/request"
	"go-bank/internal/dto/response"
	"go-bank/internal/models"
	"net/http"

	"github.com/labstack/echo/v5"
)

func CreateDeposit(c *echo.Context) error {
	accID := c.Get("accID").(uint)

	var req request.AmountRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(err))
	}

	if req.AmountCents <= 0 {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(errors.New("Invalid amount")))
	}

	var acc models.Account
	if err := database.DB.First(&acc, accID).Error; err != nil {
		return c.JSON(http.StatusNotFound, response.NewBasicErrorDto(errors.New("Account not found")))
	}

	transaction := models.Transaction{
		ToID:        &accID,
		Type:		 models.Deposit,
		AmountCents: req.AmountCents,
	}

    res := database.DB.Create(&transaction)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.NewBasicErrorDto(errors.New("Failed to create transfer")))
	}

    acc.BalanceCents += req.AmountCents

    if err := database.DB.Save(&acc).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, response.NewBasicErrorDto(errors.New("Failed to save account")))
    }

	resp := response.NewTransactionResponseDto(transaction)
	return c.JSON(http.StatusOK, response.NewBasicSuccessDto(resp))
}
