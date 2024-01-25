package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type AccessClaims struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}
