package repository

import (
	"Food-Ordering/internal/models"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Restaurant{})
	if err != nil {
		return err
	}

	return nil
}
