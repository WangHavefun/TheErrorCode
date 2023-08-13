package service

import (
	"TheErrorCode/constant"
	"TheErrorCode/controller/resp"
	"TheErrorCode/dao"
	"TheErrorCode/model"
	"TheErrorCode/utils"
)

type UserService struct {
}

var userDao = dao.UserDao{}

func (UserService) Register(user *model.User) *resp.DouYinUserRegLogResponse {
	resp := resp.DouYinUserRegLogResponse{}
	dbuser, _ := userDao.FindByUsername(user.UserName)
	if dbuser.ID != 0 {
		resp.StatusCode = 1
		resp.Token = ""
		resp.StatusMsg = "用户名已经存在"
		resp.UserId = 0
		return &resp
	}
	user.Avatar = constant.DEFAULTIMAGES
	user.BackgroundImage = constant.DEFAULTIMAGES
	userDao.AddUser(user)
	token, _ := utils.CreateJWT(user.ID, user.UserName)
	resp.StatusCode = 0
	resp.Token = token
	resp.StatusMsg = "注册成功"
	resp.UserId = user.ID
	return &resp
}

func (UserService) Login(user *model.User) *resp.DouYinUserRegLogResponse {
	dbUser, _ := userDao.FindByUsername(user.UserName)
	resp := resp.DouYinUserRegLogResponse{}
	if dbUser.ID == 0 {
		resp.StatusCode = 1
		resp.Token = ""
		resp.StatusMsg = "用户名不存在"
		resp.UserId = 0
		return &resp
	}
	if dbUser.Password != user.Password {
		resp.StatusCode = 1
		resp.Token = ""
		resp.UserId = 0
		resp.StatusMsg = "密码错误"
		return &resp
	}
	token, _ := utils.CreateJWT(dbUser.ID, dbUser.UserName)
	resp.StatusCode = 0
	resp.Token = token
	resp.UserId = dbUser.ID
	resp.StatusMsg = "登录成功"
	return &resp
}
