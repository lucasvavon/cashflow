package entities

import "gorm.io/gorm"

type TransactionType struct {
	gorm.Model
	Value        string        `gorm:"unique;not null"`
	Transactions []Transaction `gorm:"foreignKey:TransactionTypeID"`
}
