package service

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
	userDomain model.UserDomainInterface
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_error.Err
	UpdateUser(string, model.UserDomainInterface) *rest_error.Err
	FindUserById(string) (*model.UserDomainInterface, *rest_error.Err)
	FindUserByEmail(string) (*model.UserDomainInterface, *rest_error.Err)
	DeleteUser(string) *rest_error.Err
}
