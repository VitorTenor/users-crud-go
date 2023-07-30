package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &UserDomain{email, password, name, age}
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_error.Err
	UpdateUser(string) *rest_error.Err
	FindUser(string) (*UserDomain, *rest_error.Err)
	DeleteUser(string) *rest_error.Err
}
