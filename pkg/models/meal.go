package models

import (
	"time"
)

type Meal struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`

	Name        string        `json:"name"`
	Ingredients []*Ingredient `gorm:"many2many:meals_ingredients;" json:"ingredients"`
}
