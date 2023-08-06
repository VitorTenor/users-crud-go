package service

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserServices(userId string) *rest_error.Err {
	logger.Info("Init DeleteUser Service",
		zap.String("journey", "deleteUser"),
	)

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error on delete user",
			err,
			zap.String("journey", "deleteUser"),
		)

		return err
	}

	logger.Info("User updated successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	return nil
}
