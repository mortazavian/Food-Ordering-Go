package routes

import (
	"Food-Ordering/internal/handlers"
	"Food-Ordering/internal/middleware"
	"github.com/labstack/echo/v4"
)

func UserRouts(e *echo.Echo) {

	userGroup := e.Group("/api/users")
	userGroup.GET("/create", handlers.CreateUser)
	userGroup.POST("/login", handlers.LoginHandler)
	userGroup.GET("/profile", handlers.ProfileHandler, middleware.JwtMiddleware)
}
