package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := gorm.Open(mysql.Open("root:tp@tcp(172.1.1.97:3306)/gva"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// adapter, _ := gormadapter.NewAdapter("mysql", "root:tp@tcp(172.1.1.97:3306)/gva", true)
	adapter, err := gormadapter.NewAdapterByDBUseTableName(db, "sys", "casbin_rule")
	if err != nil {
		fmt.Println(err)
		return
	}

	e, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		fmt.Println(err)
		return
	}

	e.AddFunction("my_func", KeyMatchFunc)

	e.LoadPolicy()


	// ok, err := e.Enforce("twb", "data2", "read")
	ok, err := e.BatchEnforce([][]interface{}{{"twb", "data", "read"},{"twb", "data3", "read"}})

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range ok{
		if v {
			// 允许alice读取data1
			fmt.Println("允许")
		} else {
			fmt.Println("拒绝")
			// 拒绝请求，抛出异常
		}
	}



}




func KeyMatchFunc(args ...interface{}) (interface{}, error) {
    name1 := args[0].(string)
    name2 := args[1].(string)

    return (bool)(KeyMatch(name1, name2)), nil
}



func KeyMatch(key1 string, key2 string) bool {
    return key1 == key2 || key1 == "data3"
}
