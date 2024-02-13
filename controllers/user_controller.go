package controllers

import (
	"net/http"

	database "github.com/deribewsoftware/event_managemnt/Database"
	"github.com/deribewsoftware/event_managemnt/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {

	var body struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	ctx.BindJSON(&body)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "password is  not hashable",
		})
		return
	}
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := database.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user is not created",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "users successfully created",
		"user":    user,
	})
}

func Login(ctx *gin.Context) {

	var body struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	ctx.BindJSON(&body)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "password is  not hashable",
		})
		return
	}
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := database.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user is not created",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "users successfully created",
		"user":    user,
	})
}

func GetAllUsers(ctx *gin.Context) {
	var users []*models.User
	database.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "users successfully retrieved",
		"users":   users,
	})

}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := database.DB.Where("id=?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"user":    user,
		})
	}
}
