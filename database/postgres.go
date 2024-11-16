package database

import (
	"github.com/Shikha-code36/golang-crud-rest-api-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectingUsersDB() {
	UserDSN := "host=localhost user=postgres password=Root dbname=users_db port=5432 sslmode=disable"
	UserDB, err := gorm.Open(postgres.Open(UserDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := UserDB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
