#使用micro web运行web页面,也可以使用etcd
#set MICRO_REGISTRY=etcd
set MICRO_REGISTRY=consul
#set MICRO_REGISTRY_ADDRESS=localhost:2379
set MICRO_REGISTRY_ADDRESS=localhost:8500
set MICRO_API_NAMESPACE=api.wj.com
micro web