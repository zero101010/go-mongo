package services

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mongo/domain"
	"go-mongo/framework/database"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func CreateCustomer(response http.ResponseWriter, request *http.Request) {
	var database database.DB
	db := database.Connect()
	response.Header().Set("content-type", "application/json")
	var customer domain.Customer
	_ = json.NewDecoder(request.Body).Decode(&customer)
	result, _ := db.Collection("customer").InsertOne(context.TODO(), customer)
	json.NewEncoder(response).Encode(result)
}

func GetCustomers(response http.ResponseWriter, request *http.Request) {
	var client *mongo.Client
	response.Header().Set("content-type", "application/json")
	collection := client.Database("staging-db").Collection("customer")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var customers []domain.Customer
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var customer domain.Customer
		cursor.Decode(&customer)
		fmt.Println(customer)
		customers = append(customers, customer)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(customers)

}
