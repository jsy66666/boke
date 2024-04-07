package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"webapp/controller"
	"webapp/middlewares"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()

	v1 := r.Group("/api/v1")
	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JwtAuthMiddleware(), middlewares.RateLimitMiddleware(2*time.Millisecond, 1))
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/communitydetail", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/postdetail", controller.GetPostDetail)

		//投票和取消投票
		v1.POST("/vote/:id", controller.VoteHandler)
		//新建评论
		v1.POST("/createcomment", controller.NewCommentHandler)
		//修改评论
		v1.POST("/updatecomment", controller.UpdateCommentHandler)
		//删除评论
		v1.POST("/deletecomment", controller.DeleteCommentHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
