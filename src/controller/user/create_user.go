package user

import (
	"fmt"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/VitorTenor/users-crud-go/src/controller/user/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_error.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields in the request body, error=%s ", err.Error()),
			nil,
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
