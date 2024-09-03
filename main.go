package main

import (
	"github.com/twbworld/dating/dao"
	"github.com/twbworld/dating/global"
	"github.com/twbworld/dating/initialize"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			global.Log.Println(p)
		}
		if dao.DB != nil {
			if err := dao.DB.Close(); err != nil {
				global.Log.Println(`数据库关闭出错[joiasjofg]`, err)
			}
		}
	}()

	initialize.Start()

}
