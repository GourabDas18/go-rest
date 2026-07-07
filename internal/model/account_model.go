package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	Name          string `gorm:"not null;size=100"`
	AccountNumber string
	UserId        uint   `gorm:"not null;index"`
	Balance       string `gorm:"default=0"`
	IsActive      bool   `gorm:"default=true"`
}
