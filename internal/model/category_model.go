package model

import (
	"strings"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	Name        string `gorm:"not null;size:100"`
	Description string
	IsActive    bool `gorm:"default:true"`
	UserId      uint `gorm:"default=-1;index"`
}

type CategoryReq struct {
	Name        string `json:"name" validate:"required"`
	UserId      uint   `json:"userId" validate:"required"`
	Description string `json:"description"`
	IsActive    *bool  `json:"isActive"`
}

type CategoryRes struct {
	Id          uint
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (c CategoryReq) ParseToCategory() Category {
	isActivePtr := true
	if c.IsActive != nil {
		isActivePtr = *c.IsActive
	}
	return Category{
		Name:        strings.TrimSpace(c.Name),
		Description: strings.TrimSpace(c.Description),
		IsActive:    isActivePtr,
		UserId:      c.UserId,
	}
}

func (c Category) ParseToCategoryRes() CategoryRes {
	return CategoryRes{
		Id:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}
