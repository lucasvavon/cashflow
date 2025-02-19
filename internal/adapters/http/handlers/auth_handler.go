package handlers

import (
	"cashflow-backend/internal/core/entities"
	"cashflow-backend/internal/core/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	us *services.UserService
}

func NewAuthHandler(us *services.UserService) *AuthHandler {
	return &AuthHandler{us}
}

func (ah *AuthHandler) Auth(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, err := ah.us.Authenticate(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{"token": token})
}
