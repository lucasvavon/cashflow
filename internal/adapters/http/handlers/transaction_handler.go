package handlers

import (
	"cashflow-go/internal/core/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TransactionHandler struct {
	ts *services.TransactionService
}

func NewTransactionHandler(ts *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		ts: ts,
	}
}

func (th *TransactionHandler) GetTransactions(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	fmt.Printf("\n\nuser_id: %v\n\n\n", userID)
	if !ok || userID == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	trs, err := th.ts.GetTransactions(userID)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Transactions not found"})
	}

	return c.Render(200, "dashboard", map[string]interface{}{
		"Transactions": trs,
	})
}
