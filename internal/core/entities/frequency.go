package entities

import "gorm.io/gorm"

type Frequency struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type Frequencies []Frequency
