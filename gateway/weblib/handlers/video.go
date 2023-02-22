package handlers

import (
	"context"
	"gateway/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//视频流
func FeedInfo(ginCtx *gin.Context) {
	//允许传参username,password为空
	var feedReq services.FeedRequest
	PanicIfUserError(ginCtx.Bind(&feedReq))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["videoService"].(services.VideoService)

	feedResp, err := videoService.FeedInfo(context.Background(), &feedReq)
	PanicIfUserError(err)
	//ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": feedResp.StatusCode,
		"status_msg":  feedResp.StatusMsg,
		//"video_list":     feedResp.VideoList,
		//"next_time":       feedResp.NextTime,
	})
}

//发布视频
func PublishAction(ginCtx *gin.Context) {

}

//自己的发布视频列表
func PublishList(ginCtx *gin.Context) {

}

//点赞/取消点赞
func FavoriteAction(ginCtx *gin.Context) {
	var favoriteActionRequest services.FavoriteActionRequest
	PanicIfUserError(ginCtx.ShouldBind(&favoriteActionRequest))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["videoService"].(services.VideoService)

	//从上下文中获取由token拦截器获取到的userId也赋值给context上下文
	strUserId := ginCtx.GetString("userId")
	ctx := context.WithValue(ginCtx.Request.Context(), "userId", strUserId)
	ginCtx.Request = ginCtx.Request.WithContext(ctx)

	feedResp, err := videoService.FavoriteAction(ginCtx.Request.Context(), &favoriteActionRequest)
	PanicIfUserError(err)
	//ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": feedResp.StatusCode,
		"status_msg":  feedResp.StatusMsg,
	})
}

//点赞列表
func FavoriteList(ginCtx *gin.Context) {
	var favoriteListRequest services.FavoriteListRequest
	PanicIfUserError(ginCtx.ShouldBind(&favoriteListRequest))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["videoService"].(services.VideoService)

	//从上下文中获取由token拦截器获取到的userId也赋值给context上下文
	strUserId := ginCtx.GetString("userId")
	ctx := context.WithValue(ginCtx.Request.Context(), "userId", strUserId)
	ginCtx.Request = ginCtx.Request.WithContext(ctx)

	feedResp, err := videoService.FavoriteList(ginCtx.Request.Context(), &favoriteListRequest)
	PanicIfUserError(err)
	//ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": feedResp.StatusCode,
		"status_msg":  feedResp.StatusMsg,
		"video_list":  feedResp.VideoList,
	})
}

//评论
func CommentAction(ginCtx *gin.Context) {
	var commentActionRequest services.CommentActionRequest
	PanicIfUserError(ginCtx.ShouldBind(&commentActionRequest))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["videoService"].(services.VideoService)

	//从上下文中获取由token拦截器获取到的userId也赋值给context上下文
	strUserId := ginCtx.GetString("userId")
	ctx := context.WithValue(ginCtx.Request.Context(), "userId", strUserId)
	ginCtx.Request = ginCtx.Request.WithContext(ctx)

	feedResp, err := videoService.CommentAction(ginCtx.Request.Context(), &commentActionRequest)
	PanicIfUserError(err)
	//ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": feedResp.StatusCode,
		"status_msg":  feedResp.StatusMsg,
		"comment":     feedResp.Comment,
	})
}

//评论列表
func CommentList(ginCtx *gin.Context) {
	var commentListRequest services.CommentListRequest
	PanicIfUserError(ginCtx.ShouldBind(&commentListRequest))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["videoService"].(services.VideoService)

	//从上下文中获取由token拦截器获取到的userId也赋值给context上下文
	strUserId := ginCtx.GetString("userId")
	ctx := context.WithValue(ginCtx.Request.Context(), "userId", strUserId)
	ginCtx.Request = ginCtx.Request.WithContext(ctx)

	feedResp, err := videoService.CommentList(ginCtx.Request.Context(), &commentListRequest)
	PanicIfUserError(err)
	//ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code":  feedResp.StatusCode,
		"status_msg":   feedResp.StatusMsg,
		"comment_list": feedResp.CommentList,
	})
}
