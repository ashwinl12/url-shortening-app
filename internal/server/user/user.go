package user

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSrvPkg struct {
	Logger   *logrus.Logger
	Db       *mongo.Database
	Ctx      *context.Context
	Validate *validator.Validate
}

func NewUserSrvPkg(logger *logrus.Logger, db *mongo.Database, ctx *context.Context, validate *validator.Validate) *UserSrvPkg {
	return &UserSrvPkg{
		Logger: logger,
		Db:     db,
		Ctx:    ctx,
		Validate: validate,
	}
}
