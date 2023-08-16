package req

import "bytes"

// 注册登录请求体
type DouYinUserRegLogRequest struct {
	Username string `form:"username"` // 注册用户名，最长32个字符
	Password string `form:"password"` // 密码，最长32个字符
}

// 用户信息请求体
type DouYinUserRequest struct {
	UserId int64  `form:"user_id"` // 用户id
	Token  string `form:"token"`   // 用户鉴权token
}

//视频投稿请求体

type DouYinPublishActionRequest struct {
	Token string       `form:"token"` // 用户鉴权token
	Data  bytes.Buffer `form:"data"`  // 视频数据
	Title string       `form:"title"` // 视频标题
}

// feed流请求体
type FeedRequest struct {
	LatestTime int64  `form:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `form:"token"`       // 可选参数，登录用户设置
}

// 点赞请求体
type FavoriteActionRequest struct {
	Token      string `form:"token"`       // 用户鉴权token
	VideoId    int64  `form:"video_id"`    // 视频id
	ActionType int32  `form:"action_type"` // 1-点赞，2-取消点赞
}

type CommentActionRequest struct {
	Token       string `form:"token"`        // 用户鉴权token
	VideoId     int64  `form:"vide_id"`      // 视频id
	ActionType  int32  `form:"action_type"`  // 1-发布评论，2-删除评论
	CommentText string `form:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `form:"comment_id"`   // 要删除的评论id，在action_type=2的时候使用
}
