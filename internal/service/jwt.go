package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaim struct {
	userId    int
	userName  string
	countryId int
	jwt.RegisteredClaims
}

func GetToken(userId int, userName string, countryId int) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	claim := CustomClaim{userId: userId, userName: userName, countryId: countryId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-rest-buzzet",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string) (*CustomClaim, error) {
	token, err := jwt.ParseWithClaims()
}
