package mongodb

import (
	"context"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection() {
	logger.Info("Initializing mongodb connection")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	logger.Info("Mongodb connection initialized with success")
}
