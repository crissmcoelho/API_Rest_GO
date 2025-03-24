package repository

import (
	"API_Rest_GO/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRespository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}	

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query:= "SELECT id, product_name, price FROM product"
		rows, err := pr.connection.Query(query)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}	

		var productList []model.Product
		var productObject model.Product

		for rows.Next(){
			err = rows.Scan(
				&productObject.ID,
				&productObject.Name,
				&productObject.Price)

				if err != nil {
					fmt.Println(err)
					return []model.Product{}, err
				}	

				productList = append(productList, productObject)
		}
		rows.Close()
		return productList, nil
	}
