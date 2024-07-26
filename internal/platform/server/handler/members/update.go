package members

import (
	"github.com/gin-gonic/gin"
	"github.com/wodm8/wodm8-core/internal/application"
	"github.com/wodm8/wodm8-core/internal/domain"
	"net/http"
)

func MemberUpdateHandler(memberService application.MemberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.MembersRequest

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := memberService.UpdateMember(req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}
