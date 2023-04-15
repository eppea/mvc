package main

import (
	"fmt"
	"mvc/config"
	"mvc/controller"
	"os"

	// "github.com/eppea/mvc/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Koneksi ke database

func main() {
	config.InitDB()
	//var err error

	// db.AutoMigrate(&model.Transaction{})
	// defer db.Close()

	// Init Echo framework
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/transactions", controller.GetAllTransactions)
	e.GET("/transactions/:id", controller.GetTransaction)
	e.POST("/transactions", controller.CreateTransaction)
	e.PUT("/transactions/:id", controller.UpdateTransaction)
	e.DELETE("/transactions/:id", controller.DeleteTransaction)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Start(fmt.Sprintf(":%s", port))
}
