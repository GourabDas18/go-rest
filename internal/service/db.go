package service

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DBInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	apiSecret := os.Getenv("API_SECRET")
	dsn := fmt.Sprintf(`host=%s user=postgres password=%s dbname=buzzet port=%s sslmode=disable TimeZone=Asia/Kolkata`, dbHost, apiSecret, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error in DB connection %s", err.Error())
	}
	log.Println("Db is connected")
	Db = db
}
