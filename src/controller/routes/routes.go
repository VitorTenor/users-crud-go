package routes

import (
	"github.com/VitorTenor/users-crud-go/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	//r.GET("/getUsers", getUsers)
	r.GET("/getUserById/:id", user.FindUserById)
	r.GET("/getUserByEmail/:email", user.FindUserByEmail)
	r.POST("/createUser", user.CreateUser)
	r.PUT("/updateUser/:id", user.UpdateUser)
	r.DELETE("/deleteUser/:id", user.DeleteUser)
}
