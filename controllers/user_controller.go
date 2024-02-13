package controllers

import (
	"net/http"
	"os"
	"time"

	database "github.com/deribewsoftware/event_managemnt/Database"
	"github.com/deribewsoftware/event_managemnt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	// get e,ail and  password from body

	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Failed to login",
		})
	}

	// lookup the user
	var user models.User

	database.DB.First(&user, "email==?", body.Email)
	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid email",
		})
		return

	}

	// compare hash password

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Password is mismatch",
		})
		return
	}

	// Generate Jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and Get the complete  encoded payload  as String wit secrets

	tokenString, err := token.SignedString(os.Getenv("SECRET"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Fail to create token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   tokenString,
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
