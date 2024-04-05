package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RestaurantProfile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
