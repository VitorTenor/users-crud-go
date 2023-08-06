package service

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIdServices(userId string) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init FindUserById service",
		zap.String("journey", "findUserById"),
	)

	return ud.userRepository.FindUserById(userId)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init FindUserByEmail service",
		zap.String("journey", "findUserByEmail"),
	)

	return ud.userRepository.FindUserByEmail(email)
}
