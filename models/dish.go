package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dish struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string    `gorm:"size:100;not null;validate:required,max=100"`
	Description  string    `gorm:"size:255;validate:omitempty,max=255"`
	Price        float64   `gorm:"not null;validate:required,gt=0"`
	Image        string    `gorm:"size:255;validate:omitempty,url"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not null"`
}

func (dish *Dish) BeforeCreate(tx *gorm.DB) (err error) {
	dish.ID = uuid.New()
	return
}
