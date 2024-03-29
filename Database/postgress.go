package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectedToDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=alpha port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDatabase() (db *gorm.DB) {
	return db
}
