package database

import (
	"log"

	"github.com/bveranoc/mu_server/pkg/config"
	"github.com/bveranoc/mu_server/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := config.GetEnv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	log.Println("Database connected")
	db.AutoMigrate(&models.User{}, &models.Mini{})

	DB = db
}
