package repository

import (
	"database/sql"

	"github.com/eneassena/app-go-loja/internal/products/domain"
)

type ProductsRepository struct {
	Database *sql.DB
}

func NewProductsRepository(database *sql.DB) domain.ProductsRepository {
	return &ProductsRepository{
		Database: database,
	}
}

func (repository *ProductsRepository) FindAll() ([]domain.ProductRequest, error) {
	var produtosList []domain.ProductRequest

	rows, err := repository.Database.Query("SELECT id,name,price,count,category FROM products")
	if err != nil {
		return produtosList, err
	}

	defer rows.Close()
	for rows.Next() {
		var product domain.ProductRequest
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Category.Name)
		if err != nil {
			return produtosList, err
		}
		produtosList = append(produtosList, product)
	}

	return produtosList, nil
}
func (repository *ProductsRepository) FindByID(id int) (domain.ProductRequest, error) {
	rows, err := repository.Database.Query("select id,name,type,price,count,category from products where id = ?", id)
	if err != nil {
		return domain.ProductRequest{}, err
	}

	defer rows.Close()

	var product domain.ProductRequest
	if rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Price, &product.Count, &product.Category.Name); err != nil {
			return domain.ProductRequest{}, err
		}
	}
	return product, nil
}
func (repository *ProductsRepository) FindByName(name string) (domain.ProductRequest, error) {
	return domain.ProductRequest{}, nil
}
func (repository *ProductsRepository) Create(product domain.ProductRequest) (domain.ProductRequest, error) {
	stmt, err := repository.Database.Prepare("insert into products (name,type,count,price,category) VALUES(?,?,?,?,?)")
	if err != nil {
		return domain.ProductRequest{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&product.Name, &product.Type, &product.Count, &product.Price, &product.Category.Name)
	if err != nil {
		return domain.ProductRequest{}, err
	}
	return product, nil
}
func (repository *ProductsRepository) Remove(product domain.ProductRequest) error {
	return nil
}
func (repository *ProductsRepository) UpdateCount(product domain.ProductRequest) (domain.ProductRequest, error) {
	return domain.ProductRequest{}, nil
}
