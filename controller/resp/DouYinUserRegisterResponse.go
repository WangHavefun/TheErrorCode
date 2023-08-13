package resp

type DouYinUserRegLogResponse struct {
	StatusCode int32  `json:"statusCode"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"statusMsg"`  // 返回状态描述
	UserId     int64  `json:"userId"`     // 用户id
	Token      string `json:"token"`      // 用户鉴权token
}
type Resp struct {
	status_code int32       // 状态码，0-成功，其他值-失败
	status_msg  string      // 返回状态描述
	Data        interface{} // 数据
}
type User struct {
	Id              int64  `json:"id"`               // 用户id
	Name            string `json:"name"`             // 用户名称
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar"`           //用户头像
	BackgroundImage string `json:"background_image"` //用户个人页顶部大图
	Signature       string `json:"signature"`        //个人简介
	TotalFavorited  int64  `json:"total_favorited"`  //获赞数量
	WorkCount       int64  `json:"work_count"`       //作品数量
	FavoriteCount   int64  `json:"favorite_count"`   //点赞数量
}
