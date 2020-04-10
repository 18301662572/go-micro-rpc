#运行 apigw.bat,会开启HTTP :8080 监听端口(网关地址,默认:8080),就可通过POSTMAN访问
#使用micro api工具生成服务网关的命令
set MICRO_REGISTRY=consul
set MICRO_REGISTRY_ADDRESS=localhost:8500
set MICRO_API_NAMESPACE=api.wj.com
set MICRO_API_HANDLER=rpc
micro api