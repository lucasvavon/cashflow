package entities

import (
	"gorm.io/gorm"
)

type (
	Transaction struct {
		gorm.Model
		UserID            uint    `json:"user_id" gorm:"not null"`
		Amount            float32 `json:"amount" gorm:"not null"`
		TransactionTypeID uint    `json:"trans_type_id" gorm:"not null"`
		Recurring         bool    `gorm:"default:false"`
		FrequencyID       uint    `json:"frequency_id" gorm:"default:null"`
		Description       string  `json:"description"`
		CategoryID        uint    `json:"category_id" gorm:"not null"`
	}

	Transactions []Transaction
)
