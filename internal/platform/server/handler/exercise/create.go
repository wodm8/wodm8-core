package exercise

import (
	"net/http"

	"github.com/wodm8/wodm8-core/internal/application"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Name string `json:"name" binding:"required"`
}

func CreateHandler(exerciseService application.ExerciseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := exerciseService.CreateExercise(ctx, req.Name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.Status(http.StatusCreated)
	}
}
