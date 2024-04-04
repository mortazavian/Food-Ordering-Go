package routes

import (
	"Food-Ordering/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RestaurantRouts(e *echo.Echo) {
	restaurantGroup := e.Group("api/restaurant")
	restaurantGroup.POST("/create", handlers.CreateRestaurantHandler)
}
