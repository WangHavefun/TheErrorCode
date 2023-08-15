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
	StatusCode int32        `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string       `json:"status_msg"`  // 返回状态描述
	VideoList  []vo.VideoVo `json:"video_list"`  // 视频列表
	NextTime   int64        `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}
