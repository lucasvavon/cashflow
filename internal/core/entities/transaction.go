package entities

import (
	"gorm.io/gorm"
)

type (
	Transaction struct {
		gorm.Model
		UserID            uint            `json:"user_id" gorm:"not null"`
		Amount            float32         `json:"amount" gorm:"not null"`
		Description       string          `json:"description"`
		Recurring         bool            `gorm:"default:false"`
		FrequencyID       uint            `json:"frequency_id" gorm:"default:null"`
		Frequency         Frequency       `gorm:"foreignKey:FrequencyID;references:ID"`
		CategoryID        uint            `json:"category_id" gorm:"index;not null"`
		Category          Category        `gorm:"foreignKey:CategoryID;references:ID"`
		TransactionTypeID uint            `json:"trans_type_id" gorm:"index;not null"`
		TransactionType   TransactionType `gorm:"foreignKey:TransactionTypeID;references:ID"`
	}

	Transactions []Transaction
)
