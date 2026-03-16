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

func CreateTransfer(c *echo.Context) error {

	var req request.TransferRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(errors.New("Invalid transfer details")))
	}

	if err := validate.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(errors.New("Invalid transfer details")))
	}

	fromID := c.Get("accID").(uint)

	var fromAcc models.Account
	if err := database.DB.First(&fromAcc, fromID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(errors.New("Unable to process transfer")))
	}

	if fromAcc.BalanceCents < req.AmountCents {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(errors.New("Insufficient funds")))
	}

	if fromAcc.Username == req.ToUsername {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(errors.New("Transfers to the same account are not allowed.")))
	}

	var toAcc models.Account
	if err := database.DB.Where("username = ?", req.ToUsername).First(&toAcc).Error; err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(errors.New("Unable to process transfer")))
	}

	transfer := models.Transaction{
		FromID:      &fromID,
		ToID:        &toAcc.ID,
		ToUsername:  toAcc.Username,
		Type: 		 models.Transfer,
		AmountCents: req.AmountCents,
	}

	res := database.DB.Create(&transfer)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.NewBasicErrorDto(errors.New("Failed to create transfer")))
	}

	toAcc.BalanceCents += req.AmountCents
	fromAcc.BalanceCents -= req.AmountCents

	if err := database.DB.Save(&toAcc).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewBasicErrorDto(errors.New("Failed to update receiver balance")))
	}
	if err := database.DB.Save(&fromAcc).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewBasicErrorDto(errors.New("Failed to update sender balance")))
	}

	resp := response.NewTransactionResponseDto(transfer)

	return c.JSON(http.StatusOK, response.NewBasicSuccessDto(resp))
}
