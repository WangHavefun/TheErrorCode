package resp

import "TheErrorCode/vo"

type DouYinUserRegLogResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	UserId     int64  `json:"userId"`      // 用户id
	Token      string `json:"token"`       // 用户鉴权token
}
type UserResp struct {
	StatusCode int32     `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string    `json:"status_msg"`  // 返回状态描述
	User       vo.UserVo `json:"user"`        // 数据
}
type VideoListResp struct {
	StatusCode int32        `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string       `json:"status_msg"`  // 返回状态描述
	VideoList  []vo.VideoVo `json:"video_list"`  // 用户发布的视频列表
}
type FeedResponse struct {
	StatusCode int32         `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `json:"status_msg"`  // 返回状态描述
	VideoList  *[]vo.VideoVo `json:"video_list"`  // 视频列表
	NextTime   int64         `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type CommentActionResponse struct {
	StatusCode int32        `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string       `json:"status_msg"`  // 返回状态描述
	Comment    vo.CommentVo `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
}

type CommentListResponse struct {
	StatusCode  int32           `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string          `json:"status_msg"`   // 返回状态描述
	CommentList *[]vo.CommentVo `json:"comment_list"` // 评论列表
}
type RelationFollowListResponse struct {
	StatusCode int32        `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string       `json:"status_msg"`  // 返回状态描述
	UserList   *[]vo.UserVo `json:"user_list"`   // 用户信息列表
}
type RelationFriendListResponse struct {
	StatusCode int32              `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string             `json:"status_msg"`  // 返回状态描述
	UserList   *[]vo.FriendUserVo `json:"user_list"`   // 用户信息列表
}
