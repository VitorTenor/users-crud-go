package user

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser Controller",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorRest := rest_error.NewBadRequestError("Invalid id, must be a hex value", nil)
		c.JSON(errorRest.Code, errorRest)
	}

	err := uc.service.DeleteUserServices(userId)
	if err != nil {
		logger.Error("Error when trying to delete user",
			err,
			zap.String("journey", "deleteUser"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusOK)
}
