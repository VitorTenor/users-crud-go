package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/VitorTenor/users-crud-go/src/configuration/logger"
)

type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetJsonValues() (string, error) {
	logger.Info("Init UserDomain json values")
	b, err := json.Marshal(ud)
	if err != nil {
		logger.Error("Error on marshal userDomain", err)
		return "", err
	}

	logger.Info("UserDomain json values OK")
	return string(b), nil
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetJsonValues() (string, error)
	EncryptPassword()
	SetId(string)
}

func (ud *userDomain) SetId(id string) {
	ud.ID = id
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
