package controller

import (
	"encoding/json"
	"net/http"

	"github.com/GourabDas18/g-rest/internal/model"
	"github.com/GourabDas18/g-rest/internal/service"
	"github.com/GourabDas18/g-rest/utility"
)

func CreateCountry(w http.ResponseWriter, r *http.Request) {

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
