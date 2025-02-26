package ports

import (
	"cashflow-go/internal/core/entities"
)

type TransactionRepository interface {
	CreateTransaction(transaction *entities.Transaction) error
	GetTransaction(userID uint, id uint) (*entities.Transaction, error)
	GetTransactionsByDateRange(userID uint, startDate string, endDate string) (*entities.Transactions, error)
	GetTransactionsByCategory(userID uint, category entities.Category) (*entities.Transactions, error)
	GetTransactionsByType(userID uint, transType entities.TransactionType) (*entities.Transactions, error)
	GetTransactions(userID uint) (*entities.Transactions, error)
	UpdateTransaction(transaction *entities.Transaction) error
	DeleteTransaction(id uint) error
	GetGlobalBalance(id uint) (float32, error)
}
