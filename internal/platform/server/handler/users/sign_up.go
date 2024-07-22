package users

import (
	"github.com/gin-gonic/gin"
	"github.com/wodm8/wodm8-core/internal/application"
	"github.com/wodm8/wodm8-core/internal/domain"
	"net/http"
)

func UserSignUpHandler(userService application.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.CreateUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		err := userService.CreateUser(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusCreated)
	}
}
