package db

import (
	"fmt"
	"log"

	"github.com/depermana12/go-notes/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=localhost user=gonotes password=gonotes dbname=gonotes port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	if err := database.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
		log.Fatal("failed to migrate schema to database:", err)
	}

	fmt.Println("db connected and migrated successfully")
}

func GetDB() *gorm.DB {
	return database
}
