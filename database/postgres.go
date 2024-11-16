package database

import (
	"github.com/Shikha-code36/golang-crud-rest-api-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectingUsersDB() {
	dsn := "host=localhost user=your_username password=your_password dbname=users_db port=5432 sslmode=disable"
	UsersDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := UsersDB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
