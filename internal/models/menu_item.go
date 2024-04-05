package models

import "gorm.io/gorm"

type MenuItem struct {
	gorm.Model
	RestaurantId uint
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantId;references:ID"`
}
