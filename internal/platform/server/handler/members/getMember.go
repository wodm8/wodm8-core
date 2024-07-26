package members

import (
	"github.com/gin-gonic/gin"
	"github.com/wodm8/wodm8-core/internal/application"
	"github.com/wodm8/wodm8-core/internal/domain"
	"net/http"
)

func GetMemberHandler(memberService application.MemberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		userStruct, ok := user.(domain.UserContext)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		member, err := memberService.GetMember(userStruct.Email)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		c.JSON(http.StatusOK, member)
	}
}
