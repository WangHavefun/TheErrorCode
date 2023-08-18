package dao

import (
	"TheErrorCode/constant"
	"TheErrorCode/model"
	"time"
)

type MessageDao struct {
}

func (d MessageDao) GetLatestByPostIdAndReceiveId(postId int64, receiveId int64) *model.Message {
	var message *model.Message
	constant.DB.Model(message).Where("(post_id = ? and receive_id = ?)or(receive_id = ? and post_id = ?)", postId, receiveId, postId, receiveId).Order("created_at desc").Limit(1).Find(&message)
	return message
}

func (d MessageDao) Save(message *model.Message) {
	constant.DB.Model(message).Create(message)
}

func (d MessageDao) GetByPostIdAndReceiveId(postId int64, receiveId int64, preMsgTime int64) *[]model.Message {
	var message *[]model.Message
	constant.DB.Model(message).Where("((post_id = ? and receive_id = ?)or(receive_id = ? and post_id = ?)) and created_at > ?", postId, receiveId, postId, receiveId, time.Unix(preMsgTime, 0).Format("2006-01-02 15:04:05")).Order("created_at").Find(&message)
	return message
}
