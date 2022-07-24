package middleware

import (
	"code.project.com/InstantMessaging/pkg/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthLoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		userClaim, err := token.AnalyseToken(tokenString)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户未登录",
			})
			return
		}
		c.Set("user_claim", userClaim)
		c.Next()
	}
}
