package ports

import "product/internal/domain"

type ProductService interface {
	CreateProduct(product domain.Product) error
	GetProductByID(id uint) (domain.Product, error)
	UpdateProduct(product domain.Product) error
	DeleteProduct(id uint) error
	GetAllProducts() ([]domain.Product, error)
}
