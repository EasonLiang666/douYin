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

}

//点赞列表
func FavoriteList(ginCtx *gin.Context) {

}

//评论
func CommentAction(ginCtx *gin.Context) {

}

//评论列表
func CommentList(ginCtx *gin.Context) {

}
