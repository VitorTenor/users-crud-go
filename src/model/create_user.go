package model

import (
	"fmt"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_error.Err {
	logger.Info("Init CreateUser Model", zap.String("journey", "create user"))
	ud.EncryptPassword()
	fmt.Println(ud)
	return nil
}
