package dao

import (
	"TheErrorCode/constant"
	"TheErrorCode/model"
	"errors"
	"log"
)

type UserDao struct {
}

func (UserDao) AddUser(user *model.User) {
	constant.DB.Create(&user)
}
func (UserDao) FindByUsername(username string) (*model.User, error) {
	var user = model.User{}
	result := constant.DB.Model(&user).Where("username=?", username).Find(&user)
	if result.RowsAffected == 0 {
		return nil, errors.New("未查询到")
	}
	return &user, nil
}
func (d UserDao) GetById(id int64) (*model.User, error) {
	var user = model.User{}
	result := constant.DB.Model(&user).Where("id=?", id).Find(&user)
	if result.RowsAffected == 0 {
		log.Println("未查询到")
		return nil, errors.New("未查询到")
	}
	return &user, nil
}
