package controller

import (
	"TheErrorCode/model"
	"TheErrorCode/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DouYinUserRegisterRequest struct {
	Username string `form:"username"` // 注册用户名，最长32个字符
	Password string `form:"password"` // 密码，最长32个字符
}
type DouYinUserRegisterResponse struct {
	StatusCode int32  `json:"statusCode"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"statusMsg"`  // 返回状态描述
	UserId     int64  `json:"userId"`     // 用户id
	Token      string `json:"token"`      // 用户鉴权token
}
type UserController struct {
}

var userService = new(service.UserService)

func (UserController) Register(c *gin.Context) {
	var userReq = DouYinUserRegisterRequest{}
	c.ShouldBind(&userReq)
	var user = model.User{}
	user.UserName = userReq.Username
	user.Password = userReq.Password
	userService.Register(&user)
	c.JSON(http.StatusOK, DouYinUserRegisterResponse{
		UserId:     user.ID,
		StatusMsg:  "注册成功",
		StatusCode: 0,
		Token:      "",
	})
	return
}

func (UserController) Login(c *gin.Context) {
	var user = model.User{}
	c.ShouldBind(&user)
	//userService.Login(&user)
	fmt.Println(user)

}
