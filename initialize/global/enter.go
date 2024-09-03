package global

import (
	"flag"
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/twbworld/dating/global"
)

func init() {
	var config string

	//从命令参数获取配置路径
	flag.StringVar(&config, "c", "", "choose config file.")
	if !flag.Parsed() {
		testing.Init() //避免 单元测试(go test)自动加参数, 导致flag报错
		flag.Parse()
	}
	if config == "" {
		config = `config.yaml`
	}

	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic("读取配置失败[u9ij]: " + err.Error())
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件变化[djiads]: ", e.Name)
		if err := v.Unmarshal(global.Config); err != nil {
			if err := v.Unmarshal(global.Config); err != nil {
				fmt.Println(err)
			}
		}
	})
	// 将配置赋值给全局变量(结构体需要设置mapstructure的tag)
	if err := v.Unmarshal(global.Config); err != nil {
		panic("出错[dhfal]: " + err.Error())
	}
}

func Start() {
	initLog()
	initTz()
	initMiniProgram()
}
