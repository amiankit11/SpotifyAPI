package database

import (
	"fmt"
	"ltitest/src/database/domain"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
		return
	}

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Failed to connect to the database", err)
		return
	}

	DB = db
	DB.AutoMigrate(&domain.Track{}, &domain.Artist{})
}
