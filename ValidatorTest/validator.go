package main

import (
	"fmt"
	"go-micro-grpc/AppLib"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

//validator.v9 验证测试
//https://github.com/go-playground/validator/blob/v9/_examples/simple/main.go

//注：`validate:dive` dive：深入slice, array or map 内部的每一项进行后面tag的验证

type User struct {
	Username string `validate:"required,min=4,max=10" vmsg:"用户名必须大于4位"`
	Userpwd string `validate:"required,min=4,max=10" vmsg:"密码必须大于4位"`
	Testname string `validate:"required,abc" vmsg:"用户名规则不正确"`
	Usertags []string `validate:"required,min=1,max=5,unique,dive,usertag" vmsg:"用户标签不合法"` //dive 深入slice, array or map 内部的每一项进行后面tag的验证
}

func main(){
	u:=&User{
		Username: "aaaaa",
		Userpwd:  "121111",
		Testname:"abc23fsfsfs",
		Usertags:[]string{"1","2","3","5","722222"},
	}
	validate:=validator.New()
	//自定义验证
	AppLib.AddRegexTag("usertag","^[\u4e00-\u9fa5a-zA-Z0-9]{1,4}$",validate) //数值，字母大小写，数字{1,4}
	AppLib.AddRegexTag("abc","[a-zA-Z]\\w{5,19}",validate)//字母数字下划线{5,19}

	//用户名密码封装后的验证
	err:=AppLib.ValidErrMsg(u,validate.Struct(u))
	if err!=nil{
		log.Fatal(err)
	}

	//同上，用户名密码封装后的验证
	//err := validate.Struct(u)
	//if err!=nil{
	//	if errs, ok := err.(validator.ValidationErrors); ok {
	//		for _, e := range errs {
	//			//通过反射的方式获取类型
	//			AppLib.GetValidMsg(u,e.Field())
	//		}
	//	}
	//	// from here you can create your own error messages in whatever language you wish
	//	return
	//}
	fmt.Println("验证成功！")
}