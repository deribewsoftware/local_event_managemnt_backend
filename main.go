package main

import (
	"log"
	"os"

	database "github.com/deribewsoftware/event_managemnt/Database"
	"github.com/deribewsoftware/event_managemnt/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	database.ConnectedToDatabase()
	database.SyncDatabase()
}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.POST("/signup", controllers.Signup)
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "your api is working successfully",
		})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}
