package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"user_service/config"
	"user_service/core"
	"user_service/services"
)

func main() {
	config.Init()
	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"), // 127.0.0.1:2379 是 etcd 服务注册中心的默认地址
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"), // 微服务名字
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	// 启动微服务
	_ = microService.Run()
}
