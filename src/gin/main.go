package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"test/src/gin/dao"
	"test/src/gin/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()
	fi, _ := os.Create("gin.log") //记录到文件。
	gin.DefaultWriter = io.MultiWriter(fi)

	ginServer := gin.Default()

	router.Init(ginServer)

	// ginServer.Run(":8081")
	server := http.Server{Addr: ":8081", Handler: ginServer}

	dao.InitMysql()

	dao.InitModel()

	//协程启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()


	closeBy(&server)

}

//平滑优雅关闭服务
func closeBy(server *http.Server) {
	sb := make(chan os.Signal, 1)
	signal.Notify(sb, syscall.SIGINT, syscall.SIGTERM) //监听关闭(ctrl+C)指令
	<-sb                                               //阻塞等待

	//来到这 证明有关闭指令
	c, f := context.WithTimeout(context.Background(), 5*time.Second) //如果有连接就超时5s后关闭
	defer f()                                                        //关闭当前请求

	//关闭监听端口
	if err := server.Shutdown(c); nil != err {
		log.Fatal(err)
	}
	log.Println("server exiting...")
}
