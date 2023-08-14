package dao

import (
	"TheErrorCode/constant"
	"TheErrorCode/model"
)

type VideoDao struct {
}

func (VideoDao) Save(video *model.Video) {
	constant.DB.Model(&video).Create(&video)
}
