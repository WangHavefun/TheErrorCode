package model

import "gorm.io/gorm"

type Action struct {
	ID       int64 `gorm:"column:id;"`        //主键
	UserId   int64 `gorm:"column:user_id;"`   //用户id
	AuthorId int64 `gorm:"column:author_id;"` //作者id（冗余）
	VideoId  int64 `gorm:"column:video_id;"`  //视频id
	Favorite bool  `gorm:"column:favorite;"`  //是否点赞
	Love     bool  `gorm:"column:love;"`      //是否喜欢
	gorm.Model
}

func (a *Action) TableName() string {
	return "action"
}
