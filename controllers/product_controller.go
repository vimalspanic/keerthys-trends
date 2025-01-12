package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vimalspanic/keerthys-trends/models"
	"github.com/vimalspanic/keerthys-trends/services"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := services.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	err := services.CreateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}
