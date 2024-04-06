package models

import (
	"gorm.io/gorm"
	"time"
)

type RestaurantWorkingTime struct {
	gorm.Model
	RestaurantWorkingDayId uint
	RestaurantWorkingDay   RestaurantWorkingDay `gorm:"foreignKey:RestaurantWorkingDayId;references:ID"`
	OpensAt                time.Time
	ClosesAt               time.Time
}
