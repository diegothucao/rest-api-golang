package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Token struct {
	UserId string
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}
