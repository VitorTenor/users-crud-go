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

	user, _ := ud.userRepository.FindUserByEmail(userDomain.GetEmail())
	if user != nil {
		responseError := rest_error.NewBadRequestError("Email already exists", nil)

		logger.Error("Error on create user, email already exists",
			responseError,
			zap.String("email", userDomain.GetEmail()),
			zap.String("journey", "createUser"),
		)

		return nil, responseError
	}

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
