package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
)

type ClientConf struct {
	Ip   string
	Port int
	flag int
	conn net.Conn
}

//连接服务端
func (this *ClientConf) NewClient() *ClientConf {
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if nil != err {
		fmt.Println("err", err)
		return nil
	}
	this.conn = con
	return this
}


func (c *ClientConf) ListenServer(){
    io.Copy(os.Stdout, c.conn) //获取服务费返回, 并输出到界面
    // for{
    //     res := make([]byte, 10240)
    //     n2, err2 := c.conn.Read(res)
    //     if nil != err2 && err2 != io.EOF {
    //         fmt.Println("连接错误[fsdiojfsad]")
    //         return
    //     }else if 0 == n2  {
    //         fmt.Println("对方已断开连接[gpoergv]")
    //         return
    //     }else {
    //         fmt.Println(string(res))
    //         return
    //     }
    // }
}

func (c *ClientConf) ReName() {
    fmt.Println("请输入要修改的名称 :")
    var newName string
    fmt.Scanln(&newName)
    n, err := c.conn.Write([]byte("@rename=" + newName + "\n"))
    if nil != err && err != io.EOF {
        fmt.Println("连接错误[fsdqawg]")
        return
    }else if 0 == n  {
        fmt.Println("对方已断开连接[oldrtasb]")
        return
    }
}

func (c *ClientConf) PublicChat() {
    fmt.Println("已进入公聊模式(输入exit退出)")

    var res string
    for {
        fmt.Scanln(&res)
        if "exit" == res {
            return
        }
        if 0 != len(res) {
            n, err := c.conn.Write([]byte(res + "\n"))
            if nil != err && err != io.EOF {
                fmt.Println("连接错误[gdssdfsadf]")
                return
            }else if 0 == n  {
                fmt.Println("对方已断开连接[gdfdasdas]")
                return
            }
        }
        res = ""
    }

}

func (c *ClientConf) PrivateChat() {
    fmt.Println("已进入私聊模式(输入exit退出)")

    for{
        fmt.Println("请输入聊天对象")

        n, err := c.conn.Write([]byte("@who" + "\n"))
        if nil != err && err != io.EOF {
            fmt.Println("连接错误[rwrwerwe]")
            return
        }else if 0 == n  {
            fmt.Println("对方已断开连接[rwerewrwe]")
            return
        }

        var he string
        fmt.Scanln(&he)

        if "" == he {
            fmt.Println("不能为空,请重新输入")
            continue
        }
        if "exit" == he {
            return
        }

        fmt.Println("开始聊天")
        for {
            var chat string
            fmt.Scanln(&chat)

            if "" == chat {
                fmt.Println("不能为空,请重新输入")
                continue
            }
            if "exit" == chat {
                break
            }

            n, err := c.conn.Write([]byte("@" + he + "=" + chat + "\n"))
            if nil != err && err != io.EOF {
                fmt.Println("连接错误[rwrwerwe]")
                return
            }else if 0 == n  {
                fmt.Println("对方已断开连接[rwerewrwe]")
                return
            }
        }
    }
}

func (c *ClientConf) ExitOut() {
    fmt.Println("成功退出")
    runtime.Goexit()
    c.conn.Close()
}

func (c *ClientConf) Run() {
	for 9 != c.flag {
		if c.ListenShell() {
			switch c.flag {
			case 1:
				c.PublicChat()
			case 2:
				c.PrivateChat()
			case 3:
				c.ReName()
			case 9:
				c.ExitOut()
			default:
				fmt.Println("\n" + "请重新输入")
			}
		}
	}

}

//菜单显示
func (this *ClientConf) ListenShell() bool {
	fmt.Println("1:公聊模式")
	fmt.Println("2:私聊模式")
	fmt.Println("3:更新用户名")
	fmt.Println("9:退出")

	var param int
	fmt.Scanln(&param)
    this.flag = param

	if 0 > param || 3 < param {
		return false
	}
	return true
}

var (
	Cip   string
	Cport int
)

//获取命令行参数
func init() {
	flag.IntVar(&Cport, "p", 9999, "连接服务器的端口,默认是9999")
	flag.StringVar(&Cip, "i", "127.0.0.1", "连接服务器的ip,默认是127.0.0.1")
}

func index() {
	flag.Parse()

	conf := &ClientConf{
		Ip:   Cip,
		Port: Cport,
		flag: 99,
	}

	res := conf.NewClient()

	if nil == res {
		fmt.Println("连接失败")
		return
	}

	fmt.Println("连接成功")

    go conf.ListenServer()

	conf.Run()
}
