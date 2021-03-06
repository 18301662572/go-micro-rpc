package DBModels

import "time"

//数据库模型
type Users struct {
	UserId int `gorm:"column:user_id;AUTO_INCREMENT;PRIMARY_KEY"`
	UserName string `gorm:"column:user_name;type:varchar(100);UNIQUE_INDEX"`
	UserPwd string `gorm:"column:user_pwd;type:varchar(100);"`
	UserDate time.Time	`gorm:"column:user_date;"`
}
