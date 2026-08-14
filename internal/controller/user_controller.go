package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GourabDas18/g-rest/internal/model"
	"github.com/GourabDas18/g-rest/internal/service"
	"gorm.io/gorm"

	"github.com/GourabDas18/g-rest/utility"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userReq model.UserCreateReq

	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		utility.Response(w, http.StatusBadRequest, "Invalid JSON payload : "+err.Error(), nil, utility.Error)
		return
	}
	pass, errMessage := utility.ValidatorG(userReq)

	if !pass {
		utility.Response(w, http.StatusBadRequest, errMessage, nil, utility.Error)
		return
	}

	var user model.User

	user = model.UserParseFromRequest(&userReq)

	var dbUser model.User
	err = service.Db.First(&dbUser, "email = ?", user.Email).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			utility.Response(w, http.StatusExpectationFailed, "User already exist!", nil, utility.Error)
			return
		}
		utility.Response(w, http.StatusExpectationFailed, err.Error(), nil, utility.Error)
		return
	}
	rawPassword := service.Password{Value: user.Password}

	hashPassword, error := rawPassword.GenerateHash()

	if error != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, error.Error())
	}
	user.Password = hashPassword
	err = service.Db.Save(&user).Error

	if err != nil {
		utility.Response(w, http.StatusExpectationFailed, err.Error(), nil, utility.Error)
		return
	}

	userResp := model.UserResponseParser(&user)

	authToken, err := service.GetToken(int(userResp.ID), userResp.Name, int(userResp.CountryId))
	if err != nil {
		utility.Response(w, http.StatusUnauthorized, err.Error(), nil, utility.Error)
		return
	}

	userResp.Token = &authToken

	utility.Response(w, http.StatusCreated, "Created Successfuly", &userResp, utility.Success)

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var authCred model.UserAuthReq

	if err := json.NewDecoder(r.Body).Decode(&authCred); err != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	pass, errMessage := utility.ValidatorG(authCred)
	if !pass {
		utility.ErrorResponse(w, http.StatusBadRequest, errMessage)
		return
	}
	var dbUser model.User

	err := service.Db.First(&dbUser, "email = ?", authCred.Email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utility.ErrorResponse(w, http.StatusBadRequest, "No user found!")
			return
		}
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	password := service.Password{Value: dbUser.Password}
	valid := password.ValidateHash([]byte(authCred.Password))
	if !valid {
		utility.ErrorResponse(w, http.StatusBadRequest, "Wrong password")
		return
	}
	token, err := service.GetToken(int(dbUser.ID), dbUser.Name, int(dbUser.CountryId))
	if err != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	userResp := model.UserResponseParser(&dbUser)
	userResp.Token = &token
	utility.SuccessResponse(w, http.StatusOK, "Successful", &userResp)

}
