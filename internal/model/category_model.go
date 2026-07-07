package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	Name        string `gorm:"not null;size:100"`
	Description string
	IsActive    string `gorm:"not null;size:100"`
	UserId      uint   `gorm:"default=-1;index"`
}
