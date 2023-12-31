package routes

import (
	"github.com/VitorTenor/users-crud-go/src/controller/user"
	"github.com/VitorTenor/users-crud-go/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController user.UserControllerInterface) {
	r.GET("/getAllUsers", model.VerifyTokenMiddleare, userController.FindAllUsers)
	r.GET("/getUserById/:id", userController.FindUserById)
	r.GET("/getUserByEmail/:email", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:id", userController.UpdateUser)
	r.DELETE("/deleteUser/:id", userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
}
