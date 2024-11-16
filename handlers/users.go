package handlers

import (
	"github.com/Shikha-code36/golang-crud-rest-api-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var users = []models.User{
	{ID: 1, Name: "Radha"},
	{ID: 2, Name: "Krishna"},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
}

var lastID = 2 // keep track of last used ID

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastID++         // increment last used ID
	user.ID = lastID // assign new ID to user
	users = append(users, user)
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updateUser models.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, user := range users {
		if user.ID == id {
			updateUser.ID = id // set the id of updateUser
			users[i] = updateUser
			c.JSON(http.StatusOK, updateUser)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
}
