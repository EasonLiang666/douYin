package core

import (
	"context"
	"douYin/user_service/services"
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
	return nil
}
func (*UserService) MessageChatRecord(ctx context.Context, req *services.UserRegisterRequest, resp *services.UserRegisterResponse) error {
	return nil
}
func (*UserService) SendMessage(ctx context.Context, req *services.UserRegisterRequest, resp *services.UserRegisterResponse) error {
	return nil
}
func (*UserService) RelationAction(ctx context.Context, req *services.UserRegisterRequest, resp *services.UserRegisterResponse) error {
	return nil
}
func (*UserService) FollowList(ctx context.Context, req *services.UserRegisterRequest, resp *services.UserRegisterResponse) error {
	return nil
}
func (*UserService) FollowerList(ctx context.Context, req *services.UserRegisterRequest, resp *services.UserRegisterResponse) error {
	return nil
}
func (*UserService) FriendList(ctx context.Context, req *services.UserRegisterRequest, resp *services.UserRegisterResponse) error {
	return nil
}
