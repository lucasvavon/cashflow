package http

import (
	"cashflow-go/internal/adapters/http/handlers"
	"cashflow-go/internal/adapters/http/middlewares"
	"cashflow-go/internal/core/services"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, us *services.UserService, ts *services.TransactionService) {

	// Handlers
	userHandler := handlers.NewUserHandler(us)
	transactionHandler := handlers.NewTransactionHandler(ts)
	authHandler := handlers.NewAuthHandler(us)

	// Middleware
	authMiddleware := middlewares.AuthMiddleware()

	// Public routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})
	e.GET("/registration", func(c echo.Context) error {
		return c.Render(200, "registration", nil)
	})
	e.GET("/login", func(c echo.Context) error {
		return c.Render(200, "login", nil)
	})

	e.POST("/registration", userHandler.PostUser)
	e.POST("/login", authHandler.Login)

	protected := e.Group("", authMiddleware)
	protected.GET("/dashboard", transactionHandler.GetTransactions)
	protected.POST("/logout", authHandler.Logout)
}
