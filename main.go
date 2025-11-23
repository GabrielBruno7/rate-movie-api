package main

import (
	"crud/config"
	"crud/http/routes"
	"log"
	"net/http"
)

func main() {
	databaseConnection := config.SetupDatabase()
	defer databaseConnection.Close()

	router := routes.SetupRoutes(databaseConnection)

	log.Println("Servidor rodando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
