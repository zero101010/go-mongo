package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Criando estrutura
type Customer struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
