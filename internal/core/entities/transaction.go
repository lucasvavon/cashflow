package entities

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID      uint      `gorm:"not null"`
	Amount      float32   `gorm:"not null"`
	FrequencyID uint      `gorm:"not null;index"`
	Frequency   Frequency `gorm:"foreignKey:FrequencyID;references:ID;"`
	CategoryID  uint      `gorm:"not null;index"`
	Category    Category  `gorm:"foreignKey:CategoryID;references:ID;"`
}
type Transactions []Transaction
