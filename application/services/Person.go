package services

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mongo/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

// Criando assinatura dos métodos
var client *mongo.Client
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
func init(){

	// Conectando com o banco de dados
	clientOptions := options.Client().ApplyURI("mongodb://root:12345678@localhost:27017")
	// Gerando client referente a conexão
	client, _ = mongo.Connect(ctx, clientOptions)
}
func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	// Criando Objeto
	var person domain.Person
	// Verificando o tipo do Header
	response.Header().Set("content-type", "application/json")
	// Criando contexto para a chamada da api
	_ = json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	collection.InsertOne(ctx, person)
	fmt.Println(ctx)
	json.NewEncoder(response).Encode(person)
}