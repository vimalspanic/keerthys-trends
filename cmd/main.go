package main

import (
	"log"
	"net/http"

	"github.com/go-delve/delve/pkg/config"
	"github.com/vimalspanic/keerthys-trends/routes"
)

func main() {
	config.LoadConfig()
	router := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
