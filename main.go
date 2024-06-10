package main

import (
	"github.com/emr3k0c/money-transfer-system/db"
	"github.com/emr3k0c/money-transfer-system/handlers"
	"github.com/emr3k0c/money-transfer-system/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()
	e := echo.New()

	// No Auth Needed
	e.POST("/login", handlers.Login)
	e.POST("/api/customers", handlers.CreateCustomer)
	e.GET("/customers", handlers.ListCustomers)
	e.GET("/accounts", handlers.ListAccounts)

	// Auth Needed
	r := e.Group("/api")
	r.Use(middleware.JWTMiddleware())
	r.PUT("/customers/:id", handlers.UpdateCustomer)
	r.POST("/accounts", handlers.CreateAccount)
	r.PUT("/accounts/:id", handlers.UpdateAccount)
	r.POST("/accounts/transfer", handlers.TransferMoney)

	e.Start(":8080")
}
