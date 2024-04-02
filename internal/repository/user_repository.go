package repository

import "Food-Ordering/internal/models"

func CreateUser(user *models.User) error {

	result := instance.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
