package user

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/validation"
	"github.com/VitorTenor/users-crud-go/src/controller/user/model/request"
	"github.com/VitorTenor/users-crud-go/src/model"
	"github.com/VitorTenor/users-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser Controller",
		zap.String("journey", "loginUser"),
	)

	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error when trying to bind JSON",
			err,
			zap.String("journey", "loginUser"),
		)

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)

	resuult, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error when trying to call login user service",
			err,
			zap.String("journey", "loginUser"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("LoginUser Controller OK",
		zap.String("id", resuult.GetId()),
		zap.String("journey", "loginUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(resuult))
}
