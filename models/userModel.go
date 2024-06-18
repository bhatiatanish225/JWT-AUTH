package models

import (
	"go.mongodb.org/mongo-driver/bson/primitve"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID         primitve.ObjectID `bson:"_id"`
	First_Name *string           `json:"first_name" validate:"required,min=2,max=100"`
	Last_Name  *string           `json:"last_name" validate:"required,min=2,max=100"`
	Email      *string           `json:"email" validate:"email ,required"`
	Password   *string           `json:"Password" validate:"required"`
	Phone      *string           `json:"phone" validate:"required"`
}
