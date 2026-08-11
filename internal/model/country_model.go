package model

import (
	"strings"

	"gorm.io/gorm"
)

type Country struct {
	gorm.Model

	Name     string `gorm:"not null;size=200"`
	Currency string `gorm:"not null"`
	IsActive bool   `gorm:"default:true"`
}

type CountryCreateReq struct {
	Id       *uint  `json:"id"`
	Name     string `json:"name" validate:"required"`
	Currency string `json:"currency" validate:"required"`
	IsActive *bool  `json:"isActive"`
}
type CountryCreateRes struct {
	Id       uint   `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Currency string `json:"currency" validate:"required"`
	IsActive bool   `json:"isActive"`
}

func CountryParseFromReq(countryReq *CountryCreateReq) Country {
	return Country{
		Name:     strings.TrimSpace(countryReq.Name),
		Currency: strings.TrimSpace(countryReq.Currency),
		IsActive: *countryReq.IsActive,
	}
}

func BulkCountryParseFromReq(countryList []CountryCreateReq) []Country {
	countryListData := make([]Country, len(countryList))
	for i, c := range countryList {
		countryListData[i] = Country{
			Name:     strings.TrimSpace(c.Name),
			Currency: strings.TrimSpace(c.Currency),
			IsActive: *c.IsActive,
		}
		// if c.IsActive == nil {
		// 	trueValue := true
		// 	countryListData[i].IsActive = &trueValue
		// } else {
		// 	countryListData[i].IsActive = c.IsActive
		// }
	}
	return countryListData
}

func CountryParseFromRes(c Country) CountryCreateRes {
	return CountryCreateRes{
		Name:     strings.TrimSpace(c.Name),
		Currency: strings.TrimSpace(c.Currency),
		IsActive: c.IsActive,
	}
}

func BulkCountryParseFromRes(countryList []Country) []CountryCreateRes {
	countryListData := make([]CountryCreateRes, len(countryList))
	for i, c := range countryList {
		countryListData[i] = CountryCreateRes{
			Name:     strings.TrimSpace(c.Name),
			Currency: strings.TrimSpace(c.Currency),
			IsActive: c.IsActive,
		}
	}
	return countryListData
}
