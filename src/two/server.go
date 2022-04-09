package main

import (
	"fmt"
	"net"
    "sync"
)

type IpPort struct {
	Ip   string
	Port int

    OnLiveMap map[string] *User //在线用户列表
    Message chan string //广播channel

    mapLock sync.RWMutex //用于锁进程
}


func main()  {

    i, p := "127.0.0.1", 7777

    server := &IpPort{
		Ip:   i,
		Port: p,
        OnLiveMap: make(map[string]*User),
        Message: make(chan string),
	}

    server.Start()
}

func (this *IpPort) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Println("连接失败")
        return
	}

    defer lis.Close()


    go this.listenMessage()

	for {
		com, err := lis.Accept() //会在此柱塞

		if err != nil {
			fmt.Println("不存在连接")
			continue
		}

		go this.Handle(com)
	}
}

func (this *IpPort) listenMessage(){
    for{
        msg := <-this.Message

        this.mapLock.Lock()
        for _, user := range this.OnLiveMap{
            user.C <- msg
        }
        this.mapLock.Unlock()
    }
}

func (this *IpPort) Handle(com net.Conn) {
	fmt.Println("连接成功")

    user := NewUser(com)

    //加入到用户列表中
    this.mapLock.Lock()
    this.OnLiveMap[user.Name] = user
    this.mapLock.Unlock()

    //加入到广播
    this.Message <- "[" + user.Name + "]上线了"

}
