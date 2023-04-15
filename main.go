package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/eppea/mvc/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sql.DB

func main() {
	var err error

	// Koneksi ke database
	db, err = sql.Open("mysql", "root:EANHHUFWsX2ocayI4WXW@tcp(containers-us-west-143.railway.app:6475)/railway")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
