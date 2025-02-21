package entities

import "gorm.io/gorm"

type Frequency struct {
	gorm.Model
	Value string `gorm:"unique;not null"`
}
