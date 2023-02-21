package service

import (
	"github.com/eneassena/app-go-loja/internal/products/domain"
)

type ProductsService struct {
	Repository domain.ProductsRepository
}

func NewProductsService(repository domain.ProductsRepository) domain.ProductsService {
	return &ProductsService{
		Repository: repository,
	}
}

func (service *ProductsService) FindAll() ([]domain.ProductRequest, error) {

	products, erro := service.Repository.FindAll()
	if erro != nil {
		return []domain.ProductRequest{}, erro
	}
	return products, nil
}

func (service *ProductsService) FindByID(id int) (domain.ProductRequest, error) {
	product, err := service.Repository.FindByID(id)
	if err != nil {
		return domain.ProductRequest{}, err
	}
	return product, nil
}

func (service *ProductsService) FindByName(name string) (domain.ProductRequest, error) {
	productByName, err := service.Repository.FindByName(name)
	if err != nil {
		return domain.ProductRequest{}, err
	}
	return productByName, nil
}

func (service *ProductsService) Create(product domain.ProductRequest) (domain.ProductRequest, error) {
	productCreated, erro := service.Repository.Create(product)
	if erro != nil {
		return domain.ProductRequest{}, erro
	}
	return productCreated, nil
}

func (service *ProductsService) Remove(product domain.ProductRequest) error {
	prod, erro := service.Repository.FindByID(product.ID)
	if erro != nil {
		return erro
	}

	erro = service.Repository.Remove(prod)
	if erro != nil {
		return erro
	}
	return nil
}

func (service *ProductsService) UpdateCount(product domain.ProductRequest) error {
	prod, erro := service.Repository.FindByID(product.ID)
	if erro != nil {
		return erro
	}
	prod.Count = product.Count
	erro = service.Repository.UpdateCount(prod)
	if erro != nil {
		return erro
	}
	return nil
}
