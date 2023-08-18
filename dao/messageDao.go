package dao

import (
	"TheErrorCode/constant"
	"TheErrorCode/model"
)

type MessageDao struct {
}

func (d MessageDao) GetLatestByPostIdAndReceiveId(postId int64, receiveId int64) *model.Message {
	var message *model.Message
	constant.DB.Model(message).Where("(post_id = ? and receive_id = ?)or(receive_id = ? and post_id = ?)", postId, receiveId, receiveId, postId).Order("created_at").Limit(1).Find(&message)
	return message
}

func (d MessageDao) Save(message *model.Message) {
	constant.DB.Model(message).Create(message)
}
