package ports

import "product/internal/domain"

type ProductService interface {
	CreateProduct(product domain.Product) error
	GetProductByID(id int) (domain.Product, error)
	UpdateProduct(product domain.Product) error
	DeleteProduct(id int) error
	GetAllProducts() ([]domain.Product, error)
}
