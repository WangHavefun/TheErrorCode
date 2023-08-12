package dao

import (
	"TheErrorCode/constant"
	"TheErrorCode/model"
)

type UserDao struct {
}

func (UserDao) AddUser(user *model.User) {
	constant.DB.Create(user)
}
