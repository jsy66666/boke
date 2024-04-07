package controller

import (
	"github.com/gin-gonic/gin"
	"webapp/logic"
	"webapp/models"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		ResponseError(c, CodeGetParam)
		return
	}
	//2.手动校验参数
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.Repassword) == 0 || p.Password != p.Repassword {
		ResponseError(c, CodeGetParam)
		return
	}
	//2.业务处理
	if isSingupOk := logic.SignUp(p); !isSingupOk {
		ResponseError(c, CodeUserExist)
		return
	}
	//3.返回响应
	ResponseSuccess(c, CodeSingUpSuccess, nil)
}

func LoginHandler(c *gin.Context) {
	//1.获取参数
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeGetParam)
		return
	}
	//2.业务处理
	token, err := logic.Login(p)
	if err != nil {
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3.返回响应,返回响应需要返回给前端一个token
	ResponseSuccess(c, CodeLoginSuccess, token)
}
