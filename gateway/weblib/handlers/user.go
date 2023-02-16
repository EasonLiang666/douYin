package handlers

import (
	"context"
	"gateway/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取用户信息
func UserInfo(ginCtx *gin.Context) {

}

// 用户注册
func UserRegister(ginCtx *gin.Context) {
	var userRegisterReq services.UserRegisterRequest
	PanicIfUserError(ginCtx.Bind(&userRegisterReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)

	userResp, err := userService.UserRegister(context.Background(), &userRegisterReq)
	PanicIfUserError(err)
	//ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": userResp.StatusCode,
		"status_msg":  userResp.StatusMsg,
		"user_id":     userResp.UserId,
		"token":       userResp.Token,
	})
}

// 用户登录
func UserLogin(ginCtx *gin.Context) {

}

//消息记录
func MessageChatRecord(ginCtx *gin.Context) {

}

//发送消息
func SendMessage(ginCtx *gin.Context) {

}

//关注/取关
func RelationAction(ginCtx *gin.Context) {

}

//关注列表
func FollowList(ginCtx *gin.Context) {

}

//粉丝列表
func FollowerList(ginCtx *gin.Context) {

}

//朋友列表
func FriendList(ginCtx *gin.Context) {

}
