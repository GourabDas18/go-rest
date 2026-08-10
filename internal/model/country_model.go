package model

import (
	"strings"

	"gorm.io/gorm"
)

type Country struct {
	gorm.Model

	Name     string `gorm:"not null;size=200"`
	Currency string `gorm:"not null"`
	IsActive bool   `gorm:"default=true"`
}

type CountryCreateReq struct {
	Name     string `json:"name" validate:"required"`
	Currency string `json:"currency" validate:"required"`
	IsActive bool   `json:"isActive" default:"true"`
}
type CountryCreateRes struct {
	Id       uint   `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Currency string `json:"currency" validate:"required"`
	IsActive bool   `json:"isActive" default:"true"`
}

func CountryParseFromReq(countryReq CountryCreateReq) Country {
	return Country{
		Name:     strings.TrimSpace(countryReq.Name),
		Currency: strings.TrimSpace(countryReq.Currency),
		IsActive: countryReq.IsActive,
	}
}

func BulkCountryParseFromReq(countryList []CountryCreateReq) []Country {
	countryListData := make([]Country, len(countryList))
	for i, c := range countryList {
		countryListData[i] = Country{
			Name:     strings.TrimSpace(c.Name),
			Currency: strings.TrimSpace(c.Currency),
			IsActive: c.IsActive,
		}
	}
	return countryListData
}

func BulkCountryParseFromRes(countryList []Country) []CountryCreateRes {
	countryListData := make([]CountryCreateRes, len(countryList))
	for i, c := range countryList {
		countryListData[i] = CountryCreateRes{
			Id:       c.ID,
			Name:     strings.TrimSpace(c.Name),
			Currency: strings.TrimSpace(c.Currency),
			IsActive: c.IsActive,
		}
	}
	return countryListData
}
