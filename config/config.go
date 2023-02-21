package config

import "time"

// HostSSH SSH配置
//const HostSSH = "43.138.25.60"
//const UserSSH = "ftpuser"
//const PasswordSSH = "424193726"
//const TypeSSH = "password"
//const PortSSH = 22
//const MaxMsgCount = 100
//const SSHHeartbeatTime = 10 * 60

var ExpireTime = time.Hour * 48 // 设置Redis数据热度消散时间。

// VideoCount 每次获取视频流的数量
const VideoCount = 5

const ValidComment = 0   //评论状态：有效
const InvalidComment = 1 //评论状态：取消

const IsLike = 0     //点赞的状态
const Unlike = 1     //取消赞的状态
const LikeAction = 1 //点赞的行为
const Attempts = 3   //操作数据库的最大尝试次数

const DefaultRedisValue = -1 //redis中key对应的预设值，防脏读
