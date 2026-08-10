package controller

import (
	"encoding/json"
	"net/http"

	"github.com/GourabDas18/g-rest/internal/model"
	"github.com/GourabDas18/g-rest/internal/service"

	"github.com/GourabDas18/g-rest/utility"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userReq model.UserCreateReq

	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		http.Error(w, "Invalid JSON payload : "+err.Error(), http.StatusBadRequest)
		return
	}
	pass, errMessage := utility.ValidatorG(userReq)

	if !pass {
		utility.Response(w, http.StatusBadRequest, errMessage, nil, utility.Error)
		return
	}

	var user model.User

	user = model.UserParseFromRequest(&userReq)

	err = service.Db.Save(&user).Error

	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	userResp := model.UserResponseParser(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(userResp)

}
