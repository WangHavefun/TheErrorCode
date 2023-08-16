package vo

type CommentVo struct {
	Id         int64   `json:"id"`          // 视频评论id
	User       *UserVo `json:"user"`        // 评论用户信息
	Content    string  `json:"content"`     // 评论内容
	CreateDate string  `json:"create_date"` // 评论发布日期，格式 mm-dd
}
