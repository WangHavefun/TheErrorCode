package model

import "gorm.io/gorm"

// 评论
type Comment struct {
	ID      int64  `gorm:"column:id;"`             //主键
	VideoId int64  `gorm:"column:video_id;"`       //视频id
	Text    string `gorm:"column:text;;type:text"` //内容
	gorm.Model
}

func (c *Comment) TableName() string {
	return "comment"
}
