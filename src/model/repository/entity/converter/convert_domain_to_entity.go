package converter

import (
	"github.com/VitorTenor/users-crud-go/src/model"
	"github.com/VitorTenor/users-crud-go/src/model/repository/entity"
)

func ConvertUserDomainToUserEntity(userDomain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
		Name:     userDomain.GetName(),
		Age:      userDomain.GetAge(),
	}
}
