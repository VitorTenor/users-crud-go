package service

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init LoginUser Service",
		zap.String("journey", "loginUser"),
	)

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		logger.Error("Error when trying to find user by email and password",
			err,
			zap.String("journey", "loginUser"),
		)
		return nil, err
	}

	logger.Info("LoginUser Service OK")
	logger.Info("User logged successfully",
		zap.String("Id ", user.GetId()),
		zap.String("LOGIN => ", userDomain.GetEmail()),
		zap.String("journey", "loginUser"),
	)

	return user, nil
}
