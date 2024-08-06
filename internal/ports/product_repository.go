package ports

import "product/internal/domain"

type ProductRepository interface {
	Save(product domain.Product) error
	FindByID(id uint) (domain.Product, error)
	Update(product domain.Product) error
	Delete(id uint) error
	FindAll() ([]domain.Product, error)
}
