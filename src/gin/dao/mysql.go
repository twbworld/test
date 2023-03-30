package dao

import (
	"fmt"
	"test/src/gin/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitMysql() (err error) {
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:tp@tcp(172.1.1.97:3306)/gva?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return
}


func InitModel() (err error) {
	if err = DB.AutoMigrate(&model.SysUser{}); err != nil {
		fmt.Println(err)
	}
	return
}
