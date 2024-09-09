package models

import (
	"restaurant-api/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string    `gorm:"type:string" validate:"omitempty,max=100"`
	Email        string    `gorm:"size:100;unique;not null;validate:required,email"`
	Password     string    `gorm:"not null;validate:required"`
	Role         string    `gorm:"default:user"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.Password, err = utils.HashPassword(user.Password)
	return
}
