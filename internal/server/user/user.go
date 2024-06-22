package user

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSrvPkg struct {
	Logger *logrus.Logger
	Db     *mongo.Database
}

func NewUserSrvPkg(logger *logrus.Logger, db *mongo.Database) *UserSrvPkg {
	return &UserSrvPkg{
		Logger: logger,
		Db:     db,
	}
}
