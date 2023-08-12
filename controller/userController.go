package controller

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/model"
	"TheErrorCode/service"
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
	var user = model.User{}
	c.ShouldBind(&user)
	resp := userService.Login(&user)
	c.JSON(http.StatusOK, resp)
}
