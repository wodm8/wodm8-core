package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		c.JSON(http.StatusOK, gin.H{
			"message": user,
		})
	}
}
