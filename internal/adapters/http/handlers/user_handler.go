package handlers

import (
	"cashflow-go/internal/core/dto"
	"cashflow-go/internal/core/entities"
	"cashflow-go/internal/core/services"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	us *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

func (uh *UserHandler) GetUsers(c echo.Context) error {
	users, err := uh.us.FindAllUsers()

	if err != nil {
		return c.JSON(404, map[string]string{"error": "FindAllUsers not found"})
	}

	return c.JSON(200, users)
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	var user *entities.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(500, err.Error())
	}

	user, err := uh.us.FindUserByEmail(user.Email)

	if err != nil {
		return c.JSON(404, map[string]string{"error": "User not Found"})
	}

	return c.JSON(200, user)
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	u := new(dto.UserDTO)

	if err := c.Bind(&u); err != nil {
		return c.JSON(404, map[string]string{"error": "Invalid request body"})
	}

	exist, _ := uh.us.FindUserByEmail(u.Email)

	if exist != nil {
		return c.JSON(404, map[string]string{"error": "user with this email already exists"})
	}

	if err := u.Validate(); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entities.User{
		Email:          u.Email,
		HashedPassword: string(hashedPassword),
	}

	if err := uh.us.CreateUser(&user); err != nil {
		return c.JSON(404, err.Error())
	}

	if c.Request().Header.Get("HX-Request") != "" {
		c.Response().Header().Set("HX-Redirect", "/login")
	}

	return c.NoContent(http.StatusOK)
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Invalid user ID"})
	}

	err = uh.us.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "User not found"})
	}

	return c.NoContent(200)
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	var user *entities.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Invalid user ID"})
	}

	// Call the UserService to update the user.
	err = uh.us.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "User not found"})
	}

	return c.JSON(201, user)
}
