package boxes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wodm8/wodm8-core/internal/application"
	"github.com/wodm8/wodm8-core/internal/domain"
)

func CreateBoxHandler(boxService application.BoxService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user")
		userStruct, ok := user.(domain.UserContext)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		var req domain.CreateBoxRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err := boxService.CreateBox(ctx, req, userStruct.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
