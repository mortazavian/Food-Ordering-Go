package repository

import "Food-Ordering/internal/models"

func CreateRestaurantWorkingDay(day *models.RestaurantWorkingDay) error {
	result := instance.Create(&day)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetRestaurantWorkingDayByID(id int) (*models.RestaurantWorkingDay, error) {
	var restaurantWorkingDay models.RestaurantWorkingDay

	resault := instance.First(&restaurantWorkingDay, id)
	if resault.Error != nil {
		return nil, resault.Error
	}

	return &restaurantWorkingDay, nil

}
