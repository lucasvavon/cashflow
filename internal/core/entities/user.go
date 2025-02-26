package entities

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Email          string        `gorm:"unique;not null"`
		HashedPassword string        `gorm:"not null"`
		Transactions   []Transaction `gorm:"foreignKey:UserID"`
	}

	Users []User
)
