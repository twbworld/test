package lib1

import (
	"fmt"
)

func init(){
    fmt.Println("这是init")
}


//方法名小写开头是私有方法
func Index(){
	fmt.Println("这是index")
}
