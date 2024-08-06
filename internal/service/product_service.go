package service

import (
	"product/internal/domain"
	"product/internal/ports"
)

type ProductService struct {
	Repository ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &ProductService{Repository: repo}
}

func (s *ProductService) CreateProduct(product domain.Product) error {
	return s.Repository.Save(product)
}

func (s *ProductService) GetProductByID(id int) (domain.Product, error) {
	return s.Repository.FindByID(id)
}

func (s *ProductService) UpdateProduct(product domain.Product) error {
	return s.Repository.Update(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.Repository.Delete(id)
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
	return s.Repository.FindAll()
}
