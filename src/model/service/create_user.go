package service

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init CreateUser Model",
		zap.String("journey", "createUser"),
	)

	userDomain.EncryptPassword()

	userDomainRepo, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error on create user",
			err,
			zap.String("journey", "createUser"),
		)

		return nil, err
	}

	logger.Info("User created successfully",
		zap.String("id", userDomainRepo.GetId()),
		zap.String("journey", "createUser"),
	)

	return userDomainRepo, nil
}
