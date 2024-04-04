package repository

import (
	"Food-Ordering/internal/models"
	"gorm.io/gorm"
)

func CreateUser(user *models.User) error {

	result := instance.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func GetUserByEmail(email string) (*models.User, error) {

	var user models.User

	if err := instance.Where("email = ?", email).First(&user).Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	if err := instance.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			return nil, nil
		}

		return nil, err
	}

	if user.Password != password {
		return nil, nil
	}

	return &user, nil
}

func GetUserByID(userID int) (*models.User, error) {
	var user models.User
	err := instance.Where("id = ?", userID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// User not found
			return nil, nil
		}
		// Other database error
		return nil, err
	}
	return &user, nil
}
