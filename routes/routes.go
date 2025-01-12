package routes

import (
	"github.com/gorilla/mux"
	"github.com/vimalGopher/keerthys-trends/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	return router
}
