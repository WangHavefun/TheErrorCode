package controller

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageController struct {
}

var messageService = service.MessageService{}

func (MessageController) PostMessage(c *gin.Context) {
	messagePostReq := req.MessagePostRequest{}
	c.ShouldBind(&messagePostReq)
	resp := messageService.PostMessage(messagePostReq)
	c.JSON(http.StatusOK, &resp)
}
