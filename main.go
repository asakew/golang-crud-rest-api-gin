package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Shikha-code36/golang-crud-rest-api-gin/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/users", handlers.GetUsers)
	r.GET("/user/:id", handlers.GetUser)
	r.POST("/create_user", handlers.CreateUser)
	r.PUT("/update_user/:id", handlers.UpdateUser)
	r.DELETE("/delete_user/:id", handlers.DeleteUser)

	r.Run()
}
