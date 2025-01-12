package services

import (
	"github.com/vimalGopher/keerthys-trends/models"
	"github.com/vimalGopher/keerthys-trends/utils"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	rows, err := utils.DB.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func CreateProduct(product *models.Product) error {
	_, err := utils.DB.Exec("INSERT INTO products (name, description, price, stock) VALUES (?, ?, ?, ?)",
		product.Name, product.Description, product.Price, product.Stock)
	return err
}
