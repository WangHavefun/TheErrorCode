package controller

import (
	"TheErrorCode/controller/req"
	"TheErrorCode/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavoriteController struct {
}

var favoriteService = service.FavoriteService{}

func (FavoriteController) Action(c *gin.Context) {
	favoriteActionRequest := req.FavoriteActionRequest{}
	c.ShouldBind(&favoriteActionRequest)
	resp := favoriteService.Action(&favoriteActionRequest)
	c.JSON(http.StatusOK, resp)
}
