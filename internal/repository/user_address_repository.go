package repository

import "Food-Ordering/internal/models"

func CreateUserAddress(userAddress *models.UserAddress) error {
	result := instance.Create(&userAddress)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
