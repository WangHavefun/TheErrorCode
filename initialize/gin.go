package initalize

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Gin() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, "111")
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
