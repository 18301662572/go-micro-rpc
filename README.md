# go-micro微服务架构
## go-micro-rpc 项目

### micro git仓库:https://github.com/micro 
### micro文档： https://micro.mu/docs/api.html
### orm框架：https://gorm.io/zh_CN/docs/index.html
### 验证第三方库：https://github.com/go-playground/validator
### 验证第三方库文档：https://github.com/go-playground/validator/blob/v9/_examples/simple/main.go

```text
1.go-micro框架
2.go-plugins 插件
3.micro 工具集(toolkit)
4.validate 验证库: 实现了自定义验证，正则验证，dive（对切片，map,数组每个元素验证）
```

### 解析图
```text
images/gprc-api.png
images/gprc-api2.png
```

#### 总结 （go-micro-http的HTTP REST和HTTP API ， go-mico-rpc）
```text
1.使用go-micro框架：利用http api,rpc api的方式实现服务的编写，调用，服务的注册和发现
2.使用go-plugins的部分功能，实现中间件，熔断器
3.micro工具集(toolkit)，可以认为是一个微服务工具包，帮助完成 API Gateway(网关)，cli(客户端命令)，Sidecar(异构语言的接入)，
模板生成，Web管理界面
```

#### 1.架构
```text
ServiceImpl                    对服务接口的具体实现
    ProdServices.go            实现商品服务
Services
    protos
        Models.proto           商品模型
        ProdService.proto      服务接口
main.go                        服务设置，注册服务
gen.bat                        .proto生成.pb.go/micro.go的命令
apigw.bat                       使用micro api工具生成服务网关的命令
db.bat                          使用micro web运行web页面,也可以使用etcd
grpc
    grpcGW    
        main.go                 grpc gateway 网关服务的注册
    grpcGWproto                 grpc gateway 网关生成go文件
    proto                       .proto文件， 生成的 .pb.go / micro.go文件
    gengw.bat                   生成grpc网关文件及.pb.go/micro.go文件的命令
    main.go                     grpc服务的注册和实现                            
```

#### 2.安装插件
```text
1.从 go-micro-http gomod拷过来
2.安装第三方插件,生成protoc-go-inject-tag.exe工具
    处理(.proto文件)参数模型中的json tag不一致的问题(或者加vaild验证,form,uri)
    http://github.com/favadi/protoc-go-inject-tag
    go get -u github.com/favadi/protoc-go-inject-tag
2.go-micro的装饰器wrapper（中间件/拦截器）
    写在client端，
    目的：在执行目标函数之前做一些事情，比如写日志等...
    https://github.com/micro/go-plugins/tree/master/wrapper
3.熔断器：服务容错（异常，超时处理） 
    hystrix：提供服务的熔断、降级、隔离等。
    https://github.com/afex/hystrix-go    
4.安装 micro工具集
    go get github.com/micro/micro
5.安装orm 框架
    go get -u github.com/jinzhu/gorm        
6.安装验证第三方库 （v9）
    go get gopkg.in/go-playground/validator.v9  
    import "gopkg.in/go-playground/validator.v9"
```

#### 安装出现的问题
```text
1.如果版本不一致 
    replace google.golang.org/grpc v1.28.1 => google.golang.org/grpc v1.25.1
    如果本地没有 v1.25.1版本
    cmd 执行 replace google.golang.org/grpc v1.28.1 => google.golang.org/grpc v1.25.1，就可以下载下来     
2.最新版本的包下的某个文件下载不下来，手动下载放进去（有风险）
  或者 go get github.com/micro/go-micro@master 下载
```

#### 3.gen.bat 文件
```text
gen.bat                 
    proto文件生成go文件的命令
         1.cd 进入文件夹下
         2.protoc使用插件执行.proto文件，将生成的go文件存放在Services目录下（../表示上级目录）
         3.运行protoc-go-inject-tag 生成的文件.pb.go,改写.pb.go对应的json字段
         4.cd .. & cd .. 回到当前项目文件夹下（go-micro-pro）
         5.执行，终端直接执行gen.bat
注：go_out 生成的是 .pb.go文件 （模型结构体文件）                                   --message
    micro_out 生成的是 .micro.go文件(服务接口文件interface和服务注册Handler)        --service
```

#### 4.服务调用流程
```text
HTTP REST （go-micro-http）----> HTTP API（go-micro-http） --->go-micro-rpc
http 客户端访问---> go-micro client端服务--->go-micro-rpc serivices服务端
```

#### 5.micro 工具查看和调用grpc服务
```text
1.查看注册到consul中的服务列表
micro --registry consul --registry_address localhost:8500 list services
2.查看 prodservice 服务的信息
micro --registry consul --registry_address localhost:8500 get service prodservrice
micro --registry consul --registry_address localhost:8500 get service api.wj.com.prodservrice
3.调用 prodservrice 服务的GetProdList方法(方法名是服务信息里的Endpoint)
micro --registry consul --registry_address localhost:8500 call prodservrice ProdService.GetProdList "{\"size\":5}"
micro --registry consul --registry_address localhost:8500 call api.wj.com.prodservrice ProdService.GetProdList "{\"size\":2}"
```

#### 6.micro为rpc服务创建网关，访问rpc服务 (apigw.bat)
```text
文档：https://micro.mu/docs/api.html
基本命令：
    micro api -h
1.创建 apigw.bat
2.运行 apigw.bat,会开启HTTP :8080 监听端口(网关地址,默认:8080)
3.postman 操作
   1.POST http://localhost:8080/prodservrice/ProdService/GetProdList {"size":3}
    POST: 访问方式
    prodservrice:服务名（去掉MICRO_API_NAMESPACE后的）
    ProdService/GetProdList： 服务对应的EndPoint 
    {"size":3} ： 需要传的参数

   2.POST http://localhost:8080/prodservrice/UserService/UserReg
```

#### 7.grpc网关的基本设置和运行--grpc-gateway（go-micro-grpc/grpc）
```text
文档：https://github.com/grpc-ecosystem/grpc-gateway
1.安装
    go get -u  google.golang.org/grpc

    go get -u  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway 
    go get -u  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger 
    go get -u  github.com/golang/protobuf/protoc-gen-go
2. .proto文件的service 加 option 访问方式
3. 执行 gengw.bat
        生成 grpc 服务使用的文件 .pb.go / micro.go
        生成 grpc gateway 网关使用的文件 .pb.go / .gw.go
        生成后，修改包名，包名不能一样，否则会出现调用问题
4.grpcGW/main.go : grpc gateway 网关,将服务注册到网关中，
5.开启grpc服务，开启grpc gateway网关服务，监听http端口，就可通过http访问
6.postman 访问： GET http://localhost:9000/test/33
```


