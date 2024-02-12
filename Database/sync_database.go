package database

import (
	"github.com/deribewsoftware/event_managemnt/models"
)

func SyncDatabase() {

	db.AutoMigrate(&models.User{})
}
