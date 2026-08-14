package controller

import (
	"encoding/json"
	"net/http"

	"github.com/GourabDas18/g-rest/internal/model"
	"github.com/GourabDas18/g-rest/internal/service"
	"github.com/GourabDas18/g-rest/utility"
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
	err := service.Db.Find(model.Account{}, "userId = ? AND name = ?", account.UserId, account.Name).Count(&accountNo).Error
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
