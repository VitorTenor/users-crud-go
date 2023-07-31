package service

import (
	"fmt"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_error.Err {
	logger.Info("Init CreateUser Model", zap.String("journey", "create user"))
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}
