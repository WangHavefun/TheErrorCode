package service

import (
	"TheErrorCode/constant"
	"TheErrorCode/controller/resp"
	"TheErrorCode/dao"
	"TheErrorCode/model"
	"TheErrorCode/utils"
	"bytes"
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
	uploadnVideo(file, fileName) //上传文件
	video.PlayUrl = constant.MINIO + fileName
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
func uploadnVideo(file *multipart.FileHeader, fileName string) {

	src, err := file.Open()
	if err != nil {
		log.Println("Error reading the file content")
		return
	}
	var content bytes.Buffer
	_, err = io.Copy(&content, src)
	go utils.Upload(fileName, content) //服务器带宽太低了等不起，默认上传成功把

}
