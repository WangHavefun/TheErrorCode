package initalize

import (
	"TheErrorCode/controller"
	"github.com/gin-gonic/gin"
)

func Gin() {
	r := gin.Default()
	user := r.Group("/douyin/user/")
	{
		user.POST("register/", controller.UserController{}.Register) //注册
		user.POST("login/", controller.UserController{}.Login)       //登录
		user.GET("", controller.UserController{}.GetUser)            //获取用户信息
	}
	publish := r.Group("/douyin/publish/")
	{

		publish.POST("action/", controller.VideoController{}.PublishVideo) //发布视频
		publish.GET("list/", controller.VideoController{}.VideoList)       //发布的视频列表
	}
	feed := r.Group("/douyin/feed/")
	{
		feed.GET("", controller.FeedController{}.Feed) //视频流
	}
	favorite := r.Group("/douyin/favorite/")
	{
		favorite.POST("action/", controller.FavoriteController{}.Action) //点赞
		favorite.GET("list/", controller.FavoriteController{}.List)      //点赞的作品列表
	}
	comment := r.Group("/douyin/comment/")
	{
		comment.POST("action/", controller.CommentController{}.Comment) //评论
		comment.GET("list/", controller.CommentController{}.List)       //评论列表
	}
	relation := r.Group("/douyin/relation/")
	{
		relation.POST("action/", controller.FollowController{}.Follow)         //关注
		relation.GET("follow/list/", controller.FollowController{}.FollowList) //用户关注列表
		relation.GET("follower/list/", controller.FollowController{}.FanList)  //用户粉丝列表
		relation.GET("friend/list/", controller.FollowController{}.FriendList) //用户好友列表
	}

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
