package models

import (
	"gorm.io/gorm"
	"time"
)

type RestaurantWorkingDay struct {
	gorm.Model
	RestaurantId uint
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantId;references:ID"`
	WeekDay      time.Weekday
}
