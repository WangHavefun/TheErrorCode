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
