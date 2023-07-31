package service

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
)

func (*userDomainService) FindUserById(string) (*model.UserDomainInterface, *rest_error.Err) {
	return nil, nil
}

func (*userDomainService) FindUserByEmail(string) (*model.UserDomainInterface, *rest_error.Err) {
	return nil, nil
}
