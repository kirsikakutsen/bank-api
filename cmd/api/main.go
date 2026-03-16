package main

import (
	"log/slog"

	"go-bank/internal/database"
	"go-bank/internal/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
)


func main() {
  err := godotenv.Load()
  if err != nil {
    slog.Error("Error loading .env file")
  }
  database.Connect()
  e := echo.New()

  routes.SetupRouter(e)

  if err := e.Start(":8080"); err != nil {
    slog.Error("failed to start server", "error", err)
  }
}

