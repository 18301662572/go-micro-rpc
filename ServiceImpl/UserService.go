package ServiceImpl

import (
	"context"
	"go-micro-grpc/AppInit"
	"go-micro-grpc/DBModels"
	"go-micro-grpc/Services"
	"time"
)

//UserService 实现
type UserService struct {

}

func (u *UserService) UserReg(c context.Context,req *Services.UserModel,rsp *Services.RegReponse) error{
	user:=DBModels.Users{
		UserName:req.UserName,
		UserPwd:req.UserPwd,
		UserDate:time.Now(),
	}
	if err:=AppInit.GetDB().Create(&user).Error;err!=nil{
		rsp.Message=err.Error()
		rsp.Status="error"
	}else{
		rsp.Message=""
		rsp.Status="success"
	}
	return nil
}