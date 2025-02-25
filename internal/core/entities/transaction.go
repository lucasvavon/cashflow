package entities

import (
	"gorm.io/gorm"
)

type (
	Transaction struct {
		gorm.Model
		UserID            uint            `json:"user_id" sql:"not null"`
		Amount            float32         `json:"amount" sql:"not null"`
		Description       string          `json:"description"`
		FrequencyID       uint            `json:"frequency_id" sql:"default:null"`
		Frequency         Frequency       `sql:"foreignKey:FrequencyID;references:ID"`
		CategoryID        uint            `json:"category_id" sql:"index;not null"`
		Category          Category        `sql:"foreignKey:CategoryID;references:ID"`
		TransactionTypeID uint            `json:"trans_type_id" sql:"index;not null"`
		TransactionType   TransactionType `sql:"foreignKey:TransactionTypeID;references:ID"`
	}

	Transactions []Transaction
)
