package controller

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/model"
	"TheErrorCode/service"
	"TheErrorCode/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

var userService = new(service.UserService)

func (UserController) Register(c *gin.Context) {
	var userReq = req.DouYinUserRegLogRequest{}
	c.ShouldBind(&userReq) //绑定参数
	var user = model.User{}
	user.UserName = userReq.Username
	user.Password = userReq.Password
	resp := userService.Register(&user)
	c.JSON(http.StatusOK, resp)

}

func (UserController) Login(c *gin.Context) {
	var userReq = req.DouYinUserRegLogRequest{}
	c.ShouldBind(&userReq)
	var user = model.User{}
	user.UserName = userReq.Username
	user.Password = userReq.Password
	resp := userService.Login(&user)
	c.JSON(http.StatusOK, resp)
}
func (UserController) GetUser(c *gin.Context) {
	userReq := req.DouYinUserRequest{}
	c.ShouldBind(&userReq)
	toekn := userReq.Token
	claims, _ := utils.ValidateJWT(toekn) //解析token获取数据
	user := model.User{}
	user.UserName = claims.Username
	user.ID = claims.ID
	resp := userService.GetUserInfoFromDB(user)
	c.JSON(http.StatusOK, resp)
}
