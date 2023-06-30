package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewClaims(username string, duration time.Duration) (*Claims, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	return claims, nil
}

func (claims *Claims) Valid() error {
	if time.Now().After(claims.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}
