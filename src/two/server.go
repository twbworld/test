package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"runtime"
	"strings"
	"sync"
	"time"
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

func (this *Conf) ListenWrice(user *User, isLive chan bool) {
	for {
		by := make([]byte, 10240)
		n, err := user.conn.Read(by) //客户端返回的消息
		if nil != err && io.EOF != err {
			fmt.Println("消息出错")
			return
		} else if 0 == n {
			this.Msg <- "[下线]" + user.Name
			delete(this.List, user.Name)
			return
		}

		isLive <- true //在线

		msg := string(by[:n-1]) //提取消息并去除"\n"
		fmt.Println(msg)
		if "" == msg {
			continue
		}

		if "@" == msg[:1] {
			if "@who" == msg {
				list := make([]string, 0, len(this.List))
				for k := range this.List {
					list = append(list, k)
				}
				jsonStr, err := json.Marshal(list)
				if nil == err {
					msg = string(jsonStr)
					user.C <- msg
					continue
				} else {
					fmt.Println(err)
				}
			} else if 8 < len(msg) && "@rename=" == msg[:8] {

				// newName := msg[10:]
				newName := strings.Split(msg, "=")[1]

				this.mapLock.Lock()
				if _, ok := this.List[newName]; ok {
					user.C <- "用户名[" + newName + "]已存在"
					this.mapLock.Unlock()
					continue
				}
				delete(this.List, user.Name)
				user.Name = newName
				this.List[newName] = user
				this.mapLock.Unlock()
				user.C <- "修改用户名成功:" + newName
				continue
			} else {
                //私聊
				wz := strings.Index(msg, "=")
				if wz > 1 {
					//私聊;格式@张三=你好
					username := msg[1:wz]
					msg = msg[wz+1:]
					if _, ok := this.List[username]; ok && username != user.Name {
						this.List[username].C <- "[" + user.Name + "]" + msg
						continue
					}
				}
                user.C <- "输入错误, 请重新输入"
			}
		}else {
            this.Msg <- "[" + user.Name + "]" + msg //广播群发
        }

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

	var isLive chan bool
	isLive = make(chan bool)

	//监听客户端回消息
	go this.ListenWrice(user, isLive)

	//超时踢出
	for {
		//select 会循环检测条件,如果有满足则执行并退出,否则一直循环检测; 所以外侧要使用for
		select {
		case <-isLive:

		case <-time.After(time.Second * 60 * 60):
			user.C <- "您已超时被强踢"
			time.Sleep(time.Second * 1)
			close(user.C)
			user.conn.Close()
			delete(this.List, user.Name)
			this.Msg <- user.Name + " 超时被强踢"
			fmt.Println(user.Name + " 超时被强踢")
			runtime.Goexit()
			return
		}
	}

}
