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

type CategoryController struct{}

func (c CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryReq model.CategoryReq
	if err := json.NewDecoder(r.Body).Decode(&categoryReq); err != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	pass, err := utility.ValidatorG(&categoryReq)
	if !pass {
		utility.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	category := categoryReq.ParseToCategory()
	var count int64
	er := service.Db.Find(&model.Category{}, "user_id = ? AND name = ?", categoryReq.UserId, categoryReq.Name).Count(&count).Error
	if er != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, er.Error())
		return
	}
	if count > 0 {
		utility.ErrorResponse(w, http.StatusBadRequest, "Already exist!")
		return
	}
	if er := service.Db.Create(&category).Error; er != nil {
		utility.ErrorResponse(w, http.StatusBadRequest, er.Error())
		return
	}
	categoryResp := category.ParseToCategoryRes()
	utility.SuccessResponse(w, http.StatusOK, "Successful", &categoryResp)
}

func (c CategoryController) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		var category model.Category
		if err := service.Db.First(&category, "id = ?", id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				utility.ErrorResponse(w, http.StatusBadRequest, "No category found")
				return
			}
			utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utility.SuccessResponse(w, http.StatusOK, "Successful", category.ParseToCategoryRes())
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
	var categorys []model.Category
	if err := service.Db.Find(&categorys, "user_id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utility.ErrorResponse(w, http.StatusBadRequest, "No category found")
			return
		}
		utility.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	accoutResp := []model.CategoryRes{}
	for _, v := range categorys {
		accoutResp = append(accoutResp, v.ParseToCategoryRes())
	}
	utility.SuccessResponse(w, http.StatusOK, "Successful", &accoutResp)

}
