package model

import "github.com/VitorTenor/users-crud-go/src/configuration/rest_error"

type UserDomainInterface interface {
	GetId() string
	SetId(string)
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	EncryptPassword()
	GenerateToken() (string, *rest_error.Err)
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(name string, age int8) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUserLoginDomain(email string, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
