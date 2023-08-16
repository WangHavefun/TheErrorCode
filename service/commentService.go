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
	claims, _ := utils.ValidateJWT(r.Token)
	if r.ActionType == 1 {
		//发布评论
		comment := model.Comment{
			VideoId: r.VideoId,
			Text:    r.CommentText,
			UserId:  claims.ID,
		}
		commentDao.Save(&comment)
		commentVo.Content = comment.Text
		commentVo.CreateDate = comment.CreatedAt.Format("2006-01-02 15:04:05")
		commentVo.Id = comment.ID
	} else {
		//删除评论
		commentDao.DeleteById(r.CommentId)
	}

	userVo, _ := userService.GetUserInfoFromDb(claims.ID)
	commentVo.User = userVo
	commentResp := resp.CommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "操作成功",
		Comment:    commentVo,
	}
	return &commentResp
}

func (s CommentService) ListComment(r *req.CommentListRequest) interface{} {
	comments := commentDao.ListByVideoId(r.VideoId)
	var commentVos []vo.CommentVo
	for _, value := range *comments {
		commentVo := vo.CommentVo{
			Id:         value.ID,
			User:       nil,
			Content:    value.Text,
			CreateDate: value.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		userVo, _ := userService.GetUserInfoFromDb(value.UserId)
		commentVo.User = userVo
		commentVos = append(commentVos, commentVo)
	}
	listResp := resp.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "查询成功",
		CommentList: &commentVos,
	}
	return &listResp
}
