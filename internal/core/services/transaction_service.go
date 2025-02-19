package services

import (
	"cashflow-backend/internal/core/entities"
	"cashflow-backend/internal/ports"
)

type TransactionService struct {
	tr ports.TransactionRepository
}

func NewTransactionService(tr ports.TransactionRepository) *TransactionService {
	return &TransactionService{tr: tr}
}

func (s *TransactionService) GetTransactions(userID uint) (*entities.Transactions, error) {
	return s.tr.GetTransactions(userID)
}

func (s *TransactionService) GetTransaction(userID uint, id uint) (*entities.Transaction, error) {
	return s.tr.GetTransaction(userID, id)
}

func (s *TransactionService) GetTransactionsByDateRange(userID uint, startDate string, endDate string) (*entities.Transactions, error) {
	return s.tr.GetTransactionsByDateRange(userID, startDate, endDate)
}

func (s *TransactionService) GetTransactionsByCategory(userID uint, category entities.Category) (*entities.Transactions, error) {
	return s.tr.GetTransactionsByCategory(userID, category)
}

func (s *TransactionService) GetTransactionsByType(userID uint, transType entities.TransactionType) (*entities.Transactions, error) {
	return s.tr.GetTransactionsByType(userID, transType)
}

func (s *TransactionService) CreateTransaction(transaction *entities.Transaction) error {
	return s.tr.CreateTransaction(transaction)
}

func (s *TransactionService) UpdateTransaction(transaction *entities.Transaction) error {
	return s.tr.UpdateTransaction(transaction)
}

func (s *TransactionService) DeleteTransaction(id uint) error {
	return s.tr.DeleteTransaction(id)
}
