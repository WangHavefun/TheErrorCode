package resp

type DouYinUserRegLogResponse struct {
	StatusCode int32  `json:"statusCode"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"statusMsg"`  // 返回状态描述
	UserId     int64  `json:"userId"`     // 用户id
	Token      string `json:"token"`      // 用户鉴权token
}
