package entities

import "gorm.io/gorm"

type TransactionType struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
