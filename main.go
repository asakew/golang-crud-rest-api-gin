package main

import (
	"github.com/Shikha-code36/golang-crud-rest-api-gin/database"
	"github.com/Shikha-code36/golang-crud-rest-api-gin/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	database.ConnectingUsersDB()

	r.GET("/users", handlers.GetUsers)
	r.GET("/user/:id", handlers.GetUser)
	r.POST("/create_user", handlers.CreateUser)
	r.PUT("/update_user/:id", handlers.UpdateUser)
	r.DELETE("/delete_user/:id", handlers.DeleteUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
