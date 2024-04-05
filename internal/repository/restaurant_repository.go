package repository

import (
	"Food-Ordering/internal/models"
	"gorm.io/gorm"
)

func GetRestaurantByEmail(email string) *models.Restaurant {

	var restaurant models.Restaurant

	if err := instance.Where("email = ?", email).First(&restaurant).Error; err != nil {
		return nil
	}

	return &restaurant
}

func CreateRestaurant(restaurant *models.Restaurant) error {
	result := instance.Create(&restaurant)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func AuthenticateRestaurant(email, password string) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := instance.Where("email = ?", email).First(&restaurant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			return nil, nil
		}

		return nil, err
	}

	if restaurant.Password != password {
		return nil, nil
	}

	return &restaurant, nil
}

func GetRestaurantByID(userID int) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	err := instance.Where("id = ?", userID).First(&restaurant).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// User not found
			return nil, nil
		}
		// Other database error
		return nil, err
	}
	return &restaurant, nil
}
