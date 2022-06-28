package database

import (
	"fmt"
	"log"

	"github.com/bveranoc/mu_server/pkg/config"
	"github.com/bveranoc/mu_server/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_HOST = config.GetEnv("DB_HOST")
var DB_USER = config.GetEnv("DB_USER")
var DB_PASSWORD = config.GetEnv("DB_PASSWORD")
var DB_NAME = config.GetEnv("DB_NAME")
var DB_PORT = config.GetEnv("DB_PORT")

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	log.Println("Database connected")
	db.AutoMigrate(&models.User{}, &models.Mini{})

	DB = db
}
