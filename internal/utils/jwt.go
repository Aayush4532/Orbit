package utils

import (
	"Orbit/configs"
	"Orbit/internal/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID              string `json:"_id"`
	EmailId         string `json:"emailId"`
	Role            string `json:"role"`
	IsEmailVerified bool   `json:"isEmailVerified"`
	IsActive        bool   `json:"isActive"`
	jwt.RegisteredClaims
}

func CreateJwtToken(user models.User) (string, error) {
	key := configs.LoadConfig().JWT_KEY.Key

	claims := Claims{
		ID:              user.ID.Hex(),
		EmailId:         user.EmailId,
		Role:            user.Role,
		IsEmailVerified: user.IsEmailVerified,
		IsActive:        user.IsActive,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJwtToken(tokenString string) (*Claims, error) {
	key := configs.LoadConfig().JWT_KEY.Key

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			// checking signing method is same as with we are creating.
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(key), nil
		},
	)

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token provided")
	}

	return claims, nil
}