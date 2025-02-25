package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" sql:"not null"`
	Logo        string `json:"logo" sql:"not null"`
	Description string `json:"description"`
}
