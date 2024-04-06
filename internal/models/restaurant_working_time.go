package models

import "gorm.io/gorm"

type RestaurantWorkingTime struct {
	gorm.Model
	RestaurantWorkingDayId uint
	RestaurantWorkingDay   RestaurantWorkingDay `gorm:"foreignKey:RestaurantWorkingDayId;references:ID"`
}
