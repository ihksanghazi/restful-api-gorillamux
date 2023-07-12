package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ihksanghazi/restful-api-gorillamux/controllers/product_controllers"
	"github.com/ihksanghazi/restful-api-gorillamux/models"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/api/products", product_controllers.Index).Methods("GET")
	r.HandleFunc("/api/products/{id}", product_controllers.Show).Methods("GET")
	r.HandleFunc("/api/products", product_controllers.Create).Methods("POST")
	r.HandleFunc("/api/products/{id}", product_controllers.Update).Methods("PUT")
	r.HandleFunc("/api/products/{id}", product_controllers.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000",r))
	
}