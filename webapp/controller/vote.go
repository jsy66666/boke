package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"webapp/logic"
)

func VoteHandler(c *gin.Context) {
	pidStr := c.Param("id")
	postID, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeGetParam)
		return
	}
	userID, err := GetNowUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err := logic.CreatVote(postID, userID); err != nil {
		ResponseError(c, CodeSysBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess, nil)
}
