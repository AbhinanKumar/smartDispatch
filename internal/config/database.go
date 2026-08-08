package config

import (
	"github.com/AbhinanKumar/smart-dispatch/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func NewDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Customer{},
	)
	if err != nil {
		return nil, err
	}

	log.Println("Database Connected")
	return db, nil
}
