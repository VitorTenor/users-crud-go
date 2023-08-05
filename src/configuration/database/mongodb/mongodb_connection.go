package mongodb

import (
	"context"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var (
	MONGODB_URI    = "MONGODB_URI"
	MONGODB_DBNAME = "MONGODB_DBNAME"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	logger.Info("Initializing mongodb connection")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(MONGODB_URI)))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	logger.Info("Mongodb connection initialized with success")
	return client.Database(os.Getenv(MONGODB_DBNAME)), nil
}
