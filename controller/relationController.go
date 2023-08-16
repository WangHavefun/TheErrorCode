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
