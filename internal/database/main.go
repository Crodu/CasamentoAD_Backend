package database

import (
	"log"

	"github.com/Crodu/CasamentoBackend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDBConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// AutoMigrate all models from the models package
	err = db.AutoMigrate(
		&models.User{},
		&models.Guest{},
		&models.Gift{},
		&models.BoughtGift{},
		&models.Payment{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
