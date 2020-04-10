package ServiceImpl

import (
	"context"
	pb "go-micro-grpc/Services"
	"strconv"
	"time"
)

//实现商品服务

type ProdServices struct {
}

//初始化商品
func newProd(id int32,pname string) *pb.ProdModel{
	return &pb.ProdModel{
		ProdID:id,
		ProdName:pname,
	}
}


func (p *ProdServices) GetProdList(c context.Context,req *pb.ProdRequest,rsp *pb.ProdListResponse) error{
	time.Sleep(time.Second*1)
	models:=make([]*pb.ProdModel,0)
	var i int32
	for i=0;i<req.Size ;i++  {
		models=append(models,newProd(100+i,"pname"+strconv.Itoa(100+int(i))))
	}
	rsp.Data=models
	return nil
}

func (p *ProdServices) GetProdDeatil(c context.Context,req *pb.ProdRequest,rsp *pb.ProdDeatilResponse) error{
	time.Sleep(time.Second*1)
	rsp.Data=newProd(req.ProdId,"测试商品名称")
	return nil
}
