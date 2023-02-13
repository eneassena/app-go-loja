package service

import (
	"github.com/eneassena/app-go-loja/internal/products/domain"
)

type ProductsService struct {
	Repository domain.ProductsService
}

func NewProductsService(repository domain.ProductsService) domain.ProductsService {
	return &ProductsService{
		Repository: repository,
	}

}
func (service *ProductsService) FindAll() ([]domain.ProductRequest, error) {
	return service.Repository.FindAll()

}

func (service *ProductsService) FindByID(id int) (domain.ProductRequest, error) {
	return service.Repository.FindByID(id)
}

func (service *ProductsService) FindByName(name string) (domain.ProductRequest, error) {
	return domain.ProductRequest{}, nil
}

func (service *ProductsService) Create(product domain.ProductRequest) (domain.ProductRequest, error) {
	return service.Repository.Create(product)
}

func (service *ProductsService) Remove(product domain.ProductRequest) error {
	return nil
}

func (service *ProductsService) UpdateCount(product domain.ProductRequest) (domain.ProductRequest, error) {
	products := domain.ProductRequest{}
	return products, nil
}
