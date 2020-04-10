package AppLib

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"regexp"
)

//自定义的正则验证
func AddRegexTag(tagName string,patten string,v *validator.Validate)error{
	return v.RegisterValidation(tagName, func(fl validator.FieldLevel) bool {
		m,_:=regexp.MatchString(patten,fl.Field().String())
		return m
	},false)
}

//封装好的验证
func ValidErrMsg(obj interface{},err error) error{
	getObj:= reflect.TypeOf(obj)
	if err!=nil{
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				if f, exsit := getObj.Elem().FieldByName(e.Field()); exsit {
					//只要有一个error,立刻返回。
					if value,ok:=f.Tag.Lookup("vmsg");ok{
						return fmt.Errorf("%s",value)
					}else{
						return fmt.Errorf("%s",e)
					}
				}else{
					return fmt.Errorf("%s",e)
				}
			}
		}
	}
	return err
}

//通过反射的方式获取类型
func GetValidMsg(obj interface{},field string){
	getObj:= reflect.TypeOf(obj)
	if f,exsit:=getObj.Elem().FieldByName(field); exsit{
		fmt.Println(f.Tag.Get("vmsg"))
	}
}
