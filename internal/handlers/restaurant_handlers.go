package handlers

import (
	"Food-Ordering/internal/models"
	"Food-Ordering/internal/repository"
	"Food-Ordering/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createRestaurantResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Email   string `json:"email"`
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
		Message: "User Created successfully",
		Name:    restaurant.Name,
		Email:   restaurant.Email,
	}
	return c.JSON(http.StatusOK, response)
}

// RestaurantLoginHandler handles restaurant authentication and returns a JWT token
func RestaurantLoginHandler(c echo.Context) error {
	var loginReq LoginRequest
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	restaurant, err := repository.AuthenticateRestaurant(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to authenticate restaurant"})
	}

	if restaurant == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	// Store the authenticated user in the context for future use
	c.Set("restaurant", restaurant)

	// Generate a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["restaurant_id"] = restaurant.ID
	claims["email"] = restaurant.Email

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate JWT token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func RestaurantProfileHandler(c echo.Context) error {

	restaurant, ok := c.Get("restaurant").(*models.Restaurant)
	if !ok || restaurant == nil {
		// Handle the case where the type assertion failed or the user is nil
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Restaurant not authenticated"})
	}

	userProfile := models.RestaurantProfile{
		Name:  restaurant.Name,
		Email: restaurant.Email,
	}

	return c.JSON(http.StatusOK, userProfile)
}
