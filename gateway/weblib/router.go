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
		douyin.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})

		//Feed
		//douyin.GET("/feed",)
		// 用户服务
		douyin.POST("/user", handlers.UserInfo)
		douyin.POST("/user/register", handlers.UserRegister)
		douyin.POST("/user/login", handlers.UserLogin)

		//// 需要登录保护
		//authed := douyin.Group("/")
		//authed.Use(middleware.JWT())
		//{
		//	authed.GET("tasks", handlers.GetTaskList)
		//	authed.POST("task", handlers.CreateTask)
		//	authed.GET("task/:id", handlers.GetTaskDetail) // task_id
		//	authed.PUT("task/:id", handlers.UpdateTask)    // task_id
		//	authed.DELETE("task/:id", handlers.DeleteTask) // task_id
		//}
	}
	return ginRouter
}
