package core

import (
	"context"
	"video_service/services"
)

//视频流
func (*VideoService) FeedInfo(ctx context.Context, req *services.FeedRequest, resp *services.FeedResponse) error {
	return nil
}

//发布视频
func (*VideoService) PublishAction(ctx context.Context, req *services.PublishActionRequest, resp *services.PublishActionResponse) error {
	return nil
}

//自己的发布视频列表
func (*VideoService) PublishList(ctx context.Context, req *services.PublishListRequest, resp *services.PublishListResponse) error {
	return nil
}

//点赞/取消点赞
func (*VideoService) FavoriteAction(ctx context.Context, req *services.FavoriteActionRequest, resp *services.FavoriteActionResponse) error {
	return nil
}

//点赞列表
func (*VideoService) FavoriteList(ctx context.Context, req *services.FavoriteListRequest, resp *services.FavoriteListResponse) error {
	return nil
}

//评论
func (*VideoService) CommentAction(ctx context.Context, req *services.CommentActionRequest, resp *services.CommentActionResponse) error {
	return nil
}

//评论列表
func (*VideoService) CommentList(ctx context.Context, req *services.CommentListRequest, resp *services.CommentListResponse) error {
	return nil
}
