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

type UserResponse struct {
	ID        uint    `json:"id" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	CountryId uint    `json:"country_id" validate:"required"`
	SearchKey string  `json:"search_key"`
	Token     *string `json:"token"`
}

func UserResponseParser(user *User) UserResponse {
	userResp := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CountryId: user.CountryId,
		SearchKey: user.SearchKey,
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
