package service

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init CreateUser Model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	userDomainRepo, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error on create user", err, zap.String("journey", "createUser"))
		return nil, err
	}

	return userDomainRepo, nil
}
