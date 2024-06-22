package user

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDbPkg struct {
	Logger    *logrus.Logger
	Db *mongo.Client
}

func NewUserDbPkg( logger *logrus.Logger,  client *mongo.Client) *UserDbPkg {
	return &UserDbPkg{
		Logger: logger,
		Db: client,
	}
}

