package model

import "gorm.io/gorm"

// 视频
type Video struct {
	ID       int64  `gorm:"column:id;" `                                  //主键
	UserId   int64  `gorm:"column:user_id;" `                             //作者id
	PlayUrl  string `gorm:"column:play_url;NOT NULL;type:varchar(255)" `  //播放地址
	CoverUrl string `gorm:"column:cover_url;NOT NULL;type:varchar(255)" ` //封面地址
	Title    string `gorm:"column:title;NOT NULL;type:varchar(255)" `     //视频标题
	gorm.Model
}

func (v *Video) TableName() string {
	return "video"
}
