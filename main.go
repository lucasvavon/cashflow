package main

import (
	routes "cashflow-go/internal/adapters/http"
	"cashflow-go/internal/adapters/repositories/postgre"
	"cashflow-go/internal/core/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {

	// Fiber instance
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
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
