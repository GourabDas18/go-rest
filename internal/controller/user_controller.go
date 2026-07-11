package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	config "github.com/GourabDas18/g-rest/internal"
	"github.com/GourabDas18/g-rest/internal/model"
	"github.com/GourabDas18/g-rest/internal/service"
	"github.com/go-playground/validator/v10"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userReq model.UserCreateReq

	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		http.Error(w, "Invalid JSON payload : "+err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(userReq)

	if err != nil {

		validationError, ok := err.(validator.ValidationErrors)

		if ok {
			for _, e := range validationError {
				e.Translate(config.Trans)
			}
		}

		http.Error(w, strings.Join(strings.Split(err.Error(), "\n"), ","), http.StatusBadRequest)
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
