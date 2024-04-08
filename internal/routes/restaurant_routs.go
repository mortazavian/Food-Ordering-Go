package routes

import (
	"Food-Ordering/internal/handlers"
	"Food-Ordering/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RestaurantRouts(e *echo.Echo) {
	restaurantGroup := e.Group("api/restaurant")
	restaurantGroup.POST("/create", handlers.CreateRestaurantHandler)
	restaurantGroup.POST("/login", handlers.RestaurantLoginHandler)
	restaurantGroup.GET("/profile", handlers.RestaurantProfileHandler, middleware.RestaurantJwtMiddleware)
	restaurantGroup.POST("/add-menu-item", handlers.AddMenuItemHandler, middleware.RestaurantJwtMiddleware)
	restaurantGroup.GET("/all-menu-items", handlers.GetAllMenuItemHandler, middleware.RestaurantJwtMiddleware)
	restaurantGroup.POST("/create-working-day", handlers.CreateRestaurantWorkingDayHandler, middleware.RestaurantJwtMiddleware)
	restaurantGroup.POST("/create-working-time/:id", handlers.CreateRestaurantWorkingTimeHandler, middleware.RestaurantJwtMiddleware)
}
