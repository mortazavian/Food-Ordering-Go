package models

type Restaurant struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RestaurantProfile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
