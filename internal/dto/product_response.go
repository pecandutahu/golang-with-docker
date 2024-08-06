package dto

import (
	"product/internal/domain"
	"time"
)

type ProductResponse struct {
	ProductID   uint   `json:"product_id"`
	ProductCode string `gorm:"uniqueIndex;size:255"`
	ProductName string `gorm:"size:255"`
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToProductResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		ProductID:   product.ProductID,
		ProductCode: product.ProductCode,
		ProductName: product.ProductName,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ToProductResponses(products []domain.Product) []ProductResponse {
	productResponses := make([]ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = ToProductResponse(product)
	}
	return productResponses
}
