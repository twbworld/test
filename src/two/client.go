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
    }

    res := conf.NewClient()

    if nil == res {
        fmt.Println("连接失败")
        return
    }

    fmt.Println("连接成功")
    select{}
}
