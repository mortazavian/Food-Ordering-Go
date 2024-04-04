package main

import (
	"Food-Ordering/internal/repository"
	"Food-Ordering/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	db := repository.NewGormPostgres()

	err := repository.Migration(db)
	if err != nil {
		// Handle migration error
		return
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.UserRouts(e)

	routes.RestaurantRouts(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	err = e.Start(":8080")
	if err != nil {
		// Handle server start error
		return
	}
}
