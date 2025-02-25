package sql

import (
	"cashflow-go/internal/core/entities"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db}
}

func (r *GormUserRepository) FindAllUsers() (*entities.Users, error) {
	var users entities.Users

	req := r.db.Find(&users)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("messages not found: %v", req.Error))
	}

	return &users, nil
}

func (r *GormUserRepository) FindUserByID(id uint) (*entities.User, error) {
	var user entities.User

	req := r.db.First(&user, id)
	if req.Error != nil {
		// Use fmt.Errorf for error formatting and return the zero value of models.User.
		return &entities.User{}, fmt.Errorf("user not found: %v", req.Error)
	}

	return &user, nil
}

func (r *GormUserRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	req := r.db.First(&user, "email = ?", email)
	if req.Error != nil {
		if errors.Is(req.Error, gorm.ErrRecordNotFound) {
			return nil, req.Error
		}
		return nil, fmt.Errorf("error fetching user: %v", req.Error)
	}
	return &user, nil
}

func (r *GormUserRepository) CreateUser(user *entities.User) error {
	req := r.db.Create(&user)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *GormUserRepository) DeleteUser(id uint) error {
	var user entities.User

	req := r.db.Unscoped().Delete(&user, &id)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *GormUserRepository) UpdateUser(id uint, user *entities.User) error {
	user.ID = id

	req := r.db.Save(user)

	if req.Error != nil {
		return req.Error
	}

	return nil
}
