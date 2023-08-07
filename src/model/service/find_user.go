package service

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindAllUsersServices() ([]model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init FindAllUsers Service",
		zap.String("journey", "findAllUsers"),
	)

	return ud.userRepository.FindAllUsers()
}
func (ud *userDomainService) FindUserByIdServices(userId string) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init FindUserById Service",
		zap.String("journey", "findUserById"),
	)

	return ud.userRepository.FindUserById(userId)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init FindUserByEmail Service",
		zap.String("journey", "findUserByEmail"),
	)

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(email string, password string) (model.UserDomainInterface, *rest_error.Err) {
	logger.Info("Init FindUserByEmailAndPassword Service",
		zap.String("journey", "findUserByEmailAndPassword"),
	)
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
