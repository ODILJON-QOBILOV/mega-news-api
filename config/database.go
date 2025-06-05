package config

import (
	"log"
	"os"

	"github.com/newsapi/v2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the database:", err)
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.News{})
	database.AutoMigrate(&models.Category{})
	database.AutoMigrate(&models.Follow{})
	database.AutoMigrate(&models.Banner{})
	database.AutoMigrate(&models.BannerCarousel{})
	database.AutoMigrate(&models.Advertisement{})

	DB = database
}
