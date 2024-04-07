package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "user_id"

var ErrUserNotLogin = errors.New("用户未登录")

func GetNowUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrUserNotLogin
		return
	}
	return
}
