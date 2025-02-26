package sql

import (
	"cashflow-go/internal/core/entities"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"os"
)

type DB struct {
	Db *gorm.DB
}

func ConnectDb() *gorm.DB {

	if os.Getenv("ENV") != "prod" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Error loading .env file")
		}
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Printf("Failed to connect to database! %s\n", dsn)
	}

	errMi := db.AutoMigrate(&entities.User{}, &entities.Category{}, &entities.Transaction{}, &entities.Frequency{}, &entities.TransactionType{})

	db.Clauses(clause.OnConflict{DoNothing: true}).Create([]entities.Frequency{
		{Name: "one-time"},
		{Name: "annual"},
		{Name: "monthly"},
		{Name: "weekly"},
		{Name: "daily"},
	})
	
	db.Clauses(clause.OnConflict{DoNothing: true}).Create([]entities.TransactionType{
		{Name: "income"},
		{Name: "expense"},
	})

	db.Clauses(clause.OnConflict{DoNothing: true}).Create([]entities.Category{
		{Name: "Salary", Logo: "", Description: "Monthly or bi-weekly wages from employment"},
		{Name: "Government Benefits", Logo: "", Description: "Income from social security, unemployment, or other government programs"},
		{Name: "Gifts & Donations", Logo: "", Description: "Money received as gifts or donations"},
		{Name: "Other income", Logo: "", Description: "Other type of income"},
		{Name: "Rent", Logo: "", Description: "Monthly rent payment for housing"},
		{Name: "Groceries", Logo: "", Description: "Money spent on food and household essentials"},
		{Name: "Transportation", Logo: "", Description: "Expenses related to fuel, public transport, or vehicle maintenance"},
		{Name: "Entertainment", Logo: "", Description: "Spending on movies, concerts, and leisure activities"},
		{Name: "Utilities", Logo: "", Description: "Payments for electricity, water, internet, and other utilities"},
		{Name: "Health & Wellness", Logo: "", Description: "Medical expenses, insurance, and gym memberships"},
		{Name: "Other Expenses", Logo: "", Description: "Miscellaneous expenses not categorized elsewhere"},
	})

	if errMi != nil {
		fmt.Printf("Failed to migrate database! %s\n", dsn)
	}

	return db
}
