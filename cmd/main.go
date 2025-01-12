package main

import (
	"log"
	"net/http"

	"github.com/vimalGopher/keerthys-trends/config"
	"github.com/vimalGopher/keerthys-trends/routes"
)

func main() {
	config.LoadConfig()
	router := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
