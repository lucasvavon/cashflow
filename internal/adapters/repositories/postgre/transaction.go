package postgre

import (
	"cashflow-go/internal/core/entities"
	"fmt"
	"gorm.io/gorm"
)

type GormTransactionRepository struct {
	db *gorm.DB
}

func NewGormTransactionRepository(db *gorm.DB) *GormTransactionRepository {
	return &GormTransactionRepository{db}
}

func (gtr *GormTransactionRepository) GetTransactions(userID uint) (*entities.Transactions, error) {
	var transactions entities.Transactions

	err := gtr.db.Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return nil, fmt.Errorf("transactions not found: %v", err)
	}

	return &transactions, nil
}

func (gtr GormTransactionRepository) CreateTransaction(transaction *entities.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (gtr GormTransactionRepository) GetTransaction(userID uint, id uint) (*entities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (gtr GormTransactionRepository) GetTransactionsByDateRange(userID uint, startDate string, endDate string) (*entities.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func (gtr GormTransactionRepository) GetTransactionsByCategory(userID uint, category entities.Category) (*entities.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func (gtr GormTransactionRepository) GetTransactionsByType(userID uint, transType entities.TransactionType) (*entities.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func (gtr GormTransactionRepository) UpdateTransaction(transaction *entities.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (gtr GormTransactionRepository) DeleteTransaction(id uint) error {
	//TODO implement me
	panic("implement me")
}
