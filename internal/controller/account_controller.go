package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GourabDas18/g-rest/internal/model"
	"github.com/GourabDas18/g-rest/internal/service"
	"github.com/GourabDas18/g-rest/utility"
	"gorm.io/gorm"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var accountReq model.AccountReq
	if err := json.NewDecoder(r.Body).Decode(&accountReq); err != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	pass, errStr := utility.ValidatorG(&accountReq)
	if !pass {
		utility.ErrorResponse(w, http.StatusBadRequest, errStr)
		return
	}
	account := accountReq.ParseToAccount()
	var accountNo int64
	err := service.Db.Find(model.Account{}, "user_id = ? AND name = ?", account.UserId, account.Name).Count(&accountNo).Error
	if err != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if accountNo > 0 {
		utility.ErrorResponse(w, http.StatusBadRequest, "Already exists!")
		return
	}
	err = service.Db.Create(&account).Error
	if err != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	resp := account.ParseAccountRes()
	utility.SuccessResponse(w, http.StatusCreated, "Succesful", &resp)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id != "" {
		var account model.Account
		if err := service.Db.First(&account, "id = ?", id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				utility.ErrorResponse(w, http.StatusBadRequest, "No account found")
				return
			}
			utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utility.SuccessResponse(w, http.StatusOK, "Successful", account.ParseAccountRes())
		return
	}
	body := map[string]any{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	userId := body["userId"]
	if userId == nil {
		utility.ErrorResponse(w, http.StatusBadRequest, "User Id not found")
		return
	}
	var accounts []model.Account
	if err := service.Db.Find(&accounts, "user_id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utility.ErrorResponse(w, http.StatusBadRequest, "No account found")
			return
		}
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	accoutResp := []model.AccountRes{}
	for _, v := range accounts {
		accoutResp = append(accoutResp, v.ParseAccountRes())
	}
	utility.SuccessResponse(w, http.StatusOK, "Successful", &accoutResp)
}
