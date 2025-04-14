package database

import (
	"log"
	"user_service/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
	log.Printf("Connecting to PostgreSQL at %s", dsn)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	// Migrate the database schema
	err := db.AutoMigrate(
		&domain.User{},
	)
	if err != nil {
		return err
	}

	return nil
}