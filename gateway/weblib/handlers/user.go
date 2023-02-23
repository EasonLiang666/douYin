package handlers

import (
	"context"
	"gateway/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取用户信息
func UserInfo(ginCtx *gin.Context) {
	var userReq services.UserRequest
	userReq.UserId = ginCtx.Query("user_id")
	//PanicIfUserError(ginCtx.Bind(&userReq))
	//从上下文中获取由token拦截器获取到的userId也赋值给context上下文
	strUserId := ginCtx.GetString("userId")
	ctx := context.WithValue(ginCtx.Request.Context(), "userId", strUserId)
	ginCtx.Request = ginCtx.Request.WithContext(ctx)
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserInfo(ginCtx.Request.Context(), &userReq)
	PanicIfUserError(err)
	//ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": userResp.StatusCode,
		"status_msg":  userResp.StatusMsg,
		"user":        userResp.User,
	})
}

// 用户注册
func UserRegister(ginCtx *gin.Context) {
	var userRegisterReq services.UserRegisterRequest
	username := ginCtx.Query("username")
	password := ginCtx.Query("password")
	userRegisterReq.Password = password
	userRegisterReq.Username = username
	//PanicIfUserError(ginCtx.Bind(&userRegisterReq))
	//从上下文中获取由token拦截器获取到的userId也赋值给context上下文
	strUserId := ginCtx.GetString("userId")
	ctx := context.WithValue(ginCtx.Request.Context(), "userId", strUserId)
	ginCtx.Request = ginCtx.Request.WithContext(ctx)
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserRegister(ginCtx.Request.Context(), &userRegisterReq)
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
	var userLoginReq services.UserLoginRequest
	userLoginReq.Username = ginCtx.Query("username")
	userLoginReq.Password = ginCtx.Query("password")
	//PanicIfUserError(ginCtx.Bind(&userLoginReq))
	//从上下文中获取由token拦截器获取到的userId也赋值给context上下文
	strUserId := ginCtx.GetString("userId")
	ctx := context.WithValue(ginCtx.Request.Context(), "userId", strUserId)
	ginCtx.Request = ginCtx.Request.WithContext(ctx)
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserLogin(ginCtx.Request.Context(), &userLoginReq)
	PanicIfUserError(err)
	//ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": userResp.StatusCode,
		"status_msg":  userResp.StatusMsg,
		"user_id":     userResp.UserId,
		"token":       userResp.Token,
	})

}

//消息记录
func MessageChatRecord(ginCtx *gin.Context) {
	var messageChatRequest services.MessageChatRequest
	messageChatRequest.ToUserId = ginCtx.Query("to_user_id")
	messageChatRequest.Token = ginCtx.Query("token")
	//PanicIfUserError(ginCtx.Bind(&messageChatRequest))
	//从上下文中获取由token拦截器获取到的userId也赋值给context上下文
	strUserId := ginCtx.GetString("userId")
	ctx := context.WithValue(ginCtx.Request.Context(), "userId", strUserId)
	ginCtx.Request = ginCtx.Request.WithContext(ctx)
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.MessageChatRecord(ginCtx.Request.Context(), &messageChatRequest)
	PanicIfUserError(err)
	//ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code":  userResp.StatusCode,
		"status_msg":   userResp.StatusMsg,
		"message_list": userResp.MessageList,
	})

}

//发送消息
func SendMessage(ginCtx *gin.Context) {
	var relationMessageReq services.RelationMessageRequest
	relationMessageReq.Content = ginCtx.Query("content")
	relationMessageReq.ToUserId = ginCtx.Query("to_user_id")
	relationMessageReq.ActionType = ginCtx.Query("action_type")
	relationMessageReq.Token = ginCtx.Query("token")
	//PanicIfUserError(ginCtx.Bind(&relationMessageReq))
	//从上下文中获取由token拦截器获取到的userId也赋值给context上下文token
	strUserId := ginCtx.GetString("userId")
	ctx := context.WithValue(ginCtx.Request.Context(), "userId", strUserId)
	ginCtx.Request = ginCtx.Request.WithContext(ctx)
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.SendMessage(ginCtx.Request.Context(), &relationMessageReq)

	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": userResp.StatusCode,
		"status_msg":  userResp.StatusMsg,
	})
}

//关注/取关
func RelationAction(ginCtx *gin.Context) {
	var relationActionRequest services.RelationActionRequest
	PanicIfUserError(ginCtx.Bind(&relationActionRequest))
	userService := ginCtx.Keys["userService"].(services.UserService)

	userResp, err := userService.RelationAction(context.Background(), &relationActionRequest)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": userResp.StatusCode,
		"status_msg":  userResp.StatusMsg,
	})
}

//关注列表
func FollowList(ginCtx *gin.Context) {
	var relationFollowListRequest services.RelationFollowListRequest
	PanicIfUserError(ginCtx.Bind(&relationFollowListRequest))
	userService := ginCtx.Keys["userService"].(services.UserService)

	userResp, err := userService.FollowList(context.Background(), &relationFollowListRequest)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": userResp.StatusCode,
		"status_msg":  userResp.StatusMsg,
		"user_list":   userResp.UserList,
	})

}

//粉丝列表
func FollowerList(ginCtx *gin.Context) {
	var relationFollowerListRequest services.RelationFollowerListRequest
	PanicIfUserError(ginCtx.Bind(&relationFollowerListRequest))
	userService := ginCtx.Keys["userService"].(services.UserService)

	userResp, err := userService.FollowerList(context.Background(), &relationFollowerListRequest)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": userResp.StatusCode,
		"status_msg":  userResp.StatusMsg,
		"user_list":   userResp.UserList,
	})
}

//朋友列表
func FriendList(ginCtx *gin.Context) {
	var relationFriendListRequest services.RelationFriendListRequest
	PanicIfUserError(ginCtx.Bind(&relationFriendListRequest))
	userService := ginCtx.Keys["userService"].(services.UserService)

	userResp, err := userService.FriendList(context.Background(), &relationFriendListRequest)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": userResp.StatusCode,
		"status_msg":  userResp.StatusMsg,
		"user_list":   userResp.UserList,
	})
}
