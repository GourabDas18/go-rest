package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaim struct {
	UserId    int    `json:"userId"`
	UserName  string `json:"userName"`
	CountryId int    `json:"countryId"`

	jwt.RegisteredClaims
}

func GetToken(userId int, userName string, countryId int) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	claim := CustomClaim{UserId: userId, UserName: userName, CountryId: countryId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-rest-buzzet",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenStr string) (*CustomClaim, error) {
	secretKey := os.Getenv("JWT_SECRET")
	claims := CustomClaim{}
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid auth token")
	}

	return &claims, nil
}
