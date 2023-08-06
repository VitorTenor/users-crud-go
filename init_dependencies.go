package main

import (
	"github.com/VitorTenor/users-crud-go/src/controller/user"
	"github.com/VitorTenor/users-crud-go/src/model/repository"
	"github.com/VitorTenor/users-crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initUserControllerDependencies(database *mongo.Database) user.UserControllerInterface {

	repo := repository.NewUserRepository(database)
	svc := service.NewUserDomainService(repo)

	return user.NewUserControllerInterface(svc)
}
