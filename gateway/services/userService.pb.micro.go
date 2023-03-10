// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: userService.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	"github.com/micro/go-micro/v2/api"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	UserInfo(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error)
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserRegisterResponse, error)
	MessageChatRecord(ctx context.Context, in *MessageChatRequest, opts ...client.CallOption) (*MessageChatResponse, error)
	SendMessage(ctx context.Context, in *RelationMessageRequest, opts ...client.CallOption) (*RelationMessageResponse, error)
	RelationAction(ctx context.Context, in *RelationActionRequest, opts ...client.CallOption) (*RelationActionResponse, error)
	FollowList(ctx context.Context, in *RelationFollowListRequest, opts ...client.CallOption) (*RelationFollowListResponse, error)
	FollowerList(ctx context.Context, in *RelationFollowerListRequest, opts ...client.CallOption) (*RelationFollowerListResponse, error)
	FriendList(ctx context.Context, in *RelationFriendListRequest, opts ...client.CallOption) (*RelationFriendListResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) UserInfo(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.UserInfo", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.UserLogin", in)
	out := new(UserLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserRegisterResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.UserRegister", in)
	out := new(UserRegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) MessageChatRecord(ctx context.Context, in *MessageChatRequest, opts ...client.CallOption) (*MessageChatResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.MessageChatRecord", in)
	out := new(MessageChatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SendMessage(ctx context.Context, in *RelationMessageRequest, opts ...client.CallOption) (*RelationMessageResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.SendMessage", in)
	out := new(RelationMessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...client.CallOption) (*RelationActionResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.RelationAction", in)
	out := new(RelationActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) FollowList(ctx context.Context, in *RelationFollowListRequest, opts ...client.CallOption) (*RelationFollowListResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.FollowList", in)
	out := new(RelationFollowListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) FollowerList(ctx context.Context, in *RelationFollowerListRequest, opts ...client.CallOption) (*RelationFollowerListResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.FollowerList", in)
	out := new(RelationFollowerListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) FriendList(ctx context.Context, in *RelationFriendListRequest, opts ...client.CallOption) (*RelationFriendListResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.FriendList", in)
	out := new(RelationFriendListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	UserInfo(context.Context, *UserRequest, *UserResponse) error
	UserLogin(context.Context, *UserLoginRequest, *UserLoginResponse) error
	UserRegister(context.Context, *UserRegisterRequest, *UserRegisterResponse) error
	MessageChatRecord(context.Context, *MessageChatRequest, *MessageChatResponse) error
	SendMessage(context.Context, *RelationMessageRequest, *RelationMessageResponse) error
	RelationAction(context.Context, *RelationActionRequest, *RelationActionResponse) error
	FollowList(context.Context, *RelationFollowListRequest, *RelationFollowListResponse) error
	FollowerList(context.Context, *RelationFollowerListRequest, *RelationFollowerListResponse) error
	FriendList(context.Context, *RelationFriendListRequest, *RelationFriendListResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		UserInfo(ctx context.Context, in *UserRequest, out *UserResponse) error
		UserLogin(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error
		UserRegister(ctx context.Context, in *UserRegisterRequest, out *UserRegisterResponse) error
		MessageChatRecord(ctx context.Context, in *MessageChatRequest, out *MessageChatResponse) error
		SendMessage(ctx context.Context, in *RelationMessageRequest, out *RelationMessageResponse) error
		RelationAction(ctx context.Context, in *RelationActionRequest, out *RelationActionResponse) error
		FollowList(ctx context.Context, in *RelationFollowListRequest, out *RelationFollowListResponse) error
		FollowerList(ctx context.Context, in *RelationFollowerListRequest, out *RelationFollowerListResponse) error
		FriendList(ctx context.Context, in *RelationFriendListRequest, out *RelationFriendListResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) UserInfo(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.UserInfo(ctx, in, out)
}

func (h *userServiceHandler) UserLogin(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error {
	return h.UserServiceHandler.UserLogin(ctx, in, out)
}

func (h *userServiceHandler) UserRegister(ctx context.Context, in *UserRegisterRequest, out *UserRegisterResponse) error {
	return h.UserServiceHandler.UserRegister(ctx, in, out)
}

func (h *userServiceHandler) MessageChatRecord(ctx context.Context, in *MessageChatRequest, out *MessageChatResponse) error {
	return h.UserServiceHandler.MessageChatRecord(ctx, in, out)
}

func (h *userServiceHandler) SendMessage(ctx context.Context, in *RelationMessageRequest, out *RelationMessageResponse) error {
	return h.UserServiceHandler.SendMessage(ctx, in, out)
}

func (h *userServiceHandler) RelationAction(ctx context.Context, in *RelationActionRequest, out *RelationActionResponse) error {
	return h.UserServiceHandler.RelationAction(ctx, in, out)
}

func (h *userServiceHandler) FollowList(ctx context.Context, in *RelationFollowListRequest, out *RelationFollowListResponse) error {
	return h.UserServiceHandler.FollowList(ctx, in, out)
}

func (h *userServiceHandler) FollowerList(ctx context.Context, in *RelationFollowerListRequest, out *RelationFollowerListResponse) error {
	return h.UserServiceHandler.FollowerList(ctx, in, out)
}

func (h *userServiceHandler) FriendList(ctx context.Context, in *RelationFriendListRequest, out *RelationFriendListResponse) error {
	return h.UserServiceHandler.FriendList(ctx, in, out)
}
