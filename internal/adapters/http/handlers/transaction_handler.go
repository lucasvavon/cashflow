package handlers

import (
	"cashflow-go/internal/core/services"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	ts *services.TransactionService
}

func NewTransactionHandler(ts *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		ts: ts,
	}
}

func (th *TransactionHandler) GetTransactions(c *fiber.Ctx, userID uint) error {
	/*userID, ok := c.Locals("userID").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}*/

	trs, err := th.ts.GetTransactions(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Transactions not found"})
	}

	return c.Status(fiber.StatusOK).JSON(trs)
}
