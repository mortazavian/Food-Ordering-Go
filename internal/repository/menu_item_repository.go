package repository

import (
	"Food-Ordering/internal/models"
)

func CreateMenuItem(menuItem *models.MenuItem) error {
	result := instance.Create(&menuItem)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func RestaurantAllMenuItems(restaurant models.Restaurant) ([]models.MenuItem, error) {
	var menuItems []models.MenuItem

	result := instance.Find(&menuItems, "restaurant_id = ?", restaurant.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return menuItems, nil
}
