package service

import (
	"TheErrorCode/dao"
	"TheErrorCode/model"
)

type UserService struct {
}

var userDao = dao.UserDao{}

func (UserService) Register(user *model.User) {
	userDao.AddUser(user)
}

//func (UserService) Login(user *model.User) (string, error) {
//	dbUser, err := userDao.FindByUsername(user.UserName)
//	if err != nil {
//		log.Fatal("未找到用户")
//	}
//	if dbUser.Password != user.Password {
//
//	}
//	return
//}
