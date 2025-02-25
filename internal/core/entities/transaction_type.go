package entities

import "gorm.io/gorm"

type TransactionType struct {
	gorm.Model
	Value string `sql:"unique;not null"`
}
