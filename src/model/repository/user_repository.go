package repository

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(databaseConnection *mongo.Database) UserRepository {
	return &userRepository{databaseConnection: databaseConnection}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_error.Err)
	FindAllUsers() ([]model.UserDomainInterface, *rest_error.Err)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_error.Err)
	FindUserById(userId string) (model.UserDomainInterface, *rest_error.Err)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_error.Err
	DeleteUser(userId string) *rest_error.Err
}
