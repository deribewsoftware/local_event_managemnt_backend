package database

import (
	"github.com/deribewsoftware/event_managemnt/models"
)

func SyncDatabase() {

	DB.AutoMigrate(&models.User{})
}
