package service

import (
	"TheErrorCode/constant"
	"TheErrorCode/controller/resp"
	"TheErrorCode/dao"
	"TheErrorCode/model"
	"TheErrorCode/utils"
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

// 上传video到db
func saveVideoToDb(video *model.Video) {
	videoDao.Save(video)
}

// 上传文件到minio
func uploadnVideo(file *bytes.Buffer, fileName string) {

	go utils.Upload(fileName, file) //服务器带宽太低了等不起，默认上传成功把

}
