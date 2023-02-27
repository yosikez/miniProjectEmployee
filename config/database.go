package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Name)
}

func LoadDatabase() (*DatabaseConfig, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error load .env file")
		return nil, err
	}

	dbConfig := &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	return dbConfig, nil
}
