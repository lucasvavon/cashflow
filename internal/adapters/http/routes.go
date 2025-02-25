package http

import (
	"cashflow-go/internal/adapters/http/handlers"
	"cashflow-go/internal/adapters/http/middlewares"
	"cashflow-go/internal/core/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRoutes(e *echo.Echo, us *services.UserService, ts *services.TransactionService, fs *services.FrequencyService) {

	// handlers
	userHandler := handlers.NewUserHandler(us)
	transactionHandler := handlers.NewTransactionHandler(ts, fs)
	authHandler := handlers.NewAuthHandler(us)

	// middlewares
	authMiddleware := middlewares.AuthMiddleware()

	// public routes
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "login")
	})
	e.GET("/registration", func(c echo.Context) error {
		return c.Render(200, "registration", nil)
	})

	e.POST("/registration", userHandler.PostUser)
	e.POST("/login", authHandler.Login)
	e.GET("/login", func(c echo.Context) error {
		_, err := c.Cookie("token")
		if err == nil {
			return c.Redirect(http.StatusSeeOther, "/dashboard")
		}
		return c.Render(200, "login", nil)
	})

	// protected routes
	protected := e.Group("", authMiddleware)
	protected.GET("/dashboard", transactionHandler.GetTransactions)
	protected.POST("/logout", authHandler.Logout)
}
