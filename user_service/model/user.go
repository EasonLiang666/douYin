package model

type User struct {
	Id            int64  `gorm:"unique"` // 用户id
	Name          string // 用户名称
	FollowCount   int64  // 关注总数
	FollowerCount int64  // 粉丝总数
	IsFollow      bool   // true-已关注，false-未关注
}
