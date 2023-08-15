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
		user.POST("/login/", controller.UserController{}.Login)
		user.GET("/", controller.UserController{}.GetUser)
	}
	publish := r.Group("/douyin/publish")
	{
		publish.POST("/action/", controller.VideoController{}.PublishVideo)
		publish.GET("/list/", controller.VideoController{}.VideoList)
	}
	feed := r.Group("/douyin/feed/")
	{
		feed.GET("", controller.FeedController{}.Feed)
	}
	favorite := r.Group("/douyin/favorite/")
	{
		favorite.POST("action/", controller.FavoriteController{}.Action)
		favorite.GET("list/", controller.FavoriteController{}.List)
	}
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
