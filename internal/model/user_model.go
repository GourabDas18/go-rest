package model

import (
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:255;index:idx_email_password;uniqueIndex"`
	Password  string `gorm:"not null;index:idx_email_password"`
	CountryId uint   `gorm:"not null"`
	SearchKey string `gorm:"sizw:500"`
	IsActive  bool   `gorm:"default=true"`
}

type UserCreateReq struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,min=3,email" `
	Password  string `json:"password" validate:"required,min=6"`
	CountryId uint   `json:"countryId" validate:"required"`
}

type UserAuthReq struct {
	Email    string `json:"email" validate:"required,min=3,email" `
	Password string `json:"password" validate:"required,min=6"`
}

type UserResult struct {
	User
	CurrencySymbol string `json:"currency_symbol" validate:"required"`
}

type UserResponse struct {
	ID             uint    `json:"id" validate:"required"`
	Name           string  `json:"name" validate:"required"`
	CountryId      uint    `json:"countryId"`
	CurrencySymbol string  `json:"currency"`
	Email          string  `json:"email"`
	AuthToken      *string `json:"authToken"`
}

func UserResponseParser(user *UserResult) UserResponse {
	userResp := UserResponse{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		CountryId:      user.CountryId,
		CurrencySymbol: user.CurrencySymbol,
	}

	return userResp
}

func UserParseFromRequest(userReq *UserCreateReq) User {
	user := User{
		Name:      strings.TrimSpace(userReq.Name),
		Email:     strings.TrimSpace(userReq.Email),
		Password:  strings.TrimSpace(userReq.Password),
		CountryId: userReq.CountryId,
		SearchKey: strings.TrimSpace(strings.ToLower(userReq.Name)) + strings.TrimSpace(strings.ToLower(userReq.Email)),
	}

	return user
}
