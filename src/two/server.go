package main

import (
    "fmt"
    "net"
)

type IpPort struct{
    Ip string
    Port int
}

func Csh(i string, p int) (*IpPort)  {
    server := &IpPort{
        Ip: i,
        Port: p,
    }

    return server
}


func (this *IpPort)Start()  {
    lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

    if err != nil {
        fmt.Println("连接失败")
    }

    for {
        com, err := lis.Accept()

        if err != nil {
            fmt.Println("不存在连接")
            continue
        }

        go this.Handle(com)
    }
}

func (this *IpPort) Handle(com net.Conn)  {
    fmt.Println("连接成功")
}
