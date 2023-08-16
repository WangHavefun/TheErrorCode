package controller

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentController struct {
}

var commentService = service.CommentService{}

func (CommentController) Comment(c *gin.Context) {
	commentReq := req.CommentActionRequest{}
	c.ShouldBind(&commentReq)
	resp := commentService.SaveComment(&commentReq)
	c.JSON(http.StatusOK, resp)
}

func (CommentController) List(c *gin.Context) {
	listReq := req.CommentListRequest{}
	c.ShouldBind(&listReq)
	resp := commentService.ListComment(&listReq)
	c.JSON(http.StatusOK, resp)
}
