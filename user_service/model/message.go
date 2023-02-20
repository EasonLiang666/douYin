package model

type Message struct {
	Id         int64  `gorm:"unique"` // 消息id
	ToUserId   int64  // 该消息接收者的id
	FromUserId int64  // 该消息发送者的id
	Content    string // 消息内容
	CreateTime string // 消息创建时间
}
