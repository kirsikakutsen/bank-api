package handlers

import (
	"errors"
	"go-bank/internal/database"
	"go-bank/internal/dto/response"
	"go-bank/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func GetTransfers(c *echo.Context) error {
	idStr := c.QueryParam("id")
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")
	accID := c.Get("accID").(uint)
	sortType := c.QueryParam("type")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 20
	}

	if idStr != "" {
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(errors.New("Invalid id")))
		}

		id := uint(idInt)
		var transaction models.Transaction
		if err := database.DB.First(&transaction, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, response.NewBasicErrorDto(errors.New("Transaction not found")))
		}

		isOwner := false
		if transaction.FromID != nil && *transaction.FromID == accID {
			isOwner = true
		}
		if transaction.ToID != nil && *transaction.ToID == accID {
			isOwner = true
		}

		if !isOwner {
			return c.JSON(http.StatusForbidden, response.NewBasicErrorDto(errors.New("Access denied")))
		}

		resp := response.NewTransactionResponseDto(transaction)
		return c.JSON(http.StatusOK, response.NewBasicSuccessDto(resp))
	}

	offset := (page - 1) * limit

	var transactions []models.Transaction

	query := database.DB.
		Where("from_id = ? OR to_id = ?", accID, accID)

	if sortType != "" {
		switch sortType {
		case "transfer", "deposit", "withdrawal":
			query = query.Where("type = ?", sortType)
		default:
			return c.JSON(http.StatusBadRequest,
				response.NewBasicErrorDto(errors.New("Invalid type parameter")))
		}
	}

	if err := query.
    Order("created_at DESC"). 
    Offset(offset).
    Limit(limit).
    Find(&transactions).Error; err != nil {
    
    return c.JSON(http.StatusInternalServerError,
        response.NewBasicErrorDto(errors.New("Failed to fetch transactions")))
}

	resp := make([]response.TransactionResponseDto, len(transactions))
	for i, t := range transactions {
		resp[i] = response.NewTransactionResponseDto(t)
	}

	return c.JSON(http.StatusOK, response.NewBasicSuccessDto(resp))
}
