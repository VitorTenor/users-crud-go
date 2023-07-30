package user

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/validation"
	"github.com/VitorTenor/users-crud-go/src/controller/user/model/request"
	"github.com/VitorTenor/users-crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller", zap.String("journey", "create user"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error when trying to bind JSON", err, zap.String("journey", "create user"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		logger.Error("Error when trying to create user", err, zap.String("journey", "create user"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "create user"))

	c.String(http.StatusOK, "")
}
