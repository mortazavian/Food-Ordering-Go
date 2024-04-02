package main

import (
	"Food-Ordering/internal/repository"
	"Food-Ordering/internal/routes"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func main() {

	db := repository.NewGormPostgres()

	err := repository.Migration(db)
	if err != nil {
		return
	}

	e := echo.New()

	routes.UserRouts(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	err = e.Start(":8080")
	if err != nil {
		return
	}

}
