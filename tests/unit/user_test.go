package unit

import (
	"cashflow-go/internal/adapters/repositories/mock"
	"cashflow-go/internal/core/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockRepo = &mock.MockUserRepository{
	Users: &entities.Users{
		{ID: 1, Email: "lucas@example.com", Password: "password"},
		{ID: 2, Email: "alice@example.com", Password: "password"},
	},
}

func TestGetUsers(t *testing.T) {
	users, err := mockRepo.GetUsers()
	assert.NoError(t, err)
	assert.Len(t, *users, 2)
}

func TestGetUserByID(t *testing.T) {
	user, err := mockRepo.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "lucas@example.com", user.Email)
}

func TestGetUserByIDNotFound(t *testing.T) {
	_, err := mockRepo.GetUserByID(99)
	assert.Error(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	user, err := mockRepo.GetUserByEmail("lucas@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "lucas@example.com", user.Email)
}

func TestGetUserByEmailNotFound(t *testing.T) {
	_, err := mockRepo.GetUserByEmail("sdgsd@sdfg")
	assert.Error(t, err)
}

func TestCreateUser(t *testing.T) {
	newUser := &entities.User{ID: 3, Email: "john@example.com", Password: "password"}
	err := mockRepo.CreateUser(newUser)
	assert.NoError(t, err)
	assert.Len(t, *mockRepo.Users, 3)
}

func TestUpdateUser(t *testing.T) {
	updatedUser := &entities.User{ID: 1, Email: "lucas.new@example.com", Password: "newpassword"}
	err := mockRepo.UpdateUser(1, updatedUser)
	assert.NoError(t, err)
	assert.Equal(t, "lucas.new@example.com", (*mockRepo.Users)[0].Email)
}

func TestDeleteUser(t *testing.T) {
	err := mockRepo.DeleteUser(1)
	assert.NoError(t, err)
	assert.Len(t, *mockRepo.Users, 2)
}
