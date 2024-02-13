package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {

	fmt.Println("Middleware Authentication")
	c.Next()

}
