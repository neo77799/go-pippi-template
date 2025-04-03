package config

import (
	"fmt"
	"log"
	"os"

	"github.com/shibu/omikuji-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.OmikujiResult{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Failed to get database connection: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("failed to close database connection: %v", err)
	}
}
