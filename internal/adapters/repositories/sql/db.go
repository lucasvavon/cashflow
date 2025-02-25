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
		{Value: "annual"},
		{Value: "monthly"},
		{Value: "weekly"},
		{Value: "daily"},
	})
	db.Clauses(clause.OnConflict{DoNothing: true}).Create([]entities.TransactionType{
		{Value: "income"},
		{Value: "expense"},
	})

	if errMi != nil {
		fmt.Printf("Failed to migrate database! %s\n", dsn)
	}

	return db
}
