package gorm

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

	err := gtr.db.Select("transactions.*, categories.name, categories.logo").
		Joins("categories").
		Where("transactions.user_id = ?", userID).
		Order("transactions.updated_at DESC").
		Preload("Category").
		Find(&transactions).Error

	if err != nil {
		return nil, fmt.Errorf("transactions not found: %v", err)
	}

	return &transactions, nil
}

func (gtr GormTransactionRepository) CreateTransaction(transaction *entities.Transaction) error {
	return gtr.db.Create(&transaction).Error
}

func (gtr GormTransactionRepository) GetTransaction(userID uint, id uint) (*entities.Transaction, error) {
	var transaction entities.Transaction

	err := gtr.db.Select("transactions.*, categories.name, categories.logo").
		Joins("JOIN categories ON categories.id = transactions.category_id").
		Where("transactions.user_id = ?", userID).
		Where("transactions.id = ?", id).
		Order("transactions.updated_at DESC").
		Preload("Category").
		Find(&transaction).Error

	if err != nil {
		return nil, fmt.Errorf("transactions not found: %v", err)
	}

	return &transaction, nil
}

func (gtr GormTransactionRepository) GetTransactionsByDateRange(userID uint, startDate string, endDate string) (*entities.Transactions, error) {
	var transactions entities.Transactions

	err := gtr.db.Select("transactions.*, categories.name, categories.logo").
		Joins("JOIN categories ON categories.id = transactions.category_id").
		Where("transactions.user_id = ?", userID).
		Where("transactions.updated_at BETWEEN ? AND ?", startDate, endDate).
		Order("transactions.updated_at DESC").
		Preload("Category").
		Find(&transactions).Error

	if err != nil {
		return nil, fmt.Errorf("transactions not found: %v", err)
	}

	return &transactions, nil
}

func (gtr GormTransactionRepository) GetTransactionsByCategory(userID uint, category entities.Category) (*entities.Transactions, error) {
	var transactions entities.Transactions

	err := gtr.db.Select("transactions.*, categories.name, categories.logo").
		Joins("JOIN categories ON categories.id = transactions.category_id").
		Where("transactions.user_id = ?", userID).
		Where("categories. = ?", category).
		Order("transactions.updated_at DESC").
		Preload("Category").
		Find(&transactions).Error

	if err != nil {
		return nil, fmt.Errorf("transactions not found: %v", err)
	}

	return &transactions, nil
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
