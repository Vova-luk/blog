package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connecting to PostgreSQL
func Connect() (*gorm.DB, error) {

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST_POSTGRES"), os.Getenv("DB_PORT_POSTGRES"), os.Getenv("DB_USER_POSTGRES"), os.Getenv("DB_PASSWORD_POSTGRES"),
		os.Getenv("DB_NAME_POSTGRES"), os.Getenv("DB_SSLMODE_POSTGRES"))

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Printf("Connection to datebase")
	return db, nil
}
