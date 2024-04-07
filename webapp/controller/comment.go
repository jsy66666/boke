package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"webapp/logic"
	"webapp/models"
)

func NewCommentHandler(c *gin.Context) {
	comment := new(models.Comment)
	if err := c.ShouldBindJSON(&comment); err != nil {
		ResponseError(c, CodeGetParam)
		return
	}
	if len(comment.Content) == 0 {
		ResponseError(c, CodeGetParam)
		return
	}
	pidStr := c.Query("id")
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
	if err := logic.NewComment(comment, postID, userID); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeSysBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess, nil)
}

func UpdateCommentHandler(c *gin.Context) {
	commentIDStr := c.Query("comment_id")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeGetParam)
		return
	}
	updatecomment := new(models.UpadteComment)
	if err := c.ShouldBindJSON(&updatecomment); err != nil {
		ResponseError(c, CodeGetParam)
		return
	}
	if len(updatecomment.Content) == 0 {
		ResponseError(c, CodeGetParam)
		return
	}
	if err := logic.UpdateComment(updatecomment, commentID); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeSysBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess, nil)
}
func DeleteCommentHandler(c *gin.Context) {
	commentIDStr := c.Query("comment_id")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeGetParam)
		return
	}
	if err := logic.DeleteComment(commentID); err != nil {
		ResponseError(c, CodeSysBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess, nil)
}
