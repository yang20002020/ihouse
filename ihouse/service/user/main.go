package main

import (
	"bj38web_084/service/user/handler"
	"bj38web_084/service/user/model"
	user "bj38web_084/service/user/proto/user"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/micro/go-micro/util/log"
)

func main() {
	//初始化mysql 连接池
	model.InitDb()
	//redis 连接池 初始化
	model.InitRedis()
	//初始化consul
	consulReg := consul.NewRegistry()
	// New Service  --指定consul
	service := micro.NewService(
		micro.Address("192.168.63.128:12342"), //防止随机生成port
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
