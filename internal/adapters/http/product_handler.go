package http

import (
	"product/internal/domain"
	"product/internal/dto"
	"product/internal/ports"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	Service ports.ProductService
}

func NewProductHandler(service ports.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.Service.CreateProduct(product)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []dto.ErrorResponse
			for _, feildError := range validationErrors {
				errors = append(errors, dto.ErrorResponse{
					Field:   feildError.Field(),
					Message: feildError.Tag(),
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(errors)

		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	productResponse := dto.ToProductResponse(product)

	return c.Status(fiber.StatusCreated).JSON(productResponse)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product id"})
	}
	product, err := h.Service.GetProductByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}
	productResponse := dto.ToProductResponse(product)

	return c.JSON(productResponse)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product id"})
	}
	err := h.Service.UpdateProduct(product)
	if err != nil {

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []dto.ErrorResponse
			for _, feildError := range validationErrors {
				errors = append(errors, dto.ErrorResponse{
					Field:   feildError.Field(),
					Message: feildError.Tag(),
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(errors)

		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	productResponse := dto.ToProductResponse(product)

	return c.JSON(productResponse)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = h.Service.DeleteProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.Service.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	productResponses := dto.ToProductResponses(products)

	return c.JSON(productResponses)
}

func SetupRoutes(app *fiber.App, service ports.ProductService) {
	handler := NewProductHandler(service)
	// Define routes
	app.Post("/products", handler.CreateProduct)
	app.Get("/products/:id", handler.GetProductByID)
	app.Put("/products/:id", handler.UpdateProduct)
	app.Delete("/products/:id", handler.DeleteProduct)
	app.Get("/products", handler.GetAllProducts)
}
