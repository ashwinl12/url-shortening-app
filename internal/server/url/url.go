package url

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlSrvPkg struct {
	Logger *logrus.Logger
	Db *mongo.Database
}

func NewUrlSrvPkg( logger *logrus.Logger, db *mongo.Database ) *UrlSrvPkg {
	return &UrlSrvPkg{
		Logger: logger,
		Db: db,
	}
}