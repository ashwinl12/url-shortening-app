package url

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlDbPkg struct {
	Logger    *logrus.Logger
	Db *mongo.Client
}

func NewUrlDbPkg( logger *logrus.Logger,  client *mongo.Client) *UrlDbPkg {
	return &UrlDbPkg{
		Logger: logger,
		Db: client,
	}
}