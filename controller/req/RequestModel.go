package req

type DouYinUserRegLogRequest struct {
	Username string `form:"username"` // 注册用户名，最长32个字符
	Password string `form:"password"` // 密码，最长32个字符
}
type DouYinUserRequest struct {
	UserId int64  `form:"user_id"` // 用户id
	Token  string `form:"token"`   // 用户鉴权token
}
