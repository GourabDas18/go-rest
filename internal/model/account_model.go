package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	Name          string `gorm:"not null;size=100" json:"name"`
	AccountNumber string `json:"accountNumber"`
	UserId        uint   `gorm:"not null;index" json:"userId"`
	Balance       int64  `gorm:"default=0" json:"balance"`
	IsActive      *bool  `gorm:"default=true" json:"isActive"`
}

type AccountReq struct {
	Name          string `json:"name" validate:"required"`
	AccountNumber string `json:"accountNumber"`
	UserId        uint   `json:"userId" validate:"required"`
	Balance       int64  `json:"balance" validate:"required"`
	IsActive      *bool  `json:"isActive"`
}

type AccountRes struct {
	Id            uint
	Name          string `json:"name" validate:"required"`
	AccountNumber string `json:"accountNumber"`
	UserId        uint   `json:"userId" validate:"required"`
	Balance       int64  `json:"balance" validate:"required"`
}

func (acc Account) ParseAccountRes() AccountRes {
	return AccountRes{
		Id:            acc.ID,
		Name:          acc.Name,
		AccountNumber: acc.AccountNumber,
		UserId:        acc.UserId,
		Balance:       acc.Balance,
	}
}

func (acc AccountReq) ParseToAccount() Account {
	return Account{
		Name:          acc.Name,
		AccountNumber: acc.AccountNumber,
		UserId:        acc.UserId,
		Balance:       acc.Balance,
		IsActive:      acc.IsActive,
	}
}
