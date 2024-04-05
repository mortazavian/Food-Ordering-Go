package repository

import "Food-Ordering/internal/models"

func CreateMenuItem(menuItem *models.MenuItem) error {
	result := instance.Create(&menuItem)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
