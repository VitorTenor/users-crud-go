package repository

import (
	"context"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"github.com/VitorTenor/users-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_error.Err) {

	logger.Info("Init createuser repository", zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertUserDomainToUserEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error on insert userDomain", err, zap.String("journey", "createUser"))
		return nil, rest_error.NewInternalServerError("Error on insert userDomain", err)
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("Createuser repository OK", zap.String("journey", "createUser"))

	return converter.ConvertUserEntityToUserDomain(*value), nil
}
