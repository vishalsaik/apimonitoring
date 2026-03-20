package models

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTClaims struct {
	UserID   primitive.ObjectID  `json:"userId"`
	Username string              `json:"username"`
	Email    string              `json:"email"`
	Role     Role                `json:"role"`
	ClientID *primitive.ObjectID `json:"clientId,omitempty"`
	jwt.RegisteredClaims
}
