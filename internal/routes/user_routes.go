package routes

import (
	"Food-Ordering/internal/handlers"
	"github.com/labstack/echo/v4"
)

func UserRouts(e *echo.Echo) {

	userGroup := e.Group("/api/users")
	userGroup.GET("/create", handlers.CreateUser)
}
