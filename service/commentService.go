package service

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/controller/resp"
	"TheErrorCode/dao"
	"TheErrorCode/model"
	"TheErrorCode/utils"
	"TheErrorCode/vo"
)

type CommentService struct {
}

var commentDao = dao.CommentDao{}

func (s CommentService) SaveComment(r *req.CommentActionRequest) interface{} {
	commentVo := vo.CommentVo{}
	if r.ActionType == 1 {
		//发布评论
		comment := model.Comment{
			VideoId: r.VideoId,
			Text:    r.CommentText,
		}
		commentDao.Save(&comment)
		commentVo.Content = comment.Text
		commentVo.CreateDate = string(comment.CreatedAt.Month()) + "-" + string(comment.CreatedAt.Day())
		commentVo.Id = comment.ID
	} else {
		//删除评论
		commentDao.DeleteById(r.CommentId)
	}
	claims, _ := utils.ValidateJWT(r.Token)
	userVo, _ := userService.GetUserInfoFromDb(claims.ID)
	commentVo.User = userVo
	commentResp := resp.CommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "操作成功",
		Comment:    commentVo,
	}
	return &commentResp
}
