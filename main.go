package main

import (
	//"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"go-mongo/domain"
	"go-mongo/application/services"
	"net/http"
)
// Criando cliente com conexão
//var client *mongo.Client


// Criando assinatura dos métodos

//func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
//	response.Header().Set("content-type", "application/json")
//	var person domain.Person
//	_ = json.NewDecoder(request.Body).Decode(&person)
//	collection := client.Database("thepolyglotdeveloper").Collection("people")
//	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//	collection.InsertOne(ctx, person)
//	fmt.Println(ctx)
//	json.NewEncoder(response).Encode(person)
//}

func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/person", services.CreatePersonEndpoint).Methods("POST")
	//router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	//router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	http.ListenAndServe(":8080", router)
}
