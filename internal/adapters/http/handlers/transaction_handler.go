package handlers

import (
	"cashflow-go/internal/core/dto"
	"cashflow-go/internal/core/entities"
	"cashflow-go/internal/core/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	transactions, err := th.ts.GetTransactions(userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to load transactions")
	}

	frequencies, err := th.fs.FindAllFrequencies()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to load frequencies")
	}

	balance, err := th.ts.GetGlobalBalance(userID)
	return c.Render(200, "dashboard", map[string]interface{}{
		"Transactions": transactions,
		"Frequencies":  frequencies,
		"Balance":      balance,
	})
}

func (th *TransactionHandler) CreateTransaction(c echo.Context) error {
	t := new(dto.TransactionDTO)

	if err := c.Bind(t); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	categoryID, err := strconv.ParseUint(t.CategoryID, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid category ID")
	}

	frequencyID, err := strconv.ParseUint(t.FrequencyID, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid frequency ID")
	}

	amount, err := strconv.ParseFloat(t.Amount, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid amount")
	}
	
	if t.TransactionType == "expense" {
		amount = -amount
	}

	transaction := entities.Transaction{
		UserID:      c.Get("user_id").(uint),
		Amount:      float32(amount),
		FrequencyID: uint(frequencyID),
		CategoryID:  uint(categoryID),
	}

	if err := th.ts.CreateTransaction(&transaction); err != nil {
		return c.String(404, err.Error())
	}

	if c.Request().Header.Get("HX-Request") != "" {
		c.Response().Header().Set("HX-Redirect", "/dashboard")
	}

	return c.NoContent(http.StatusOK)
}
