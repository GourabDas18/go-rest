package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:255;index:idx_email_password;uniqueIndex"`
	Password  string `gorm:"not null;index:idx_email_password"`
	CountryId uint   `gorm:"not null"`
	SearchKey string `gorm:"sizw:500"`
	IsActive  bool   `gorm:"default=true"`
}
