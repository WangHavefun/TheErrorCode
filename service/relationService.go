package service

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/controller/resp"
	"TheErrorCode/dao"
	"TheErrorCode/model"
	"TheErrorCode/utils"
	"TheErrorCode/vo"
)

type FollowService struct {
}

var messageDao = dao.MessageDao{}

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

func (s FollowService) FollowList(r *req.DouYinUserRequest) interface{} {
	follows := followDao.ListByFanId(r.UserId)
	var userVos []vo.UserVo
	for _, value := range *follows {
		userInfo, _ := userService.GetUserInfoFromDb(value.FollowId, r.UserId)
		userVos = append(userVos, *userInfo)
	}
	followResp := resp.RelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		UserList:   &userVos,
	}
	return &followResp
}

func (s FollowService) FanList(r *req.DouYinUserRequest) interface{} {
	follows := followDao.ListByFollowId(r.UserId)
	var userVos []vo.UserVo
	for _, value := range *follows {
		userInfo, _ := userService.GetUserInfoFromDb(value.FanId, r.UserId)
		userVos = append(userVos, *userInfo)
	}
	followResp := resp.RelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		UserList:   &userVos,
	}
	return &followResp
}

func (s FollowService) FriendList(r *req.DouYinUserRequest) interface{} {
	follows := followDao.ListByFollowId(r.UserId)
	var friendUserVos []vo.FriendUserVo
	for _, value := range *follows {
		friendUserVo := vo.FriendUserVo{}
		userInfo, _ := userService.GetUserInfoFromDb(value.FanId, r.UserId)
		friendUserVo.UserVo = userInfo
		message := messageDao.GetLatestByPostIdAndReceiveId(r.UserId, userInfo.Id)
		friendUserVo.Message = message.Msg
		if message.PostId == r.UserId {
			friendUserVo.MsgType = 1
		} else {
			friendUserVo.MsgType = 0
		}
		friendUserVos = append(friendUserVos, friendUserVo)
	}
	followResp := resp.RelationFriendListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		UserList:   &friendUserVos,
	}
	return &followResp
}
