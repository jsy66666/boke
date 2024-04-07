package mysql

import (
	"errors"
	"webapp/models"
)

var UserInsertFailed string = "添加用户信息失败"
var UserNotExit string = "登陆失败,用户不存在"
var UserInfoError string = "登陆失败,用户名或密码错误"

func CheckUserExit(name string) (bool, error) {
	sqlStr := "select count(user_id) from user where user_name = ? "
	var count int
	if err := db.Get(&count, sqlStr, name); err == nil {
		if count == 0 {
			return false, nil
		} else {
			return true, nil
		}
	} else {
		return true, err
	}
}

func InsertUser(user *models.User) (err error) {
	sqlStr := "insert into user ( user_id , user_name , password) values (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	if err != nil {
		return errors.New(UserInsertFailed)
	}
	return nil
}

func Login(p *models.ParamLogin) (err error) {

	sqlStr := "select user_id,password from user where user_name = ?"
	var user = new(models.User)
	if err := db.Get(user, sqlStr, p.Username); err != nil {
		return errors.New(UserNotExit)
	}
	if user.Password != p.Password {
		return errors.New(UserInfoError)
	}
	p.UserID = user.UserID
	return nil
}
