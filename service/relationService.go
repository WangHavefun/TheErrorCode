package service

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/controller/resp"
	"TheErrorCode/model"
	"TheErrorCode/utils"
)

type FollowService struct {
}

func (s FollowService) Follow(r *req.RelationActionRequest) interface{} {
	claims, _ := utils.ValidateJWT(r.Token)
	follow := model.Follow{
		FollowId: r.ToUserId,
		FanId:    claims.ID,
	}
	if r.ActionType == 1 {
		followDao.Save(&follow)
	} else {
		followDao.DeleteByFollowIdAndFanId(r.ToUserId, claims.ID)
	}
	relationResp := resp.UserResp{
		StatusCode: 0,
		StatusMsg:  "操作成功",
	}

	return &relationResp
}
