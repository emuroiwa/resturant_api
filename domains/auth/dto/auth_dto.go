package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name         string `json:"name" validate:"required,min=2,max=100"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	RestaurantID string `json:"restaurantid" validate:"required,uuid"`
}
