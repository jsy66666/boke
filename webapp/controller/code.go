package controller

// 定义状态码  每个状态码对应返回给前端的json数据
type ResCode int64

const (
	CodeSingUpSuccess ResCode = 1000 + iota
	CodeLoginSuccess
	CodeGetParam
	CodeCheckParam
	CodeUserExist

	CodeNeedLogin
	CodeInvalidToken

	CodeUserNotExist
	CodeInvalidPassword

	CodeInvalidParam

	CodeSysBusy
	CodeSuccess
)

var codeMsgMap = map[ResCode]string{
	CodeGetParam:        "请求参数有误",
	CodeCheckParam:      "校验参数有误",
	CodeSingUpSuccess:   "注册成功",
	CodeLoginSuccess:    "登录成功",
	CodeUserExist:       "用户已经存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",

	CodeNeedLogin:    "需要先登录",
	CodeInvalidToken: "无效的token",
	CodeSysBusy:      "系统繁忙",
	CodeInvalidParam: "参数错误",
	CodeSuccess:      "成功",
}

func (code ResCode) getMsg() string {
	return codeMsgMap[code]
}
