package service

import (
	"product/internal/domain"
	"product/internal/ports"

	"github.com/go-playground/validator/v10"
)

type ProductService struct {
	Repository ports.ProductRepository
	Validator  *validator.Validate
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &ProductService{
		Repository: repo,
		Validator:  validator.New(),
	}
}

func (s *ProductService) CreateProduct(product domain.Product) error {
	if err := s.Validator.Struct(product); err != nil {
		return err
	}
	return s.Repository.Save(product)
}

func (s *ProductService) GetProductByID(id uint) (domain.Product, error) {
	return s.Repository.FindByID(id)
}

func (s *ProductService) UpdateProduct(product domain.Product) error {
	if err := s.Validator.Struct(product); err != nil {
		return err
	}
	return s.Repository.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repository.Delete(id)
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
	return s.Repository.FindAll()
}
