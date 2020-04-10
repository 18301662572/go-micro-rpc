package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	testGW "go-micro-grpc/grpc/grpcGWproto"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

//grpc gateway 网关

func main(){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	//创建路由处理器--对应.proto文件设置的 option
	mux := runtime.NewServeMux()

	//创建一个[]grpc.DialOption
	opts := []grpc.DialOption{grpc.WithInsecure()}
	//grpc服务地址
	gRpcEndPoint:="localhost:8012"

	//RestService服务相关的REST接口中转到后面的GRPC服务
	err := testGW.RegisterTestServiceHandlerFromEndpoint(
		ctx, mux, gRpcEndPoint,
		opts,
	)
	if err != nil {
		log.Fatal(err)
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":9000", mux)
}
