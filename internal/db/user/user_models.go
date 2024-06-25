package user

import (
	//"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// table names
const UserTName = "users"

type User struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name" validate:"required,min=3,max=32"`
	Email   string             `json:"email" bson:"email" validate:"required,email"`
	Age     int                `json:"age" bson:"age" validate:"gte=18,lte=100"`
	Country string             `json:"country" bson:"country" validate:"required"`
}
