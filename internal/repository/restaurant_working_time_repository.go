package repository

import (
	"Food-Ordering/internal/models"
)

func CreateRestaurantWorkingTime(workingTime *models.RestaurantWorkingTime) error {

	result := instance.Create(workingTime)
	if result.Error != nil {
		return instance.Error
	}

	return nil
}
