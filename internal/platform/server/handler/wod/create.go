package wod

import (
	"net/http"

	"github.com/gin-gonic/gin"
	crossfit "github.com/wodm8/wodm8-core/internal"
)

type createRequest struct {
	ID             string `json:"id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Rounds         int32  `json:"rounds" binding:"required"`
	NumberSections int32  `json:"number_sections" binding:"required"`
	TimerType      int32  `json:"timer_type_id" binding:"required"`
}

func CreateWodHandler(wodRepository crossfit.WodRepository) gin.HandlerFunc {
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
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
