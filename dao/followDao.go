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

func (f *FollowDao) Save(follow *model.Follow) {
	constant.DB.Model(follow).Create(follow)
}

func (f *FollowDao) DeleteByFollowIdAndFanId(followId int64, fanId int64) {
	constant.DB.Model(&model.Follow{}).Where("follow_id = ? and fan_id = ?", followId, fanId).Delete(&model.Follow{})
}

func (f *FollowDao) ListByFanId(fanId int64) *[]model.Follow {
	var follows []model.Follow
	constant.DB.Model(&model.Follow{}).Where("fan_id = ?", fanId).Find(&follows)
	return &follows
}

func (f *FollowDao) ListByFollowId(followId int64) *[]model.Follow {
	var follows []model.Follow
	constant.DB.Model(&model.Follow{}).Where("follow_id = ?", followId).Find(&follows)
	return &follows
}

func (f *FollowDao) GetFollowByFollowIdAndFanId(followId int64, fanId int64) *model.Follow {
	var follow *model.Follow
	constant.DB.Model(&model.Follow{}).Where("follow_id = ? and fan_id = ?", followId, fanId).Find(&follow)
	return follow
}
