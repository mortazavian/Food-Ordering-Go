package models

import (
	"gorm.io/gorm"
	"time"
)

type RestaurantWorkingDay struct {
	gorm.Model
	RestaurantId uint         `json:"restaurant_id"`
	Restaurant   Restaurant   `gorm:"foreignKey:RestaurantId;references:ID"`
	WeekDay      time.Weekday `json:"week_day"`
}
