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
func (UserDao) FindByUsername(username string) (*model.User, error) {
	var user = model.User{}
	constant.DB.Where("username=?", username).Find(&user)
	return &user, nil
}
