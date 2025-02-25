package entities

import (
	"gorm.io/gorm"
	"net/mail"
	"strings"
)

type (
	User struct {
		gorm.Model
		Email           string        `json:"email" sql:"unique;not null"`
		Password        string        `json:"password" sql:"not null"`
		ConfirmPassword string        `json:"confirm-password" sql:"-"`
		Transactions    []Transaction `sql:"foreignKey:UserID"`
	}

	Users []User
)

func (user *User) Validate() error {
	if user.Password == "" || user.Email == "" {
		return ErrEmptyUserField
	}

	if user.Password != user.ConfirmPassword {
		return ErrPasswordEq
	}

	if strings.ContainsAny(user.Password, "\t\r\n") {
		return ErrFieldWithSpaces
	}

	if len(user.Password) < 6 {
		return ErrShortPassword
	}

	if len(user.Password) > 72 {
		return ErrLongPassword
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}
