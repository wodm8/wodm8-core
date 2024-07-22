package kit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wodm8/wodm8-core/initializers"
	"github.com/wodm8/wodm8-core/internal/users"
	"net/http"
	"time"
)

func RequireAuth(c *gin.Context) {
	var cfg = initializers.Cfg

	// Get the cookie off the request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(cfg.JwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token Subject
		var user users.Users
		initializers.DB.First(&user, "id = ?", claims["sub"])

		if user.ID == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		userMap := make(map[string]string)
		userMap["id"] = user.ID
		userMap["email"] = user.Email
		userMap["first_name"] = user.FirstName
		userMap["last_name"] = user.LastName

		// Attach the request
		c.Set("user", userMap)

		//Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
