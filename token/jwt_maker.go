package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Claims, error) {
	claims, err := NewClaims(username, duration)
	if err != nil {
		return "", claims, err
	}

	jwToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwToken.SignedString([]byte(maker.secretKey))
	return token, claims, err
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Claims, error) {

	jwToken, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
	}
	if !jwToken.Valid {
		return nil, ErrInvalidToken
	}

	if claims, ok := jwToken.Claims.(*Claims); ok && jwToken.Valid {
		return claims, nil
	} else {
		return nil, ErrInvalidToken
	}
}
