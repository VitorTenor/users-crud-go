package service

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(userId string, userDomain model.UserDomainInterface) *rest_error.Err {
	logger.Info("Init UpdateUser Service",
		zap.String("journey", "updateUser"),
	)

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error on update user",
			err,
			zap.String("journey", "updateUser"),
		)

		return err
	}

	logger.Info("User updated successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	return nil
}
