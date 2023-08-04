package repository

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(databaseConnection *mongo.Database) UserRepository {
	return &userRepository{databaseConnection: databaseConnection}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_error.Err)
}
