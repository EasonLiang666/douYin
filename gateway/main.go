package main

import (
	"gateway/services"
	"gateway/weblib"
	"gateway/wrappers"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"time"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// User用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	// 用户服务调用实例
	userService := services.NewUserService("rpcUserService", userMicroService.Client())

	// Video视频
	videoMicroService := micro.NewService(
		micro.Name("videoService.client"),
		micro.WrapClient(wrappers.NewVideoWrapper),
	)
	videoService := services.NewVideoService("rpcTaskService", videoMicroService.Client())

	//创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),

		//将服务调用实例使用gin处理
		web.Handler(weblib.NewRouter(userService, videoService)),

		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
