package service

import (
	"log"

	"github.com/GourabDas18/g-rest/internal/model"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Account{},
		&model.Category{},
		&model.Transaction{},
		&model.User{},
		&model.Country{},
	)

	if err != nil {
		log.Fatalf("Error in table migration : %s", err.Error())
	}
}
