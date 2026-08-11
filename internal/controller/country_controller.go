package controller

import (
	"encoding/json"
	"net/http"

	"github.com/GourabDas18/g-rest/internal/model"
	"github.com/GourabDas18/g-rest/internal/service"
	"github.com/GourabDas18/g-rest/utility"
)

func CountrySave(w http.ResponseWriter, r *http.Request) {
	var countryReq *model.CountryCreateReq
	err := json.NewDecoder(r.Body).Decode(&countryReq)
	if err != nil {
		utility.Response(w, http.StatusBadRequest, err.Error(), nil, utility.Error)
		return
	}
	var country model.Country
	if countryReq.Id != nil {
		err = service.Db.First(&country).Where("id = ?", *countryReq.Id).Error
		if err != nil {
			utility.Response(w, http.StatusNotFound, "No country found!", nil, utility.Error)
			return
		}
		err = service.Db.Model(&model.Country{}).Where("id = ?", *countryReq.Id).Updates(&country).Error
		if err != nil {
			utility.Response(w, http.StatusInternalServerError, err.Error(), nil, utility.Error)
			return
		}
	} else {
		country = model.CountryParseFromReq(countryReq)
		respErr := service.Db.Create(&country).Error

		if respErr != nil {
			utility.Response(w, http.StatusInternalServerError, respErr.Error(), nil, utility.Error)
			return
		}
	}
	countryResp := model.CountryParseFromRes(country)
	utility.Response(w, http.StatusCreated, "Successfull", &countryResp, utility.Success)
}

func CreateBulkCountry(w http.ResponseWriter, r *http.Request) {
	var countryList []model.CountryCreateReq
	err := json.NewDecoder(r.Body).Decode(&countryList)
	if err != nil {
		utility.Response(w, http.StatusBadRequest, err.Error(), nil, utility.Error)
		return
	}

	countryModels := model.BulkCountryParseFromReq(countryList)

	result := service.Db.Create(&countryModels)

	if result.Error != nil {
		utility.Response(w, http.StatusInternalServerError, result.Error.Error(), nil, utility.Error)
		return
	}
	countryListResult := model.BulkCountryParseFromRes(countryModels)
	utility.Response(w, http.StatusCreated, "Created Successfuly", &countryListResult, utility.Success)
}

func UpdateCountry(w http.ResponseWriter, r *http.Request) {
	var countryReq model.CountryCreateReq

	err := json.NewDecoder(r.Body).Decode(&countryReq)
	if err != nil {
		utility.Response(w, http.StatusBadRequest, err.Error(), nil, utility.Error)
		return
	}

	if countryReq.Id == nil || (countryReq.Id != nil && *countryReq.Id < 1) {
		print(countryReq.Id)
		utility.Response(w, http.StatusBadRequest, "Give a valid Id", nil, utility.Error)
		return
	}

	country := model.CountryParseFromReq(&countryReq)

	err = service.Db.UpdateColumns(&country).Error
	if err != nil {
		utility.Response(w, http.StatusInternalServerError, err.Error(), nil, utility.Error)
		return
	}
	countryData := model.CountryParseFromRes(country)
	utility.Response(w, http.StatusAccepted, "Updated Successfuly", &countryData, utility.Success)
}

func GetCountryList(w http.ResponseWriter, r *http.Request) {
	var countryList []model.Country

	respError := service.Db.Find(&countryList).Where("isActive = ?", true).Error
	if respError != nil {
		utility.Response(w, http.StatusInternalServerError, respError.Error(), nil, utility.Error)
		return
	}
	countryData := model.BulkCountryParseFromRes(countryList)
	utility.Response(w, http.StatusFound, "Fetched Successfuly", &countryData, utility.Success)
}
