package service

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/controller/resp"
	"TheErrorCode/model"
	"TheErrorCode/utils"
	"TheErrorCode/vo"
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

func (s MessageService) MessageList(chatReq req.MessageChatRequest) interface{} {
	claims, _ := utils.ValidateJWT(chatReq.Token)
	messages := messageDao.GetByPostIdAndReceiveId(claims.ID, chatReq.ToUserId, chatReq.PreMsgTime)
	var messageVos []vo.MessageVo
	for _, value := range *messages {
		messageVo := vo.MessageVo{
			Id:         value.ID,
			ToUserId:   value.ReceiveId,
			FromUserId: value.PostId,
			Content:    value.Msg,
			CreateTime: value.CreatedAt.Unix(),
		}
		messageVos = append(messageVos, messageVo)
	}
	messageListResp := resp.MessageChatResponse{
		StatusCode:  0,
		StatusMsg:   "获取成功",
		MessageList: &messageVos,
	}
	return &messageListResp
}
