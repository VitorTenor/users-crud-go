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
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_error.Err)
	UpdateUserServices(string, model.UserDomainInterface) *rest_error.Err
	FindAllUsersServices() ([]model.UserDomainInterface, *rest_error.Err)
	FindUserByIdServices(string) (model.UserDomainInterface, *rest_error.Err)
	FindUserByEmailServices(string) (model.UserDomainInterface, *rest_error.Err)
	DeleteUserServices(string) *rest_error.Err
	LoginUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_error.Err)
}
