package handlers

import (
	"errors"
	"go-bank/internal/auth"
	"go-bank/internal/database"
	"go-bank/internal/dto/request"
	"go-bank/internal/dto/response"
	"go-bank/internal/models"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func Login(c *echo.Context) error {
	var req request.AuthRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(err))
	}

	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(err))
	}

	var acc models.Account
	if err := database.DB.Where("username = ?", req.Username).First(&acc).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, response.NewBasicErrorDto(errors.New("No such account")))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, response.NewBasicErrorDto(errors.New("Wrong credentials")))
	}

	token, err := auth.SignPayload(acc.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewBasicErrorDto(errors.New("Failed to generate token")))
	}

	resp := response.AuthDto{
		Token: token,
		Profile: response.NewAccountDto(acc),
	}

	return c.JSON(http.StatusOK, response.NewBasicSuccessDto(resp))
}

func Signup(c *echo.Context) error {
	var req request.AuthRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(err))
	}

	if len(req.Password) < 8 {
		return c.JSON(http.StatusBadRequest,
			response.NewBasicErrorDto(errors.New("Password must be at least 8 characters")))
	}

	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBasicErrorDto(err))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewBasicErrorDto(errors.New("Hashing password error, try again")))
	}

	acc := models.Account{
		Username: req.Username,
		Password: string(hashedPassword),
		BalanceCents: 0,
	}

	if err := database.DB.Create(&acc).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewBasicErrorDto(errors.New("Failed to create account")))
	}

	token, err := auth.SignPayload(acc.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewBasicErrorDto(errors.New("Failed to generate token")))
	}

	resp := response.AuthDto{
		Token: token,
		Profile: response.NewAccountDto(acc),
	}

	return c.JSON(http.StatusCreated, response.NewBasicSuccessDto(resp))
}

func Logout(c *echo.Context) error {
	return c.JSON(http.StatusOK, response.NewBasicSuccessDto("Logged out successfully"))
}
