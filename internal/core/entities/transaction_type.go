package entities

type TransactionType struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"unique;not null"`
}
