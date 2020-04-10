package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	_ "go-micro-grpc/AppInit" //初始化数据库
	"go-micro-grpc/ServiceImpl"
	"go-micro-grpc/Services"
)

func main(){
	//1.初始化consul
	consulReg:=consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	//将consul迁移到etcd里面
	//etcdReg:=etcd.NewRegistry(
	//	registry.Addrs("localhost:2379"),
	//	)

	//2.注册rpc service，使用micro
	prodService:=micro.NewService(
		micro.Name("api.wj.com.prodservrice"),//micro.Name("prodservrice"),
		micro.Address(":8011"),
		micro.Registry(consulReg),//etcdReg
	)
	prodService.Init()
	//3.注册服务
	Services.RegisterProdServiceHandler(prodService.Server(),new(ServiceImpl.ProdServices))
	Services.RegisterUserServiceHandler(prodService.Server(),new(ServiceImpl.UserService))
	//4.监听服务
	prodService.Run()
}
