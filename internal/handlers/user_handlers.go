package handlers

import (
	"Food-Ordering/internal/models"
	"Food-Ordering/internal/repository"
	"Food-Ordering/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createUserResponse struct {
	Message   string `json:"message"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func CreateUser(c echo.Context) error {

	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return err
	}

	// Check if the email pattern is OK
	isValidEmail := utils.IsValidEmail(user.Email)
	if isValidEmail == false {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid email address"})
	}

	existingUser, err := repository.GetUserByEmail(user.Email)

	if err != nil {

	}
	if existingUser != nil {
		// Email is already in use, return error response
		return c.JSON(http.StatusConflict, map[string]string{"error": "Email is already in use"})
	}

	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	response := createUserResponse{
		Message:   "User Created successfully",
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return c.JSON(http.StatusOK, response)
}
