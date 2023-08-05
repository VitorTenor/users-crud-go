package converter

import (
	"github.com/VitorTenor/users-crud-go/src/model"
	"github.com/VitorTenor/users-crud-go/src/model/repository/entity"
)

func ConvertUserEntityToUserDomain(userEntity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(userEntity.Email, userEntity.Password, userEntity.Name, userEntity.Age)
	domain.SetId(userEntity.ID.Hex())
	return domain
}
