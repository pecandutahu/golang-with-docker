package main

import (
	"log"
	"product/internal/adapters/db"
	"product/internal/adapters/http"
	"product/internal/adapters/middleware"
	"product/internal/service"
	"product/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	//Connect to mysql db
	mysqlDB, err := database.ConnectMySQL()
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %s", err.Error())
	}
	// defer mysqlDB.Close()

	// Connect to MongoDB
	mongoClient, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}

	// Setup Product Repo and service

	productRepo := db.NewProductRepositoryDB(mysqlDB)
	productService := service.NewProductService(productRepo)

	// Setup Middleware
	app.Use(middleware.MonitorFunctionPerformance(mongoClient, "function_performance"))

	// Setup Routes
	http.SetupRoutes(app, productService, mongoClient)

	// Start Server
	log.Fatal(app.Listen(":8080"))
}
