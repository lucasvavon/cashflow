package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique,not null"`
	Logo        string `json:"logo" gorm:"not null"`
	Description string `json:"description"`
}
