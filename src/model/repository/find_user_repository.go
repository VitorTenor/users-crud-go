package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"github.com/VitorTenor/users-crud-go/src/model/repository/entity"
	"github.com/VitorTenor/users-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"),
	)

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{"email", email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)

			logger.Error(errorMessage, err,
				zap.String("journey", "findUserByEmail"),
			)

			return nil, rest_error.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error on find user by email"

		logger.Error(errorMessage, err,
			zap.String("journey", "findUserByEmail"),
		)

		return nil, rest_error.NewInternalServerError(errorMessage, err)
	}

	logger.Info("FindUserByEmail repository OK",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertUserEntityToUserDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(userId string) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init findUserById repository",
		zap.String("journey", "findUserById"),
	)

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{"_id", userId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with this id: %s", userId)

			logger.Error(errorMessage, err,
				zap.String("journey", "findUserById"),
			)

			return nil, rest_error.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error on find user by id"

		logger.Error(errorMessage, err,
			zap.String("journey", "findUserById"),
		)

		return nil, rest_error.NewInternalServerError(errorMessage, err)
	}

	logger.Info("FcindUserById repository OK",
		zap.String("journey", "findUserById"),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertUserEntityToUserDomain(*userEntity), nil
}
