package routes

import (
	"go-bank/internal/auth"
	"go-bank/internal/handlers"

	"github.com/labstack/echo/v5"
)

func SetupRouter(e *echo.Echo) {

	api := e.Group("/api")

	api.GET("/ping", handlers.HealthCheck)

	api.POST("/auth/signup", handlers.Signup)
	api.POST("/auth/login", handlers.Login)

	protected := api.Group("")
	protected.Use(auth.VerifyAuth)
	protected.GET("/account", handlers.GetAccount)
	protected.POST("/auth/logout", handlers.Logout)


	protected.POST("/transfer", handlers.CreateTransfer)
	protected.GET("/transfers", handlers.GetTransfers)
	protected.GET("/accounts/lookup", handlers.LookupUser)
	protected.POST("/withdraw", handlers.CreateWithdraw)
	protected.POST("/deposit", handlers.CreateDeposit)
}
