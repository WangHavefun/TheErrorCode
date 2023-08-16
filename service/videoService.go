package service

import (
	"TheErrorCode/constant"
	"TheErrorCode/controller/req"
	"TheErrorCode/controller/resp"
	"TheErrorCode/dao"
	"TheErrorCode/model"
	"TheErrorCode/utils"
	"TheErrorCode/vo"
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strings"
)

type VideoService struct {
}

var videoDao = dao.VideoDao{}
var userService = UserService{}

func (s VideoService) UploadVideoAndSaveVideoToDb(file *multipart.FileHeader, video *model.Video) interface{} {
	fileName := utils.GenerateSubId() + "." + strings.Split(file.Filename, ".")[1]
	src, err := file.Open()
	if err != nil {
		resp := resp.UserResp{
			StatusCode: 1,
			StatusMsg:  "投稿失败",
		}
		log.Println("Error reading the file content")
		return resp
	}
	var content = bytes.Buffer{}
	_, err = io.Copy(&content, src)
	uploadnVideo(&content, fileName) //上传视频
	video.PlayUrl = constant.MINIO + fileName
	imageBuf, err := utils.ExtractFirstFrameAsImage(&content) //截取视频的第一帧图片
	if err != nil {
		resp := resp.UserResp{
			StatusCode: 1,
			StatusMsg:  "投稿失败",
		}
		log.Println(err.Error())
		return resp
	}
	fileName = utils.GenerateSubId() + ".jpeg" //图片名字
	fmt.Println(imageBuf)
	uploadnVideo(imageBuf, fileName)
	video.CoverUrl = constant.MINIO + fileName
	saveVideoToDb(video) //保存video信息到Db
	resp := resp.UserResp{
		StatusCode: 0,
		StatusMsg:  "投稿成功",
	}
	return resp
}

// 获取视频流的视频
func (s VideoService) GetPublicVideoList(feedReq *req.FeedRequest) (*[]vo.VideoVo, int64) {
	//TODO 增加最新时间
	var videoVos []vo.VideoVo
	var nextTime int64
	videos := videoDao.ListLimitCount(30)
	for _, value := range *videos {
		videoVo := vo.VideoVo{}
		videoVo.Title = value.Title
		videoVo.PlayUrl = value.PlayUrl
		videoVo.CoverUrl = value.CoverUrl
		videoVo.Id = value.ID
		userInfo, _ := userService.GetUserInfoFromDb(value.UserId)
		videoVo.Author = userInfo
		nextTime = value.CreatedAt.Unix() //获取最早投稿时间
		//获取是否点赞
		if feedReq.Token == "" {
			videoVo.IsFavorite = false
		} else {
			claims, _ := utils.ValidateJWT(feedReq.Token)
			action := favoriteDao.GetByUserIdAndVideoId(claims.ID, value.ID)
			videoVo.IsFavorite = action.Favorite
		}
		//获取点赞数
		count := favoriteDao.GetCountByVideoId(value.ID)
		videoVo.FavoriteCount = count
		videoVos = append(videoVos, videoVo)
	}
	return &videoVos, nextTime
}

// 获取自己喜欢的视频列表
func (VideoService) GetSelfFavoriteVideoList(userId int64) *[]vo.VideoVo {
	actions := favoriteDao.ListByUserId(userId)
	var videoVos []vo.VideoVo
	for _, value := range *actions {
		video := videoDao.GetById(value.VideoId)
		videoVo := vo.VideoVo{}
		videoVo.Title = video.Title
		videoVo.PlayUrl = video.PlayUrl
		videoVo.CoverUrl = video.CoverUrl
		videoVo.Id = video.ID
		//获取用户信息
		userInfo, _ := userService.GetUserInfoFromDb(video.UserId)
		videoVo.Author = userInfo
		//获取是否点赞
		videoVo.IsFavorite = true //都查点赞列表了还能不点赞？
		//获取点赞数
		count := favoriteDao.GetCountByVideoId(value.VideoId)
		videoVo.FavoriteCount = count
		videoVos = append(videoVos, videoVo)
	}
	return &videoVos

}

// 获取自己投稿的视频
func (s VideoService) GetSelfVideoList(userId int64) interface{} {

	videos := videoDao.ListByUserId(userId)
	var videoVos []vo.VideoVo
	userInfo, err := userService.GetUserInfoFromDb(userId)
	if err != nil {
		log.Println(err.Error())
		resp := resp.VideoListResp{
			StatusCode: 1,
			StatusMsg:  "获取失败",
			VideoList:  nil,
		}
		return resp
	}
	for _, value := range *videos {
		videoVo := vo.VideoVo{}
		videoVo.Title = value.Title
		videoVo.PlayUrl = value.PlayUrl
		videoVo.CoverUrl = value.CoverUrl
		videoVo.Id = value.ID
		videoVo.Author = userInfo
		//获取是否点赞
		action := favoriteDao.GetByUserIdAndVideoId(userId, value.ID)
		videoVo.IsFavorite = action.Favorite
		//获取点赞数
		count := favoriteDao.GetCountByVideoId(value.ID)
		videoVo.FavoriteCount = count
		videoVos = append(videoVos, videoVo)
	}

	resp := resp.VideoListResp{
		StatusCode: 0,
		StatusMsg:  "获取成功",
		VideoList:  videoVos,
	}
	return resp
}

// 上传video到db
func saveVideoToDb(video *model.Video) {
	videoDao.Save(video)
}

// 上传文件到minio
func uploadnVideo(file *bytes.Buffer, fileName string) {

	go utils.Upload(fileName, file) //服务器带宽太低了等不起，默认上传成功把

}
