package models

import (
	"time"
)

type Ingredient struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`

	Name  string  `json:"name"`
	Meals []*Meal `gorm:"many2many:meals_ingredients" json:"meals"`
}
