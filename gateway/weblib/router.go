package weblib

import (
	"gateway/weblib/handlers"
	"gateway/weblib/middleware"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

		//Feed
		douyin.POST("/feed", handlers.FeedInfo)

		// 用户服务
		douyin.POST("/user/register", handlers.UserRegister)
		douyin.POST("/user/login", handlers.UserLogin)

		// 需要登录保护
		authed := douyin.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("user", handlers.UserInfo)
			authed.POST("publish/action/", handlers.PublishAction)
			authed.POST("publish/list/", handlers.PublishList)
			authed.POST("favorite/action/", handlers.PublishAction)
			authed.POST("favorite/list/", handlers.FavoriteList)
			authed.POST("comment/action/", handlers.CommentAction)
			authed.POST("comment/list/", handlers.CommentList)
			authed.POST("relation/action/", handlers.RelationAction)
			authed.POST("relation/follow/list/", handlers.FollowList)
			authed.POST("relation/follower/list/", handlers.FollowerList)
			authed.POST("relation/friend/list/", handlers.FriendList)
			authed.POST("message/action/", handlers.SendMessage)
			authed.POST("message/chat/", handlers.MessageChatRecord)
		}
	}
	return ginRouter
}
