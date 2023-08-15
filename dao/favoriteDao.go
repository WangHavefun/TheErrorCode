package dao

import (
	"TheErrorCode/constant"
	"TheErrorCode/model"
)

type FavoriteDao struct {
}

func (d FavoriteDao) UpdateIfExistOrSave(action *model.Action) {
	var count int64

	constant.DB.Model(&action).Where("user_id = ? and video_id = ?", action.UserId, action.VideoId).Count(&count)
	if count == 0 {
		constant.DB.Model(&action).Create(action)
	} else {

		constant.DB.Model(&action).Where("user_id = ? and video_id = ?", action.UserId, action.VideoId).Delete(&action)
	}
}

func (d FavoriteDao) GetByUserIdAndVideoId(userId int64, videoId int64) *model.Action {
	var action = model.Action{}
	constant.DB.Model(&action).Where("user_id = ? and video_id = ?", userId, videoId).Find(&action)
	return &action
}

func (d FavoriteDao) GetCountByVideoId(videoId int64) int64 {
	var count int64
	constant.DB.Model(&model.Action{}).Where("video_id = ? ", videoId).Count(&count)
	return count
}

func (d FavoriteDao) GetCountByUserId(userId int64) int64 {
	var count int64
	constant.DB.Model(&model.Action{}).Where("user_id = ? ", userId).Count(&count)
	return count
}

func (d FavoriteDao) ListByUserId(userId int64) *[]model.Action {
	var actions []model.Action
	constant.DB.Model(&model.Action{}).Where("user_id = ?", userId).Find(&actions)
	return &actions
}
