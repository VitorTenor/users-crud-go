package repository

import (
	"context"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"github.com/VitorTenor/users-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_error.Err {

	logger.Info("Init updateUser repository",
		zap.String("journey", "updateUser"),
	)

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertUserDomainToUserEntity(userDomain)
	userIdHex, err := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{"_id", userIdHex}}
	update := bson.D{{"$set", value}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error on update userDomain",
			err,
			zap.String("journey", "updateUser"),
		)

		return rest_error.NewInternalServerError("Error on update userDomain", err)
	}

	logger.Info("UpdateUser repository OK",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	return nil
}
