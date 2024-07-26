package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wodm8/wodm8-core/internal/application"
	"github.com/wodm8/wodm8-core/internal/domain"
	"net/http"
	"time"
)

func UserSignUpHandler(userService application.UsersService, memberService application.MemberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tmNow := time.Now()

		var req domain.CreateUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		cu := make(chan any, 1)
		cm := make(chan any, 1)

		go func() {
			cu <- userService.CreateUser(req)
			close(cu)
		}()

		go func() {
			cm <- memberService.CreateMember(req)
			close(cm)
		}()

		resCu := <-cu
		resCm := <-cm

		if resCu != nil {
			fmt.Printf("***this is a go rutine 1 validation\n")
			c.JSON(http.StatusInternalServerError, resCu)
		}

		if resCm != nil {
			fmt.Printf("***this is a go rutine 2 validation\n")
			c.JSON(http.StatusInternalServerError, resCm)
		}

		//if err := userService.CreateUser(req); err != nil {
		//	fmt.Printf("error in service member %v\n", err)
		//	c.JSON(http.StatusInternalServerError, err.Error())
		//}
		//
		//if err := memberService.CreateMember(req); err != nil {
		//	fmt.Printf("error in service member %v\n", err)
		//	c.JSON(http.StatusInternalServerError, err.Error())
		//}

		fmt.Println(time.Since(tmNow))

		c.Status(http.StatusCreated)
	}
}
