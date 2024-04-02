package main

import (
	"Food-Ordering/internal/repository"
	"Food-Ordering/internal/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	db := repository.NewGormPostgres()

	err := repository.Migration(db)
	if err != nil {
		return
	}

	e := echo.New()

	routes.UserRouts(e)

	err = e.Start(":8080")
	if err != nil {
		return
	}
}
