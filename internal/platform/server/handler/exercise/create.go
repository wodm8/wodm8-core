package exercise

import (
	"github.com/wodm8/wodm8-core/internal/creating"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func CreateHandler(exerciseService creating.ExerciseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := exerciseService.CreateExercise(ctx, req.ID, req.Name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		//exercise, err := crossfit.NewExercise(req.ID, req.Name)
		//if err != nil {
		//	ctx.JSON(http.StatusBadRequest, err.Error())
		//	return
		//}
		//if err := exerciseRepository.Save(ctx, exercise); err != nil {
		//	ctx.JSON(http.StatusInternalServerError, err.Error())
		//	return
		//}

		ctx.Status(http.StatusCreated)
	}
}
