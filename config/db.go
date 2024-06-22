package config

import (
	"context"
	"fmt"
	"time"

	log "go-url-shortner/pkg/logger"

	"github.com/sirupsen/logrus"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectToDB(logger *logrus.Logger) (*mongo.Client, context.Context, context.CancelFunc, error) {
	// define mongodb uri
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// set a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// connect to mongodb
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.LogError(err, "Error inside connectToDB: error while connecting to mongodb", map[string]interface{}{"URI":"mongodb://localhost:27017"})
		return nil, nil, nil, err
	}

	// ping the database to verify the conn
	if err := client.Ping(ctx, nil); err != nil {
		log.LogError(err, "Error inside connectToDB: error while verifying db connection", nil)
		return nil, nil, nil, err
	}

	fmt.Println("Connected to mongodb !!")
	return client, ctx, cancel, nil
}