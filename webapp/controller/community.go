package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"webapp/logic"
)

func CommunityHandler(c *gin.Context) {
	//查询所有的社区(id,name)以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		ResponseError(c, CodeSysBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess, data)
}

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Query("id")
	communityID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidParam)
		return
	}
	communityDetail, err := logic.GetCommunityDetail(communityID)
	if err != nil {
		ResponseError(c, CodeSysBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess, communityDetail)
}
