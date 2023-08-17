package vo

type FriendUserVo struct {
	*UserVo
	Message string `json:"message"`
	MsgType int64  `json:"msgType"`
}
