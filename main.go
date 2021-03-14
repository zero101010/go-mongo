package main

import (
	"fmt"
	"go-mongo/application/services"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {

	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/customer", services.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer", services.GetCustomers).Methods("GET")
	http.ListenAndServe(":8080", router)
}
