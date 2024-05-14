package middleware

import (
	"TARG_revenue_report_backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		token = token[7:]

		fmt.Println(token)
		if token == "" {
			c.JSON(500, gin.H{
				"code":    1,
				"message": "沒有携带token",
			})
			c.Abort()
			return
		} else {
			claims, err := utils.ParseToken(token)

			if err != nil {
				c.JSON(500, gin.H{
					"code":    1,
					"message": "错误的token" + err.Error(),
				})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(500, gin.H{
					"code":    1,
					"message": "token超时了,请重新登录",
				})
				c.Abort()
				return
			}
		}
		c.JSON(200, gin.H{
			"code":    0,
			"message": "success",
		})
		c.Next()
	}
}
