package http

import (
	"product/internal/ports"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, productService ports.ProductService, mongoClient *mongo.Client) {
	productHandler := NewProductHandler(productService)
	monitoringHandler := NewMonitoringHandler(mongoClient)

	// Define product routes
	app.Post("/products", productHandler.CreateProduct)
	app.Get("/products/:id", productHandler.GetProductByID)
	app.Put("/products/:id", productHandler.UpdateProduct)
	app.Delete("/products/:id", productHandler.DeleteProduct)
	app.Get("/products", productHandler.GetAllProducts)

	// Define monitoring route
	app.Get("/monitoring", monitoringHandler.GetMonitoringData)
}
