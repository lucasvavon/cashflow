package ports

import (
	"cashflow-backend/internal/core/entities"
)

type UserRepository interface {
	FindAllUsers() (*entities.Users, error)
	FindUserByID(id uint) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	CreateUser(user *entities.User) error
	UpdateUser(id uint, user *entities.User) error
	DeleteUser(id uint) error
}
