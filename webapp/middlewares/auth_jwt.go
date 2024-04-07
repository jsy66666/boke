package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"webapp/controller"
	"webapp/pkg/jwt"
)

func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		//判断格式是否正确
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		//parts[1]获取到的是tokenstring,解析tokenstring
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		//将当前请求的userID保存到请求的上下文c上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
