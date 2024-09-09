package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Restaurant struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name   string    `gorm:"type:string" validate:"omitempty,max=100"`
	Dishes []Dish
	Users  []User
}

func (restaurant *Restaurant) BeforeCreate(tx *gorm.DB) (err error) {
	restaurant.ID = uuid.New()
	return
}
