package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/service/grpc"
	test "go-micro-grpc/grpc/proto"
)

//grpc服务
func main(){

	//1.初始化consul
	consulReg:=consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	//2.注册grpc service，使用grpc
	prodService:=grpc.NewService(
		micro.Name("api.wj.com.testservrice"),
		micro.Address(":8012"),
		micro.Registry(consulReg),
	)
	prodService.Init()
	//3.注册服务
	test.RegisterTestServiceHandler(prodService.Server(),new(testServiceImp))
	//4.监听服务
	prodService.Run()
}


//服务的实现
type testServiceImp struct {
}

func (t *testServiceImp) Call(ctx context.Context, in *test.TestRequest, out *test.TestResponse) error{
	out.Data=fmt.Sprintf("%s: Text!",in.Id)
	return nil
}


