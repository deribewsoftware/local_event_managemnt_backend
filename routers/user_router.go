package routers

import (
	"github.com/deribewsoftware/event_managemnt/controllers"
	"github.com/deribewsoftware/event_managemnt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.POST("/signup", controllers.Signup)

	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.Authentication, controllers.ValidateUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/user/:id", controllers.GetUser)

}
