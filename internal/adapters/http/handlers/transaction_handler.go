package handlers

import (
	"cashflow-go/internal/core/entities"
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

	transactions, err := th.ts.GetTransactions(userID)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Transactions not found"})
	}

	frequencies, err := th.fs.FindAllFrequencies()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to load frequencies"})
	}

	return c.Render(200, "dashboard", map[string]interface{}{
		"Transactions": transactions,
		"Frequencies":  frequencies,
	})
}

func (th *TransactionHandler) CreateTransaction(c echo.Context) error {
	transaction := new(entities.Transaction)

	fmt.Println(c.FormValue("category"))
	fmt.Println(c.FormValue("transaction_type"))
	fmt.Println(c.FormValue("frequency"))
	fmt.Println(c.FormValue("amount"))

	userID, ok := c.Get("user_id").(uint)
	if !ok || userID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Unauthorized"})
	}

	transaction.UserID = userID

	fmt.Printf("transac : %v\n\n", transaction)

	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := th.ts.CreateTransaction(transaction); err != nil {
		return c.JSON(404, err.Error())
	}

	if c.Request().Header.Get("HX-Request") != "" {
		c.Response().Header().Set("HX-Redirect", "/dashboard")
	}

	return c.NoContent(http.StatusOK)
}
