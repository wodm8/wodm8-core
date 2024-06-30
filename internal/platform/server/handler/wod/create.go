package wod

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	crossfit "github.com/wodm8/wodm8-core/internal"
)

type ExerciseInWod struct {
	ID               string  `json:"id" binding:"required"`
	Reps             int32   `json:"reps" binding:"required"`
	Weight           float32 `json:"weight"`
	WeightUnit       string  `json:"weight_unit"`
	Section          int32   `json:"section" binding:"required"`
	SectionTimerType int32   `json:"section_timer_type"`
	SectionCap       int32   `json:"section_cap"`
}

type createRequest struct {
	ID             string          `json:"id" binding:"required"`
	Name           string          `json:"name" binding:"required"`
	Rounds         int32           `json:"rounds" binding:"required"`
	NumberSections int32           `json:"number_sections" binding:"required"`
	TimerType      int32           `json:"timer_type_id" binding:"required"`
	Exercises      []ExerciseInWod `json:"exercises" binding:"required"`
}

func CreateWodHandler(wodRepository crossfit.WodRepository, exerciseWodRepository crossfit.ExerciseWodRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		wod, err := crossfit.NewWod(req.ID, req.Name, req.Rounds, req.NumberSections, req.TimerType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err := wodRepository.Save(ctx, wod); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		for _, exercise := range req.Exercises {
			fmt.Println(exercise)

			id, err := uuid.NewRandom()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}

			exerciseWod, err := crossfit.NewExerciseWod(id.String(), wod.ID().String(), exercise.ID, exercise.Reps, exercise.Weight, exercise.WeightUnit, exercise.Section, exercise.SectionTimerType, exercise.SectionCap)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}

			if err := exerciseWodRepository.Save(ctx, exerciseWod); err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
