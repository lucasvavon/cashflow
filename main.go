package main

import (
	routes "cashflow-go/internal/adapters/http"
	"cashflow-go/internal/adapters/repositories/sql"
	"cashflow-go/internal/core/services"
	"cashflow-go/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Echo instance
	e := echo.New()
	e.Static("/static", "views/assets")
	e.Renderer, _ = views.NewTemplate()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// conn to postgresgl instance
	db := sql.ConnectDb()
	db = db.Debug()

	// init store repositories
	userStore := sql.NewGormUserRepository(db)
	transactionStore := sql.NewGormTransactionRepository(db)

	// init services
	userService := services.NewUserService(userStore)
	transactionService := services.NewTransactionService(transactionStore)

	routes.InitRoutes(e, userService, transactionService)

	e.Logger.Fatal(e.Start("localhost:1323"))
}
