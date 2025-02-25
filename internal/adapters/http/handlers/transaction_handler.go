package handlers

import (
	"cashflow-go/internal/core/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TransactionHandler struct {
	ts *services.TransactionService
	fs *services.FrequencyService
}

func NewTransactionHandler(ts *services.TransactionService, fs *services.FrequencyService) *TransactionHandler {
	return &TransactionHandler{
		ts: ts,
		fs: fs,
	}
}

func (th *TransactionHandler) GetTransactions(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)

	if !ok || userID == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	_, err := th.ts.GetTransactions(userID)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Transactions not found"})
	}

	frequencies, err := th.fs.FindAllFrequencies()
	for _, frequency := range frequencies {
		fmt.Printf("frequency: %v\n", frequency.Name)
	}
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to load frequencies"})
	}

	return c.Render(200, "dashboard", map[string]interface{}{
		"Frequencies": frequencies,
	})
}

func (th *TransactionHandler) CreateTransaction(c echo.Context) error {
	//userID, ok := c.Get("user_id").(uint)
	return nil
}
