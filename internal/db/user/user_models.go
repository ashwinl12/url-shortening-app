package user

import (
	//"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// table names
const UserTName = "users"

type User struct {
	ID      primitive.ObjectID `bson:"id" validate:"required"`
	Name    string             `bson:"name" validate:"required"`
	Email   string             `bson:"email" validate:"required,email"`
	Country string             `bson:"country" validate:"required"`
}
