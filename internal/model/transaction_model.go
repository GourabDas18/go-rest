package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model

	Amount      float64 `gorm:"not null;default=0"`
	Description string  `gorm:"default=''"`
	UserId      uint    `gorm:"not null;index"`
	CategoryId  uint    `gorm:"not null;index"`
	Type        uint16  `gorm:"not null;index"`
	AccountId   uint    `gorm:"not null;index"`
}

type TransactionResult struct {
	Transaction
	AccountName string `json:"accountName"`
	UserName    string `json:"userName"`
}

type TransactionReq struct {
	Amount      float64 `json:"amount" validate:"gt=0"`
	Description string  `json:"description" `
	UserId      uint    `json:"userId" validate:"required"`
	CategoryId  *uint   `json:"categoryId" validate:"required"`
	Type        uint16  `json:"type" validate:"required"`
	AccountId   *uint   `json:"accountId" validate:"required"`
}
type TransactionRes struct {
	Id          uint    `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description" `
	CategoryId  uint    `json:"categoryId"`
	Type        uint16  `json:"type"`
	AccountId   uint    `json:"accountId"`
	AccountName string  `json:"accountName"`
	UserName    string  `json:"userName"`
	CreatedAt   string  `json:"createdAt"`
}

func (t TransactionResult) ParseToTransactionRes() TransactionRes {
	return TransactionRes{
		Id:          t.ID,
		Amount:      t.Amount,
		Description: t.Description,
		CategoryId:  t.CategoryId,
		Type:        t.Type,
		AccountId:   t.AccountId,
		AccountName: t.AccountName,
		UserName:    t.UserName,
		CreatedAt:   t.CreatedAt.Format(time.DateTime),
	}
}

func (t TransactionReq) ParseToTransaction() Transaction {
	return Transaction{
		Amount:      t.Amount,
		Description: t.Description,
		CategoryId:  *t.AccountId,
		Type:        t.Type,
		AccountId:   *t.AccountId,
		UserId:      t.UserId,
	}
}
