package dto

import (
	"errors"
	"gorm.io/gorm"
	"net/mail"
	"strings"
)

type UserDTO struct {
	gorm.Model
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm-password"`
}

func (user *UserDTO) Validate() error {
	if user.Password == "" || user.Email == "" {
		return errors.New("email and password can't be empty")
	}

	if user.Password != user.ConfirmPassword {
		return errors.New("passwords are not equal")
	}

	if strings.ContainsAny(user.Password, "\t\r\n") {
		return errors.New("email and password can't have spaces")
	}

	if len(user.Password) < 6 {
		return errors.New("password shorter than 6 characters")
	}

	if len(user.Password) > 72 {
		return errors.New("password longer than 72 characters")
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.New("invalid email address")
	}

	return nil
}
