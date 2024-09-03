package initialize

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/twbworld/dating/global"
	initGlobal "github.com/twbworld/dating/initialize/global"
	"github.com/twbworld/dating/initialize/system"
	"github.com/twbworld/dating/router"
	"github.com/twbworld/dating/service"
	"github.com/twbworld/dating/utils"

	"github.com/gin-gonic/gin"
)

var server *http.Server

func init() {
	initGlobal.Start()

	ginfile, err := os.OpenFile(global.Config.GinLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("打开文件错误[nvcbjkdfgu]: " + err.Error())
	}
	gin.DefaultWriter, gin.DefaultErrorWriter = io.MultiWriter(ginfile), global.Log.Out //记录所有日志
	gin.DisableConsoleColor()                                                           //将日志写入文件时不需要控制台颜色
	mode := gin.ReleaseMode
	if global.Config.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	ginServer := gin.Default()
	router.Start(ginServer)

	// ginServer.Run(":80")
	server = &http.Server{
		Addr:    global.Config.GinAddr,
		Handler: ginServer,
	}
}

func Start() {
	s := system.Start()
	defer s.Stop()
	// service.Match(1)

	//协程启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Panic("服务出错[isjfio]: ", err.Error()) //外部并不能捕获Panic
		}
	}()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	global.Log.Infof("启动成功, version: %s, port: %s, pid: %d, mem: %gMiB", runtime.Version(), global.Config.GinAddr, syscall.Getpid(), utils.NumberFormat(float32(m.Alloc)/1024/1024))

	service.Service.UserServiceGroup.TgService.TgSend("启动成功")

	//监听关闭(ctrl+C)指令
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	<-ctx.Done() //阻塞等待

	//来到这 证明有关闭指令,将进行平滑优雅关闭服务

	global.Log.Infof("程序关闭中..., port: %s, pid: %d", global.Config.GinAddr, syscall.Getpid())

	stop()

	//给程序最多5秒处理余下请求
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//关闭监听端口
	if err := server.Shutdown(timeoutCtx); err != nil {
		global.Log.Panicln("服务关闭出错[oijojiud]", err)
	}
	service.Service.UserServiceGroup.TgService.TgSend("服务退出成功")
	global.Log.Infoln("服务退出成功")

}
