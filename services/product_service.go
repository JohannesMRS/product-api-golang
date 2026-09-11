package services

import (
	"manajemen-product/models"
	"manajemen-product/repositories"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(input models.CreateProductInput) (models.Product, error)
	UpdateProduct(id int, input models.UpdateProductInput) (models.Product, error)
	DeleteProduct(id int) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) CreateProduct(input models.CreateProductInput) (models.Product, error) {
	product := models.Product{
		Name:  input.Name,
		Price: input.Price,
	}

	return s.repo.Create(product)
}

func (s *productService) UpdateProduct(id int, input models.UpdateProductInput) (models.Product, error) {
	product := models.Product{
		Name:  input.Name,
		Price: input.Price,
	}

	return s.repo.Update(id, product)
}

func (s *productService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
