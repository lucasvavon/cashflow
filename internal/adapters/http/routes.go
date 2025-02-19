package http

import (
	"cashflow-backend/internal/adapters/http/handlers"
	"cashflow-backend/internal/core/services"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, us *services.UserService, ts *services.TransactionService) {

	// Handlers
	userHandler := handlers.NewUserHandler(us)
	transactionHandler := handlers.NewTransactionHandler(ts)
	authHandler := handlers.NewAuthHandler(us)

	// Middleware
	/*protected := app.Group("", middlewares.AuthMiddleware(ss))
	 */
	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return transactionHandler.GetTransactions(c, c.Locals("userID").(uint))
	})

	app.Post("/registration", func(c *fiber.Ctx) error {
		return userHandler.PostUser(c)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return userHandler.GetUsers(c)
	})

	app.Post("/auth", authHandler.Auth)

	app.Get("/transactions", func(c *fiber.Ctx) error {
		return transactionHandler.GetTransactions(c, c.Locals("userID").(uint))
	})

}
