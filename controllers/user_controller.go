package controllers

import (
	"net/http"

	database "github.com/deribewsoftware/event_managemnt/Database"
	"github.com/deribewsoftware/event_managemnt/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	//get email name, and password from body

	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Body of user",
		})

		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}
	// create user
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := database.CreateDatabase().Create(user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Body of user",
		})

		return
	}

	// get response
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
