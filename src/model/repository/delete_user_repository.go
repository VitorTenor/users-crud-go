package repository

import (
	"context"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) DeleteUser(userId string) *rest_error.Err {

	logger.Info("Init DeleteUser repository",
		zap.String("journey", "deleteUser"),
	)

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, err := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{"_id", userIdHex}}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error on delete userDomain",
			err,
			zap.String("journey", "deleteUser"),
		)

		return rest_error.NewInternalServerError("Error on delete userDomain", err)
	}

	logger.Info("DeleteUser repository OK",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	return nil
}
