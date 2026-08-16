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

type TransactionController struct {
}

func (t TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tReq model.TransactionReq
	if err := json.NewDecoder(r.Body).Decode(&tReq); err != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pass, err := utility.ValidatorG(&tReq)
	if !pass {
		utility.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	tData := tReq.ParseToTransaction()
	if err := service.Db.Create(&tData).Error; err != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	var txnResult model.TransactionResult
	errResp := service.Db.
		Table("transactions t").
		Joins("JOIN users u on u.id = t.user_id").
		Joins("JOIN accounts a on a.id= t.account_id").
		Where("t.id = ?", tData.ID).
		Select("t.*,u.name as user_name,a.name as account_name").
		Scan(&txnResult).Error
	if errResp != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, errResp.Error())
		return
	}
	utility.SuccessResponse(w, http.StatusOK, "Successful", txnResult.ParseToTransactionRes())
}

func (t TransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		var txn model.TransactionResult
		if err := service.Db.
			Table("transactions t").
			Joins("JOIN users u on u.id = t.user_id").
			Joins("JOIN accounts a on a.id= t.account_id").
			Where("t.id = ?", id).
			Select("t.*,u.name as user_name,a.name as account_name").
			Scan(&txn).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				utility.ErrorResponse(w, http.StatusBadRequest, "No account found")
				return
			}
			utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utility.SuccessResponse(w, http.StatusOK, "Successful", txn.ParseToTransactionRes())
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
	var txns []model.TransactionResult
	if err := service.Db.
		Table("transactions t").
		Joins("JOIN users u on u.id = t.user_id").
		Joins("JOIN accounts a on a.id= t.account_id").
		Where("t.user_id = ?", userId).
		Select("t.*,u.name as user_name,a.name as account_name").
		Scan(&txns).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utility.ErrorResponse(w, http.StatusBadRequest, "No account found")
			return
		}
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	txnsResp := []model.TransactionRes{}
	for _, v := range txns {
		txnsResp = append(txnsResp, v.ParseToTransactionRes())
	}
	utility.SuccessResponse(w, http.StatusOK, "Successful", &txnsResp)
}
