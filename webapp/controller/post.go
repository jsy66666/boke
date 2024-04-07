package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"webapp/logic"
	"webapp/models"
)

func CreatePostHandler(c *gin.Context) {
	//1.获取参数以及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c中获取当前发送请求的user_id
	userID, err := GetNowUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeSysBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, CodeSuccess, nil)
}

func GetPostDetail(c *gin.Context) {
	//1.从url中获取参数
	pidStr := c.Query("id")
	//将字符串类型解析成int64
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeGetParam)
		return
	}
	//2.根据id获取帖子
	postdetail, err := logic.GetPostDetail(pid)
	if err != nil {
		fmt.Println(err)
		ResponseError(c, CodeSysBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, CodeSuccess, postdetail)
}
