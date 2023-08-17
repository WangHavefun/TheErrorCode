package controller

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowController struct {
}

var followService = service.FollowService{}

func (FollowController) Follow(c *gin.Context) {
	relationReq := req.RelationActionRequest{}
	c.ShouldBind(&relationReq)
	relationResp := followService.Follow(&relationReq)
	c.JSON(http.StatusOK, relationResp)
}

func (FollowController) FollowList(c *gin.Context) {
	relationReq := req.DouYinUserRequest{}
	c.ShouldBind(&relationReq)
	relationResp := followService.FollowList(&relationReq)
	c.JSON(http.StatusOK, relationResp)
}

func (FollowController) FanList(c *gin.Context) {
	relationReq := req.DouYinUserRequest{}
	c.ShouldBind(&relationReq)
	relationResp := followService.FanList(&relationReq)
	c.JSON(http.StatusOK, relationResp)
}

func (FollowController) FriendList(c *gin.Context) {
	relationReq := req.DouYinUserRequest{}
	c.ShouldBind(&relationReq)
	relationResp := followService.FriendList(&relationReq)
	c.JSON(http.StatusOK, relationResp)
}
