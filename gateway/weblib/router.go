package weblib

import (
	"gateway/weblib/handlers"
	"gateway/weblib/middleware"
	"gateway/weblib/middleware/jwt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	douyin := ginRouter.Group("/douyin")
	{
		//测试
		douyin.GET("/ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})

		douyin.GET("/feed/", jwt.AuthWithoutLogin(), handlers.FeedInfo)
		douyin.POST("/publish/action/", jwt.AuthBody(), handlers.PublishAction)
		douyin.GET("/publish/list/", jwt.Auth(), handlers.PublishList)
		douyin.GET("/user/", jwt.Auth(), handlers.UserInfo)
		douyin.POST("/user/register/", handlers.UserRegister)
		douyin.POST("/user/login/", handlers.UserLogin)

		douyin.POST("/favorite/action/", jwt.Auth(), handlers.FavoriteAction)
		douyin.GET("/favorite/list/", jwt.Auth(), handlers.FavoriteList)
		douyin.POST("/comment/action/", jwt.Auth(), handlers.CommentAction)
		douyin.GET("/comment/list/", jwt.AuthWithoutLogin(), handlers.CommentList)

		douyin.POST("/relation/action/", jwt.Auth(), handlers.RelationAction)
		douyin.GET("/relation/follow/list/", jwt.Auth(), handlers.FollowList)
		douyin.GET("/relation/follower/list", jwt.Auth(), handlers.FollowerList)

	}
	return ginRouter
}
