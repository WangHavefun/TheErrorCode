package model

import "gorm.io/gorm"

// 关注
type Follow struct {
	ID       int64 `gorm:"column:id;"`                 //主键
	FollowId int64 `gorm:"column:follow_id;NOT NULL;"` //关注的ID
	FanId    int64 `gorm:"column:fan_id;NOT NULL;"`    //粉丝的ID
	gorm.Model
}

func (f *Follow) TableName() string {
	return "follow"
}
