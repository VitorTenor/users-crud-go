package model

import (
	"github.com/VitorTenor/users-crud-go/src/configuration/rest_error"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_error.Err) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_error.NewInternalServerError("Error when trying to generate token", err)
	}

	return tokenString, nil
}
