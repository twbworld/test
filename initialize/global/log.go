package global
import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/twbworld/dating/utils"

	"github.com/twbworld/dating/global"
)
func initLog() {
	if err := utils.CreateFile(global.Config.RunLogPath); err != nil {
		panic("创建文件错误[oirdtug]: " + err.Error())
	}

	global.Log = logrus.New()
	global.Log.SetFormatter(&logrus.JSONFormatter{})
	global.Log.SetLevel(logrus.InfoLevel)

	runfile, err := os.OpenFile(global.Config.RunLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("打开文件错误[0atrpf]: " + err.Error())
	}
	global.Log.SetOutput(io.MultiWriter(os.Stdout, runfile)) //同时输出到终端和日志
}
