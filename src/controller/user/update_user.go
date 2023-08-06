package user

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/configuration/validation"
	"github.com/VitorTenor/users-crud-go/src/controller/user/model/request"
	"github.com/VitorTenor/users-crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser Controller",
		zap.String("journey", "updateUser"),
	)

	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		logger.Error("Error when trying to bind JSON",
			err,
			zap.String("journey", "updateUser"),
		)

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorRest := rest_error.NewBadRequestError("Invalid id, must be a hex value", nil)
		c.JSON(errorRest.Code, errorRest)
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)

	err := uc.service.UpdateUserServices(userId, domain)
	if err != nil {
		logger.Error("Error when trying to create user",
			err,
			zap.String("journey", "updateUser"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully",
		zap.String("id", userId),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)
}
