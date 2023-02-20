package model

type NewUser struct {
	Id       int64  `gorm:"unique"` // 用户id
	Name     string // 用户名称
	Password string // 密码
	Token    string //鉴权
}
