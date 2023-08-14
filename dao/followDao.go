package dao

import (
	"TheErrorCode/constant"
	"TheErrorCode/model"
	"errors"
)

type FollowDao struct {
}

func (f *FollowDao) GetFollowByFollowId(followId int64) (*[]model.Follow, error) {
	var follows []model.Follow
	constant.DB.Model(&model.Follow{}).Where("follow_id=?", followId).Find(&follows)
	if follows == nil {
		return nil, errors.New("未查询到")
	}
	return &follows, nil
}
func (f *FollowDao) GetFollowByFanId(fanId int64) (*[]model.Follow, error) {
	var follows []model.Follow
	constant.DB.Model(&model.Follow{}).Where("fan_id=?", fanId).Find(&follows)
	if follows == nil {
		return nil, errors.New("未查询到")
	}
	return &follows, nil
}
