package database

import (
	"fmt"
	"miniProject/config"
	"miniProject/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dbConfig, err := config.LoadDatabase()
	if err != nil {
		return fmt.Errorf("failed to load database config : %v", err)
	}

	dsn := dbConfig.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database : %v", err)
	}

	DB = db

	if err := migrate(); err != nil {
		return fmt.Errorf("failed to migrate database : %v", err)
	}

	return nil
}

func migrate() error {
	if err := DB.AutoMigrate(&model.Employee{}); err != nil {
		return err
	}

	return nil
}