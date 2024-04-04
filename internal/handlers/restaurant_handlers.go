package handlers

import (
	"Food-Ordering/internal/models"
	"Food-Ordering/internal/repository"
	"Food-Ordering/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createRestaurantResponse struct {
	message string
	name    string
	email   string
}

func CreateRestaurantHandler(c echo.Context) error {
	restaurant := new(models.Restaurant)
	err := c.Bind(restaurant)
	if err != nil {
		return err
	}

	isValidEmail := utils.IsValidEmail(restaurant.Email)
	if !isValidEmail {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid email address"})
	}

	existingRestaurant := repository.GetRestaurantByEmail(restaurant.Email)

	//if err != nil {
	//	return err
	//}

	if existingRestaurant != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Email is already in use"})
	}

	err = repository.CreateRestaurant(restaurant)
	if err != nil {
		return err
	}

	response := createRestaurantResponse{
		message: "User Created successfully",
		name:    restaurant.Name,
		email:   restaurant.Email,
	}
	return c.JSON(http.StatusOK, response)
}
