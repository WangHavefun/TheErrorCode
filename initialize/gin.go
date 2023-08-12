package initalize

import (
	"TheErrorCode/controller"
	"github.com/gin-gonic/gin"
)

func Gin() {
	r := gin.Default()
	user := r.Group("/douyin/user")
	{
		user.POST("/register/", controller.UserController{}.Register)
		user.POST("/login/")
	}
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
