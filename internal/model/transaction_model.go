package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	Amount      string `gorm:"not null;default=0"`
	Description string `gorm:"default=''"`
	UserId      uint   `gorm:"not null;index"`
	CategoryId  uint   `gorm:"not null;index"`
	Type        uint   `gorm:"not null;index"`
	AccountId   uint   `gorm:"not null;index"`
}
