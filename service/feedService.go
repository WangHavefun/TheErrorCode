package service

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/controller/resp"
	"time"
)

type FeedService struct {
}

var videoService = VideoService{}

func (s FeedService) GetFeed(feedReq *req.FeedRequest) interface{} {
	if feedReq.LatestTime == 0 {
		feedReq.LatestTime = time.Now().Unix()
	}
	videoVos, nextTime := videoService.GetPublicVideoList(feedReq)
	resp := resp.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "获取成功",
		VideoList:  videoVos,
		NextTime:   nextTime,
	}
	return &resp

}
