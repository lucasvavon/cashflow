package mock

import (
	"cashflow-go/internal/core/entities"
	"errors"
)

type MockUserRepository struct {
	Users       *entities.Users
	UserByID    *entities.User
	UserByEmail *entities.User
	SaveError   error
	UpdateError error
	DeleteError error
	ExistsByE   bool
}

// GetUsers return all users
func (m *MockUserRepository) GetUsers() (*entities.Users, error) {
	return m.Users, nil
}

// GetUserByID return user by id
func (m *MockUserRepository) GetUserByID(id int) (*entities.User, error) {
	for _, user := range *m.Users {
		if user.ID == uint(id) {
			return &user, nil
		}
	}
	return &entities.User{}, errors.New("user not found")
}

// GetUserByEmail return user by email
func (m *MockUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	for _, user := range *m.Users {
		if user.Email == email {
			return &user, nil
		}
	}
	return &entities.User{}, errors.New("user not found")
}

// CreateUser return error if user creation was not successful
func (m *MockUserRepository) CreateUser(user *entities.User) error {
	if m.SaveError != nil {
		return m.SaveError
	}
	*m.Users = append(*m.Users, *user)
	return nil
}

// UpdateUser update user by id
func (m *MockUserRepository) UpdateUser(id int, user *entities.User) error {
	if m.UpdateError != nil {
		return m.UpdateError
	}
	for i, u := range *m.Users {
		if u.ID == uint(id) {
			(*m.Users)[i] = *user
			return nil
		}
	}
	return errors.New("user not found")
}

// DeleteUser delete user by id
func (m *MockUserRepository) DeleteUser(id int) error {
	if m.DeleteError != nil {
		return m.DeleteError
	}
	for i, user := range *m.Users {
		if user.ID == uint(id) {
			*m.Users = append((*m.Users)[:i], (*m.Users)[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
