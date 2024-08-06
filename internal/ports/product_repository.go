package ports

import "product/internal/domain"

type ProductRepository interface {
	Save(product domain.Product) error
	FindByID(id int) (domain.Product, error)
	Update(product domain.Product) error
	Delete(id int) error
	FindAll() ([]domain.Product, error)
}
