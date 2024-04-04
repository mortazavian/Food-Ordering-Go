package handlers

import (
	"Food-Ordering/internal/models"
	"Food-Ordering/internal/repository"
	"Food-Ordering/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createUserResponse struct {
	Message   string `json:"message"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return err
	}

	// Check if the email pattern is valid
	if !utils.IsValidEmail(user.Email) {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid email address"})
	}

	// Check if the email is already in use
	existingUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Email is already in use"})
	}

	// Create the new user
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

// LoginHandler handles user authentication and returns a JWT token
func LoginHandler(c echo.Context) error {
	var loginReq LoginRequest
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	user, err := repository.AuthenticateUser(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to authenticate user"})
	}

	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	// Store the authenticated user in the context for future use
	c.Set("user", user)

	// Generate a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate JWT token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

// ProfileHandler retrieves the authenticated user from the context
func ProfileHandler(c echo.Context) error {

	user, ok := c.Get("user").(*models.User)
	if !ok || user == nil {
		// Handle the case where the type assertion failed or the user is nil
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User not authenticated"})
	}

	userProfile := models.UserProfile{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email}

	return c.JSON(http.StatusOK, userProfile)
}
