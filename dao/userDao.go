package dao

import (
	"TheErrorCode/constant"
	"TheErrorCode/model"
)

type UserDao struct {
}

func (UserDao) AddUser(user *model.User) {
	constant.DB.Create(&user)
}
func (UserDao) FindByUsername(username string) *model.User {
	var user = model.User{}
	constant.DB.Model(&user).Where("username=?", username).Find(&user)
	return &user
}

func (d UserDao) GetById(id int64) *model.User {
	var user = model.User{}
	constant.DB.Model(&user).Where("id=?", id).Find(&user)
	return &user
}
