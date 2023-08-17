package model

import "gorm.io/gorm"

type Message struct {
	ID        int64  `gorm:"column:id;"`
	PostId    int64  `gorm:"column:post_id;"`    //发送方id
	ReceiveId int64  `gorm:"column:receive_id;"` //接收方id
	Msg       string `gorm:"column:message;"`    //信息
	gorm.Model
}

func (Message) TableName() string {
	return "message"
}
