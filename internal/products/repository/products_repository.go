package repository

import (
	"database/sql"
	"errors"

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
	produtosList := []domain.ProductRequest{}

	rows, err := repository.Database.Query("SELECT name,type,count,price,category FROM products")
	if err != nil {
		return []domain.ProductRequest{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var product domain.ProductRequest
		err := rows.Scan(&product.Name, &product.Type, &product.Count, &product.Price, &product.Category.Name)
		if err != nil {
			return []domain.ProductRequest{}, err
		}
		produtosList = append(produtosList, product)
	}

	return produtosList, nil
}

func (repository *ProductsRepository) FindByID(id int) (domain.ProductRequest, error) {
	rows, err := repository.Database.Query("select id,name,type,count,price,category from products where id = ?", id)
	if err != nil {
		return domain.ProductRequest{}, err
	}

	defer rows.Close()

	if rows.Next() {
		var product domain.ProductRequest
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.Category.Name); err != nil {
			return domain.ProductRequest{}, err
		}
		return product, nil
	}
	return domain.ProductRequest{}, errors.New("product not found")

}
func (repository *ProductsRepository) FindByName(name string) (domain.ProductRequest, error) {
	query := "select name, type, count, price, category from products where name=?;"
	rows := repository.Database.QueryRow(query, name)

	productByName := domain.ProductRequest{}
	erro := rows.Scan(&productByName.Name, &productByName.Type, &productByName.Count, &productByName.Price, &productByName.Category.Name)
	if erro != nil {
		return domain.ProductRequest{}, erro
	}

	return productByName, nil
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
	query := "delete from products where id=?"
	result, erro := repository.Database.Exec(query, product.ID)
	if erro != nil {
		return erro
	}
	if _, erro = result.RowsAffected(); erro != nil {
		return erro
	}
	return nil
}

func (repository *ProductsRepository) UpdateCount(product domain.ProductRequest) error {
	query := "update products set count=? where id=?"
	result, erro := repository.Database.Exec(query, product.Count, product.ID)
	if erro != nil {
		return erro
	}

	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}
