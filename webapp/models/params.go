package models

// 定义请求参数的结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Repassword string `json:"repassword" binding:"required"`
}

type ParamLogin struct {
	UserID   int64  `json:"userid"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
