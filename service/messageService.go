package service

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/controller/resp"
	"TheErrorCode/model"
	"TheErrorCode/utils"
)

type MessageService struct {
}

func (s MessageService) PostMessage(req req.MessagePostRequest) interface{} {
	claims, _ := utils.ValidateJWT(req.Token)
	message := model.Message{
		PostId:    claims.ID,
		ReceiveId: req.ToUserId,
		Msg:       req.Content,
	}
	messageDao.Save(&message)
	messageResp := resp.UserResp{
		StatusCode: 0,
		StatusMsg:  "发送成功",
	}
	return &messageResp
}
