package repository

import (
	"context"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"os"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_error.Err) {

	logger.Info("Init createuser repository")
	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJsonValues()
	if err != nil {
		logger.Error("Error on marshal userDomain", err)
		return nil, rest_error.NewInternalServerError("Error on marshal userDomain", err)
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error on insert userDomain", err)
		return nil, rest_error.NewInternalServerError("Error on insert userDomain", err)
	}

	userDomain.SetId(result.InsertedID.(string))

	logger.Info("Createuser repository OK")

	return userDomain, nil
}
