package routes

import (
	"github.com/gorilla/mux"
	"github.com/vimalspanic/keerthys-trends/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	return router
}
