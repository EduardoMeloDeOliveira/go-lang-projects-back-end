package repository

import (
	"database/sql"
	"fmt"
	"go-api/dto/message"
	"go-api/model"
)

type ProductRepository struct {
	dbConnection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		dbConnection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, product_price FROM product"
	rows, err := pr.dbConnection.Query(query)

	if err != nil {
		fmt.Println("error", err)
		return []model.Product{}, err
	}

	var productsList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println("error", err)
			return []model.Product{}, err
		}

		productsList = append(productsList, productObj)

	}

	rows.Close()
	return productsList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	query := "INSERT INTO product (product_name, product_price) VALUES ($1,$2) RETURNING ID , product_name , product_price"

	var savedProduct model.Product
	err := pr.dbConnection.QueryRow(query, product.Name, product.Price).Scan(
		&savedProduct.ID,
		&savedProduct.Name,
		&savedProduct.Price)

	if err != nil {
		return model.Product{}, err
	}

	return savedProduct, nil

}

func (pr *ProductRepository) DeleteProduct(id int64) (message.SuccessMessage, error) {
	query := "DELETE FROM product where ID = $1"

	result, err := pr.dbConnection.Exec(query, id)

	if err != nil {
		return message.SuccessMessage{}, err
	}

	rowsAffectd, err := result.RowsAffected()

	if err != nil {
		return message.SuccessMessage{}, nil
	}

	if rowsAffectd == 0 {
		return message.SuccessMessage{}, fmt.Errorf("nenhum produto com o ID: %d foi encontrado", id)
	}

	return message.SuccessMessage{Message: "Produto deletado com sucesso", StatusCode: 201}, nil

}
