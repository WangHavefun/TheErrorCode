package service

import (
	"TheErrorCode/controller/req"
	resp2 "TheErrorCode/controller/resp"
	"TheErrorCode/dao"
	"TheErrorCode/model"
	"TheErrorCode/utils"
)

type FavoriteService struct {
}

var favoriteDao = dao.FavoriteDao{}

func (s FavoriteService) Action(r *req.FavoriteActionRequest) interface{} {
	claims, _ := utils.ValidateJWT(r.Token)
	action := model.Action{}
	action.UserId = claims.ID
	action.VideoId = r.VideoId
	if r.ActionType == 1 {
		action.Favorite = true
	} else {
		action.Favorite = false
	}
	favoriteDao.UpdateIfExistOrSave(&action)
	resp := resp2.UserResp{
		StatusCode: 1,
		StatusMsg:  "操作成功",
	}
	return &resp
}
