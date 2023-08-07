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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) FindAllUsers() ([]model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init findAll repository",
		zap.String("journey", "findAll"),
	)

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	users, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		logger.Error("Error on find all users",
			err,
			zap.String("journey", "findAll"),
		)

		return nil, rest_error.NewInternalServerError("Error on find all users", err)
	}

	var usersDomain []model.UserDomainInterface

	for users.Next(context.Background()) {
		userEntity := &entity.UserEntity{}

		err := users.Decode(userEntity)
		if err != nil {
			logger.Error("Error on decode userEntity",
				err,
				zap.String("journey", "findAll"),
			)

			return nil, rest_error.NewInternalServerError("Error on decode userEntity", err)
		}

		userDomain := converter.ConvertUserEntityToUserDomain(*userEntity)

		usersDomain = append(usersDomain, userDomain)
	}

	logger.Info("FindAll repository OK",
		zap.String("journey", "findAll"),
	)

	return usersDomain, nil
}

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

	objectId, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{"_id", objectId}}
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

func (ur *userRepository) FindUserByEmailAndPassword(email string, password string) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init FindUserByEmailAndPassword Repository",
		zap.String("journey", "findUserByEmailAndPassword"),
	)

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{"email", email}, {"password", password}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := "Email or password invalid!"

			logger.Error(errorMessage, err,
				zap.String("journey", "findUserByEmailAndPassword"),
			)

			return nil, rest_error.NewForbiddenError(errorMessage)
		}

		errorMessage := "Error on find user by email and password"

		logger.Error(errorMessage, err,
			zap.String("journey", "findUserByEmailAndPassword"),
		)

		return nil, rest_error.NewInternalServerError(errorMessage, err)
	}

	logger.Info("FindUserByEmailAndPassword Repository OK",
		zap.String("journey", "findUserByEmailAndPassword"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertUserEntityToUserDomain(*userEntity), nil
}
