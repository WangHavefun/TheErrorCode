package controller

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/model"
	"TheErrorCode/service"
	"TheErrorCode/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoController struct {
}

var videoService = service.VideoService{}

func (VideoController) PublishVideo(c *gin.Context) {
	var req = req.DouYinPublishActionRequest{}
	c.ShouldBind(&req)
	file, _ := c.FormFile("data") //获取视频数据
	claims, _ := utils.ValidateJWT(req.Token)
	var video = model.Video{UserId: claims.ID}
	resp := videoService.UploadVideoAndSaveVideoToDb(file, &video)
	c.JSON(http.StatusOK, resp)
}

func (VideoController) VideoList(c *gin.Context) {
	var req = req.DouYinUserRequest{}
	c.ShouldBind(&req)
	claims, _ := utils.ValidateJWT(req.Token)
	resp := videoService.GetVideoList(claims.ID)
	c.JSON(http.StatusOK, resp)
}
