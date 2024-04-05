package middleware

import (
	"Food-Ordering/internal/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RestaurantJwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
		restaurantId, ok := claims["restaurant_id"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid user ID"})
		}

		// Retrieve user from repository using the user ID
		restaurant, err := repository.GetRestaurantByID(int(restaurantId))
		if err != nil || restaurant == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "user not found"})
		}

		// Set user in context for future use
		c.Set("restaurant", restaurant)

		// If the token is valid and the user is authenticated, proceed to the next middleware or handler
		return next(c)
	}
}
