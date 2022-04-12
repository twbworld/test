package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Conf struct {
	Ip   string
	Port int

	List map[string]*User //在线用户列表
	Msg  chan string      //广播channel

	mapLock sync.RWMutex //用于锁进程
}

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func main() {

	i, p := "127.0.0.1", 9999

	server := &Conf{
		Ip:   i,
		Port: p,
		List: make(map[string]*User),
		Msg:  make(chan string),
	}

	server.Start()
}

func (this *Conf) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Println("连接失败")
		return
	}

	defer lis.Close()

	go this.ListenMsg()

	for {
		com, err := lis.Accept() //会在此柱塞

		if err != nil {
			fmt.Println("不存在连接")
			continue
		}

		go this.Handle(com)
	}
}


func (this *Conf) ListenMsg() {
	for {
		msg := <-this.Msg

		this.mapLock.Lock()
		for _, user := range this.List {
			user.C <- msg //广播
		}
		this.mapLock.Unlock()
	}
}

//监听当前User channel的方法, 一旦有消息, 就直接发送给对端客户端
func (this *User) ListenC() {

	for {
		msg := <-this.C

		//给客户端写入
		this.conn.Write([]byte(msg + "\n"))
	}

}

func (this *Conf) ListenWrice(user *User){
    for {
        by := make([]byte, 10240)
        n, err := user.conn.Read(by)
        if nil != err && io.EOF != err  {
            fmt.Println("消息出错")
            continue
        }else if 0 == n {
            this.Msg <- "[下线]" + user.Name
            return
        }

        this.Msg <- "[" + user.Name + "]" + string(by[:n-1]) //提取消息并去除"\n"
    }
}

func (this *Conf) Handle(com net.Conn) {
	fmt.Println("连接成功")


    //新连接,创建用户
    ipStr := com.RemoteAddr().String()
	user := &User{
		Name: "用户" + ipStr,
		Addr: ipStr,
		C:    make(chan string),
		conn: com,
	}

	//启动监听当前user channel消息的goroutine
	go user.ListenC()



	//加入到用户列表中
	this.mapLock.Lock()
	this.List[user.Name] = user
	this.mapLock.Unlock()

	//加入到广播
	this.Msg <- "[上线]" + user.Name

    go this.ListenWrice(user)

}
