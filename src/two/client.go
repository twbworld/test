package main

import (
	"flag"
	"fmt"
	"net"
)

type ClientConf struct{
    Ip string
    Port int
    Name string
    flag int
    conn net.Conn
}

func (this *ClientConf)NewClient() *ClientConf{
    con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
    if nil != err {
        fmt.Println("err", err)
        return nil
    }
    this.conn = con
    return this
}

func (c *ClientConf)Run(){
    for 0 != c.flag{
        if c.ListenShell() {
            switch c.flag{
            case 1:
                fmt.Println("公聊模式")
            case 2:
                fmt.Println("私聊模式")
            case 3:
                fmt.Println("更新用户名")
            case 0:
                fmt.Println("退出")
            default:
                fmt.Println("请重新输入")
            }
        }
    }

}

func (this *ClientConf) ListenShell() bool{
    fmt.Println("1:公聊模式")
    fmt.Println("2:私聊模式")
    fmt.Println("3:更新用户名")
    fmt.Println("0:退出")

    var param int
    fmt.Scanln(&param)

    if 0 > param || 3 < param {
        return false
    }
    this.flag = param
    return true
}

var (
    Cip string
    Cport int
)

//获取命令行参数
func init(){
    flag.IntVar(&Cport,"p", 9999, "连接服务器的端口,默认是9999")
    flag.StringVar(&Cip,"i", "127.0.0.1", "连接服务器的ip,默认是127.0.0.1")
}


func main(){
    flag.Parse()

    conf := &ClientConf{
        Ip: Cip,
        Port: Cport,
        flag: 99,
    }

    res := conf.NewClient()

    if nil == res {
        fmt.Println("连接失败")
        return
    }

    fmt.Println("连接成功")

    conf.Run()
}
