package repository

import "Food-Ordering/internal/models"

func CreateRestaurantWorkingDay(day *models.RestaurantWorkingDay) error {
	result := instance.Create(&day)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
