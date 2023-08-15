package controller

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedController struct {
}

var feedService = service.FeedService{}

func (f FeedController) Feed(c *gin.Context) {
	feedReq := req.FeedRequest{}
	c.ShouldBind(&feedReq)
	resp := feedService.GetFeed(&feedReq)
	c.JSON(http.StatusOK, resp)
}
