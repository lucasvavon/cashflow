package services

import (
	"cashflow-backend/internal/core/entities"
	"cashflow-backend/internal/ports"
	"cashflow-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Users() (*entities.Users, error) {
	return s.repo.FindAllUsers()
}

func (s *UserService) FindUserByID(id uint) (*entities.User, error) {
	return s.repo.FindUserByID(id)
}

func (s *UserService) FindUserByEmail(email string) (*entities.User, error) {
	return s.repo.FindUserByEmail(email)
}

func (s *UserService) CreateUser(user *entities.User) error {

	if err := user.Validate(); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := entities.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	return s.repo.CreateUser(&u)
}

func (s *UserService) UpdateUser(id uint, user *entities.User) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := entities.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	return s.repo.UpdateUser(id, &u)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) Authenticate(email, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
