package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Id    uint
	Email string
	jwt.StandardClaims
}
