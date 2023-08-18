package vo

type MessageVo struct {
	Id         int64  `json:"id"`           // 消息id
	ToUserId   int64  `json:"to_user_id"`   // 该消息接收者的id
	FromUserId int64  `json:"from_user_id"` // 该消息发送者的id
	Content    string `json:"content"`      // 消息内容
	CreateTime int64  `json:"create_time"`  // 消息创建时间
}
