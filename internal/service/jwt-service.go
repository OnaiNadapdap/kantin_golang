package service

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}


