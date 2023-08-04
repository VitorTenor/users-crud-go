package main

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/controller/routes"
	"github.com/VitorTenor/users-crud-go/src/controller/user"
	"github.com/VitorTenor/users-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application...")
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	svc := service.NewUserDomainService()
	userController := user.NewUserControllerInterface(svc)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error when trying to start the application", err)
	}
}
