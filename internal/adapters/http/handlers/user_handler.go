package handlers

import (
	"cashflow-backend/internal/core/entities"
	"cashflow-backend/internal/core/services"
	"github.com/gofiber/fiber/v2"
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

func (uh *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := uh.us.Users()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Users not found"})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (uh *UserHandler) GetUser(c *fiber.Ctx) error {
	var user *entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := uh.us.FindUserByEmail(user.Email)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not Found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (uh *UserHandler) PostUser(c *fiber.Ctx) error {
	var user *entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	exist, _ := uh.us.FindUserByEmail(user.Email)

	if exist != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "user with this email already exists"})
	}

	if err := uh.us.CreateUser(user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	err = uh.us.DeleteUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var user *entities.User
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Call the UserService to update the user.
	err = uh.us.DeleteUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    user,
	})
}
