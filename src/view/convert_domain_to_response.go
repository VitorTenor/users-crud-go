package view

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
	"github.com/VitorTenor/users-crud-go/src/controller/user/model/response"
	"github.com/VitorTenor/users-crud-go/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	logger.Info("Converting domain to response")
	return response.UserResponse{
		ID:    userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

func ConvertDomainToResponseList(userDomainList []model.UserDomainInterface) []response.UserResponse {
	logger.Info("Converting domain list to response list")

	var userResponseList []response.UserResponse

	for _, user := range userDomainList {
		userResponseList = append(userResponseList, ConvertDomainToResponse(user))
	}

	return userResponseList
}
