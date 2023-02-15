package model

type Video struct {
	Id            int64  `gorm:"unique"` // 视频唯一标识
	Author        User   // 视频作者信息
	PlayUrl       string // 视频播放地址
	CoverUrl      string // 视频封面地址
	FavoriteCount int64  // 视频的点赞总数
	CommentCount  int64  // 视频的评论总数
	IsFavorite    bool   // true-已点赞，false-未点赞
	Title         string // 视频标题
}
