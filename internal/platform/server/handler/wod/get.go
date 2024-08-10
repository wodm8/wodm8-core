package wod

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wodm8/wodm8-core/internal/application"
)

func GetWodHandler(wodService application.WodService) gin.HandlerFunc {
	return func(c *gin.Context) {
		memberId := c.Param("memberid")
		fmt.Printf("Param: %v\n", memberId)
		response, err := wodService.GetWod(c, memberId)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
