package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回响应

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.getMsg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, code ResCode, data interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.getMsg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
