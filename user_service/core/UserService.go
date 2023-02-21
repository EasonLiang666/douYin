package core

import (
	"context"
	"math/rand"
	"strconv"
	"time"
	"user_service/model"
	"user_service/services"
	"user_service/utils"
)

//用户信息
func (*UserService) UserInfo(ctx context.Context, req *services.UserRequest, resp *services.UserResponse) error {

	var users model.NewUser
	var user1 = resp.User
	model.DB.Where("Id = ? AND Token = ?", req.UserId, req.Token).Find(&users)
	if users.Name != "" {
		model.DB.Where("Id = ? ", req.UserId).Find(&user1)
		resp.StatusCode = 0
		resp.User = user1
		resp.StatusMsg = "ok"
	}
	return nil
}

//用户登录
func (*UserService) UserLogin(ctx context.Context, req *services.UserLoginRequest, resp *services.UserLoginResponse) error {
	var user model.NewUser
	model.DB.Where(" Password =? AND Name = ? ", req.Password, req.Username).Find(&user)
	if user.Id != 0 {
		resp.StatusCode = 0
		resp.StatusMsg = "ok"
		resp.UserId = user.Id
		resp.Token = user.Token
	}
	return nil
}

//用户注册
func (*UserService) UserRegister(ctx context.Context, req *services.UserRegisterRequest, resp *services.UserRegisterResponse) error {
	// 插入新用户记录
	// 以时间作为种子，确保每次运行的结果不同
	rand.Seed(time.Now().UnixNano())
	// 生成一个介于 0 和 100 之间的不重复的随机数序列
	nums := rand.Intn(101)
	user := &model.User{
		Id:            int64(nums),
		Name:          req.Username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}

	err2 := model.DB.Create(user)
	if err2 != nil {
		resp.StatusCode = 0
		resp.StatusMsg = "ok"
		resp.UserId = int64(nums)
		var int, err1 = strconv.Atoi(string(resp.UserId))
		if err1 != nil {
			panic(err1)
		}
		token, _ := utils.GenerateToken(uint(int))
		resp.Token = token
		user1 := &model.NewUser{
			Id:       int64(nums),
			Name:     req.Username,
			Password: req.Password,
			Token:    token,
		}
		_ = model.DB.Create(user1)
	} else {
		resp.StatusCode = 1
		resp.StatusMsg = "error"
	}
	return nil
}

//消息记录
func (*UserService) MessageChatRecord(ctx context.Context, req *services.MessageChatRequest, resp *services.MessageChatResponse) error {
	var user model.NewUser
	model.DB.Where("Token = ?", req.Token).Find(&user)
	if user.Id != 0 {

		rows, _ := model.DB.Where(" ToUserId = ? AND FromUserId = ?", req.ToUserId, user.Id).Rows()
		for rows.Next() {
			rows.Scan(&resp.MessageList)
		}
		resp.StatusCode = 0
		resp.StatusMsg = "ok"
	}
	return nil
}

//发送消息
func (*UserService) SendMessage(ctx context.Context, req *services.RelationMessageRequest, resp *services.RelationMessageResponse) error {
	rand.Seed(time.Now().UnixNano())
	// 生成一个介于 0 和 100 之间的不重复的随机数序列
	nums := rand.Intn(1001)
	var user model.NewUser
	model.DB.Where("Token = ?", req.Token).Find(&user)
	if user.Id != 0 {
		mess := &model.Message{
			Id:         int64(nums),
			ToUserId:   req.ToUserId,        // 该消息接收者的id
			FromUserId: user.Id,             // 该消息发送者的id
			Content:    req.Content,         // 消息内容
			CreateTime: time.Now().String(), // 消息创建时间
		}
		err := model.DB.Create(mess)
		if err != nil {
			resp.StatusCode = 0
			resp.StatusMsg = "ok"

		} else {
			resp.StatusCode = 1
			resp.StatusMsg = "error"
		}
	}
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
