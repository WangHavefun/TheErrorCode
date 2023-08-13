package middleware

import (
	"TheErrorCode/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// 验证jwt
func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, flag := c.Params.Get("token")
		if !flag {
			c.Abort()
		}
		_, err := utils.ValidateJWT(token)
		if err != nil {
			log.Println(err)
			c.Abort()
		}
		c.Next()
	}
}
