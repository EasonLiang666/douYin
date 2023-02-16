package core

import (
	"context"
	"douYin/user_service/services"
	"douYin/user_service/utils"
	"strconv"
)

//用户信息
func (*UserService) UserInfo(ctx context.Context, req *services.UserRequest, resp *services.UserResponse) error {
	return nil
}

//用户登录
func (*UserService) UserLogin(ctx context.Context, req *services.UserLoginRequest, resp *services.UserLoginResponse) error {
	return nil
}

//用户注册
func (*UserService) UserRegister(ctx context.Context, req *services.UserRegisterRequest, resp *services.UserRegisterResponse) error {
	resp.StatusCode = 1
	resp.StatusMsg = "ok"
	resp.UserId = 1
	int, err := strconv.Atoi("123")
	if err != nil {
		panic(err)
	}
	token, err := utils.GenerateToken(uint(int))
	resp.Token = token

	return nil
}

//消息记录
func (*UserService) MessageChatRecord(ctx context.Context, req *services.MessageChatRequest, resp *services.MessageChatResponse) error {
	return nil
}

//发送消息
func (*UserService) SendMessage(ctx context.Context, req *services.RelationMessageRequest, resp *services.RelationMessageResponse) error {
	return nil
}

//关注/取关
func (*UserService) RelationAction(ctx context.Context, req *services.RelationActionRequest, resp *services.RelationActionResponse) error {
	return nil
}

//关注列表
func (*UserService) FollowList(ctx context.Context, req *services.RelationFollowListRequest, resp *services.RelationFollowListResponse) error {
	return nil
}

//粉丝列表
func (*UserService) FollowerList(ctx context.Context, req *services.RelationFollowerListRequest, resp *services.RelationFollowerListResponse) error {
	return nil
}

//朋友列表
func (*UserService) FriendList(ctx context.Context, req *services.RelationFriendListRequest, resp *services.RelationFriendListResponse) error {
	return nil
}
