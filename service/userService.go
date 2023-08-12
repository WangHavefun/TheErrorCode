package service

import (
	"TheErrorCode/dao"
	"TheErrorCode/model"
)

type UserService struct {
}

var userDao = dao.UserDao{}

func (UserService) AddUserToDB(user *model.User) {
	userDao.AddUser(user)
}
