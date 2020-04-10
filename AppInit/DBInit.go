package AppInit

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql","root:123456@tcp(localhost:3306)/gomicro?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		log.Fatal(err)
	}
	//最大空闲数
	db.DB().SetMaxIdleConns(10)
	//最大连接数
	db.DB().SetMaxOpenConns(50)
}

func GetDB() *gorm.DB{
	return db
}
