package main

import (
	"context"
	"github.com/VitorTenor/users-crud-go/src/configuration/database/mongodb"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	logger.Info("Starting the application...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatal("Error when trying to connect to database", err)
	}

	userController := initUserControllerDependencies(databaseConnection)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error when trying to start the application", err)
	}
}
