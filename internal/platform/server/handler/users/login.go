package users

import (
	"github.com/gin-gonic/gin"
	"github.com/wodm8/wodm8-core/internal/application"
	"github.com/wodm8/wodm8-core/internal/domain"
	"net/http"
)

func UserLoginHandler(service application.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.LoginRequest

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		tokenString, err := service.Login(req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}
