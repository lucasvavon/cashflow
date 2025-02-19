package main

import (
	routes "cashflow-backend/internal/adapters/http"
	"cashflow-backend/internal/adapters/repositories/postgre"
	"cashflow-backend/internal/core/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {

	// Fiber instance
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Vue.js address
		AllowMethods: "GET,POST,PUT,DELETE",   // Allowing all methods
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// conn to postgresgl instance
	db := postgre.ConnectDb()
	db = db.Debug()

	// init store repositories
	userStore := postgre.NewGormUserRepository(db)
	transactionStore := postgre.NewGormTransactionRepository(db)

	// init services
	userService := services.NewUserService(userStore)
	transactionService := services.NewTransactionService(transactionStore)

	routes.InitRoutes(app, userService, transactionService)

	log.Fatal(app.Listen("localhost:1323"))
}
