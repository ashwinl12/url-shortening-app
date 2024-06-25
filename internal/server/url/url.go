package url

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlSrvPkg struct {
	Logger *logrus.Logger
	Db *mongo.Database
	Ctx *context.Context
}

func NewUrlSrvPkg( logger *logrus.Logger, db *mongo.Database, ctx *context.Context ) *UrlSrvPkg {
	return &UrlSrvPkg{
		Logger: logger,
		Db: db,
		Ctx: ctx,
	}
}