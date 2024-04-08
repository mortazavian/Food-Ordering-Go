package handlers

import (
	"Food-Ordering/internal/models"
	"Food-Ordering/internal/repository"
	"Food-Ordering/internal/utils"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
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

func AddMenuItemHandler(c echo.Context) error {
	menuItem := new(models.MenuItem)
	err := c.Bind(menuItem)
	if err != nil {
		return err
	}

	loggedRestaurant := c.Get("restaurant").(*models.Restaurant)

	menuItem.RestaurantId = loggedRestaurant.ID

	err = repository.CreateMenuItem(menuItem)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "food added to the database successfully"})
}

func GetAllMenuItemHandler(c echo.Context) error {
	// The restaurant can see its menu item only
	restaurant := c.Get("restaurant").(*models.Restaurant)

	menuItems, err := repository.RestaurantAllMenuItems(*restaurant)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, menuItems)

}

func CreateRestaurantWorkingDayHandler(c echo.Context) error {
	restaurantWorkingDay := new(models.RestaurantWorkingDay)
	err := c.Bind(restaurantWorkingDay)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid JSON format"})
	}

	restaurantWorkingDay.RestaurantId = c.Get("restaurant").(*models.Restaurant).ID

	err = repository.CreateRestaurantWorkingDay(restaurantWorkingDay)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "problem with database"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "restaurant weekday added successfully"})
}

func CreateRestaurantWorkingTimeHandler(c echo.Context) error {
	workingDayId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "the request does not have id as path parameter "})
	}

	workingTime := new(models.RestaurantWorkingTime)
	err = c.Bind(&workingTime)
	if err != nil {
		return err
	}

	//fmt.Println("------------")
	//fmt.Println(workingTime.OpensAt)
	//fmt.Println("------------")

	workingDay, err := repository.GetRestaurantWorkingDayByID(workingDayId)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "there is no working day which has the provided id"})
	}

	if workingDay.RestaurantId != c.Get("restaurant").(*models.Restaurant).ID {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "you are not the owner of this working day"})
	}

	workingTime.RestaurantWorkingDayId = uint(workingDayId)

	err = repository.CreateRestaurantWorkingTime(workingTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "problem storing working time in database"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "working time added successfully"})
}
