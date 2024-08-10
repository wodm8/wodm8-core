package wod

import (
	"net/http"

	"github.com/wodm8/wodm8-core/internal/application"
	"github.com/wodm8/wodm8-core/internal/domain"

	"github.com/gin-gonic/gin"
)

func CreateWodHandler(wodService application.WodService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user")
		userStruct, ok := user.(domain.UserContext)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		var req domain.CreateWodRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err := wodService.CreateWod(ctx, req, userStruct.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
