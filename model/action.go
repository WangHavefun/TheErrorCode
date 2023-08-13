package model

type Action struct {
	ID       int64 `gorm:"column:id;"`       //主键
	VideoId  int64 `gorm:"column:video_id;"` //视频id
	Favorite bool  `gorm:"column:favorite;"` //是否点赞
}

func (a *Action) TableName() string {
	return "action"
}
