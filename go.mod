module go-micro-grpc

go 1.12

require (
	github.com/denisenkom/go-mssqldb v0.0.0-20200206145737-bbfc9a55622e // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/mock v1.4.3 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.3
	github.com/hashicorp/go-rootcerts v1.0.1 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/memberlist v0.1.5 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/micro/go-micro v1.18.0
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/pascaldekloe/goe v0.1.0 // indirect
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71 // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/sys v0.0.0-20200409092240-59c9f1ba88fa // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20200410040751-3bd20875a2eb // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20200408120641-fbb3ad325eb7
	google.golang.org/grpc v1.28.1
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/resty.v1 v1.12.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	github.com/micro/go-micro v1.18.0 => github.com/micro/go-micro v1.13.2 //替换版本号，1.18.0没有consul,selector包
	google.golang.org/grpc v1.28.1 => google.golang.org/grpc v1.25.1
)
