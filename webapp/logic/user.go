package logic

import (
	"webapp/dao/mysql"
	"webapp/models"
	"webapp/pkg/jwt"
	"webapp/pkg/snowflake"
)

func SignUp(u *models.ParamSignUp) bool {
	//判断用户存不存在,如果用户存在就插入失败,如果不存在就插入数据库
	if isExsit, _ := mysql.CheckUserExit(u.Username); isExsit {
		return false
	} else {
		//生成UID,将新注册用户的信息保存到数据库
		var userID int64
		node, err := snowflake.NewWorker(1)
		if err != nil {
			return false
		} else {
			userID = node.GetId()
		}
		user := &models.User{
			UserID:   userID,
			Username: u.Username,
			Password: u.Password,
		}
		//保存到数据库
		mysql.InsertUser(user)
		return true
	}
}

func Login(u *models.ParamLogin) (token string, err error) {
	//如果登陆失败返回空的token 返回错误
	if err = mysql.Login(u); err != nil {
		return "", err
	}
	//如果登录成功那就生成jwt token
	return jwt.GenToken(u.UserID)
}
