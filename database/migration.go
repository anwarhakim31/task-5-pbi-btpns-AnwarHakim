package database

import (
	"log"
	"pbi-btpns-hakim/models"
)

func Migrate() {
	Instance.AutoMigrate(&models.User{}, &models.Photo{})
	log.Println("Database Migration Completed!")
}