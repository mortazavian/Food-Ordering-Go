package repository

import (
	"Food-Ordering/internal/models"
)

func CreateUser(user *models.User) error {

	result := instance.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func GetUserByEmail(email string) (*models.User, error) {
	// Perform a database query to find a user by their email address
	// Example using GORM:
	var user models.User

	if err := instance.Where("email = ?", email).First(&user).Error; err != nil {
		// Handle the error, user not found, or other database-related issues

		return nil, err
	}

	// Return the user if found
	return &user, nil
}
