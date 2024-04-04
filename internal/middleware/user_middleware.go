package middleware

import (
	"Food-Ordering/internal/models"
	"Food-Ordering/internal/repository"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var jwtSigningKey = []byte("your-secret-key")

func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*models.User)
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized access"})
		}
		return next(c)
	}
}

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing token"})
		}

		// Parse JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSigningKey, nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
		}

		// Check if token is valid
		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
		}

		// Extract claims from token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token claims"})
		}

		// Extract user ID from claims
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid user ID"})
		}

		// Retrieve user from repository using the user ID
		user, err := repository.GetUserByID(int(userID))
		if err != nil || user == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "user not found"})
		}

		// Set user in context for future use
		c.Set("user", user)

		// If the token is valid and the user is authenticated, proceed to the next middleware or handler
		return next(c)
	}
}
