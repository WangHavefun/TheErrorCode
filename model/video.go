package model

// 视频
type Video struct {
	ID       int64  `gorm:"column:id;" `                //主键
	PlayUrl  string `gorm:"column:play_url;NOT NULL" `  //播放地址
	CoverUrl string `gorm:"column:cover_url;NOT NULL" ` //封面地址
	Title    string `gorm:"column:title;NOT NULL" `     //视频标题
}

func (v *Video) TableName() string {
	return "video"
}
