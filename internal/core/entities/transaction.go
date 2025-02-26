package entities

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID            uint            `json:"user_id" gorm:"not null"`
	Amount            float32         `json:"amount" gorm:"not null"`
	FrequencyID       uint            `json:"frequency"`
	Frequency         Frequency       `gorm:"foreignKey:FrequencyID;references:ID;" json:"-"`
	CategoryID        uint            `json:"category" gorm:"not null;index"`
	Category          Category        `gorm:"foreignKey:CategoryID;references:ID;" json:"-"`
	TransactionTypeID uint            `json:"transaction_type" gorm:"not null;index"`
	TransactionType   TransactionType `gorm:"foreignKey:TransactionTypeID;references:ID;" json:"-"`
}
type Transactions []Transaction
