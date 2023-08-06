package user

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById Controller",
		zap.String("journey", "findUserById"),
	)

	userId := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying validate UUID",
			err,
			zap.String("journey", "findUserById"),
		)

		errorMessage := rest_error.NewBadRequestError(
			"UserId is not a valid UUID",
			nil,
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByIdServices",
			err,
			zap.String("journey", "findUserById"),
		)

		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail Controller",
		zap.String("journey", "findUserByEmail"),
	)

	email := c.Param("email")
	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error("Error trying validate email",
			err,
			zap.String("journey", "findUserByEmail"),
		)

		errorMessage := rest_error.NewBadRequestError(
			"Email is not a valid email",
			nil,
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(email)
	if err != nil {
		logger.Error("Error trying to call findUserByEmailServices",
			err,
			zap.String("journey", "findUserByEmail"),
		)

		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
