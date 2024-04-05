package models

import (
	"gorm.io/gorm"
)

type UserAddress struct {
	gorm.Model
	City        string `json:"city"`
	AddressLine string `json:"address_line"`
	UserId      uint
	User        User `gorm:"foreignKey:UserId;references:ID"`
	XCordinate  int  `json:"x_cordinate"`
	YCordinate  int  `json:"y_cordinate"`
	IsDefault   bool `json:"is_default"`
}
