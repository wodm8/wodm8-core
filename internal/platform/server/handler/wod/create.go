package wod

import (
	"github.com/wodm8/wodm8-core/internal/creating"
	"github.com/wodm8/wodm8-core/internal/platform/server/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWodHandler(wodService creating.WodService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req handler.CreateWodRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err := wodService.CreateWod(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		//wod, err := crossfit.NewWod(req.ID, req.Name, req.Rounds, req.NumberSections, req.TimerType)
		//if err != nil {
		//	ctx.JSON(http.StatusBadRequest, err)
		//	return
		//}
		//
		//if err := wodRepository.Save(ctx, wod); err != nil {
		//	ctx.JSON(http.StatusInternalServerError, err.Error())
		//	return
		//}
		//
		//for _, exercise := range req.Exercises {
		//	fmt.Println(exercise)
		//
		//	id, err := uuid.NewRandom()
		//	if err != nil {
		//		ctx.JSON(http.StatusInternalServerError, err)
		//		return
		//	}
		//
		//	exerciseWod, err := crossfit.NewExerciseWod(id.String(), wod.ID().String(), exercise.ID, exercise.Reps, exercise.Weight, exercise.WeightUnit, exercise.Section, exercise.SectionTimerType, exercise.SectionCap)
		//	if err != nil {
		//		ctx.JSON(http.StatusBadRequest, err.Error())
		//		return
		//	}
		//
		//	if err := exerciseWodRepository.Save(ctx, exerciseWod); err != nil {
		//		ctx.JSON(http.StatusInternalServerError, err)
		//		return
		//	}
		//}

		ctx.Status(http.StatusCreated)
	}
}
