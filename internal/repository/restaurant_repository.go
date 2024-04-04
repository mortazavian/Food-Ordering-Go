package repository

import "Food-Ordering/internal/models"

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
