package service

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"github.com/VitorTenor/users-crud-go/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_error.Err)
	UpdateUser(string, model.UserDomainInterface) *rest_error.Err
	FindUserById(string) (*model.UserDomainInterface, *rest_error.Err)
	FindUserByEmail(string) (*model.UserDomainInterface, *rest_error.Err)
	DeleteUser(string) *rest_error.Err
}
