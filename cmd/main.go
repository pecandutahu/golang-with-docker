package main

import (
	"log"
	"product/internal/adapters/db"
	"product/internal/adapters/http"
	"product/internal/service"
	"product/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	//Connect to mysql db

	mysqlDB, err := database.ConnectMySQL()
	if err != nil {
		log.Fatal(err)
	}
	defer mysqlDB.Close()

	// Setup Product Repo and service

	productRepo := db.NewProductRepositoryDB(mysqlDB)
	productService := service.NewProductService(productRepo)

	// Setup Routes
	http.SetupRoutes(app, productService)

	// Start Server
	log.Fatal(app.Listen(":8080"))
}
