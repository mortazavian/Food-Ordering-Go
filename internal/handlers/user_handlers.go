package handlers

import (
	"Food-Ordering/internal/models"
	"Food-Ordering/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateUser(c echo.Context) error {

	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return err
	}

	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
