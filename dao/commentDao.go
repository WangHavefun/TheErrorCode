package dao

import (
	"TheErrorCode/constant"
	"TheErrorCode/model"
)

type CommentDao struct {
}

func (d CommentDao) Save(comment *model.Comment) {
	constant.DB.Model(&comment).Create(&comment)
}

func (d CommentDao) DeleteById(id int64) {
	constant.DB.Model(&model.Comment{}).Where("id = ? ", id).Delete(&model.Comment{})
}

func (d CommentDao) ListByVideoId(videoId int64) *[]model.Comment {
	var comments []model.Comment
	constant.DB.Model(model.Comment{}).Where("video_id = ?", videoId).Find(&comments)
	return &comments
}

func (d CommentDao) GetCountByVideoId(videoId int64) int64 {
	var count int64
	constant.DB.Model(&model.Comment{}).Where("video_id = ?", videoId).Count(&count)
	return count
}
