package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"strings"
	"time"
	"video_service/model"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("video_service/config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadMysqlData(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

//HostSSH SSH配置
const HostSSH = "43.138.25.60"
const UserSSH = "ftpuser"
const PasswordSSH = "424193726"
const TypeSSH = "password"
const PortSSH = 22
const MaxMsgCount = 100
const SSHHeartbeatTime = 10 * 60
const Secret = "TikTok"

// OneDayOfHours 时间
var OneDayOfHours = 60 * 60 * 24
var OneMinute = 60 * 1
var OneMonth = 60 * 60 * 24 * 30
var OneYear = 365 * 60 * 60 * 24

// ConConfig ftp服务器地址
const ConConfig = "43.138.25.60:21"
const FtpUser = "ftpuser"
const FtpPsw = "424193726"
const HeartbeatTime = 2 * 60

// PlayUrlPrefix 存储的图片和视频的链接
const PlayUrlPrefix = "http://43.138.25.60/"
const CoverUrlPrefix = "http://43.138.25.60/images/"

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
