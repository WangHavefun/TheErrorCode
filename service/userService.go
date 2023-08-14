package service

import (
	"TheErrorCode/constant"
	"TheErrorCode/controller/resp"
	"TheErrorCode/dao"
	"TheErrorCode/model"
	"TheErrorCode/utils"
	"log"
)

type UserService struct {
}

var userDao = dao.UserDao{}
var followDao = dao.FollowDao{}

func (UserService) Register(user *model.User) *resp.DouYinUserRegLogResponse {
	resp := resp.DouYinUserRegLogResponse{}
	_, err := userDao.FindByUsername(user.UserName)
	if err == nil {
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
	dbUser, err := userDao.FindByUsername(user.UserName)
	resp := resp.DouYinUserRegLogResponse{}

	if err != nil {
		log.Println(err.Error())
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

func (s UserService) GetUserInfoFromDB(user model.User) interface{} {
	userResp := resp.User{}
	dbUser, err := userDao.GetById(user.ID)
	if err != nil {
		log.Println(err.Error())
		resp := resp.UserResp{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}
		return &resp
	}
	userResp.Avatar = dbUser.Avatar
	userResp.BackgroundImage = dbUser.BackgroundImage
	userResp.Id = dbUser.ID
	userResp.Name = dbUser.Name
	userResp.Signature = dbUser.Signature
	//获取粉丝
	fans, err := followDao.GetFollowByFollowId(dbUser.ID)
	if err != nil {
		log.Println(err.Error())
		userResp.FollowerCount = 0
	} else {
		userResp.FollowerCount = int64(len(*fans))
	}
	//获取关注的
	follows, err := followDao.GetFollowByFanId(dbUser.ID)
	if err != nil {
		log.Println(err.Error())
		userResp.FollowCount = 0
	} else {
		userResp.FollowCount = int64(len(*follows))
	}
	resp := resp.UserResp{
		StatusCode: 0,
		StatusMsg:  "获取成功",
		User:       userResp,
	}
	return &resp
}
