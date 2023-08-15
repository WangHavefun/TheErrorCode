package dao

import (
	"TheErrorCode/constant"
	"TheErrorCode/model"
)

type VideoDao struct {
}

// 保存一条数据
func (VideoDao) Save(video *model.Video) {
	constant.DB.Model(&video).Create(&video)
}

// 查询很多数据
func (VideoDao) ListByUserId(userId int64) *[]model.Video {
	videos := []model.Video{}
	constant.DB.Model(&videos).Where("user_id = ?", userId).Find(&videos)
	return &videos
}
func (VideoDao) CountByUserId(userId int64) (count int64) {
	constant.DB.Model(&model.Video{}).Where("user_id = ?", userId).Count(&count)
	return count
}
