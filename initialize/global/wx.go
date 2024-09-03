package global

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/twbworld/dating/global"
)

func initMiniProgram() {
	var err error
	global.MiniProgramApp, err = miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     global.Config.Weixin.XcxAppid,
		Secret:    global.Config.Weixin.XcxSecret,
		Debug:     global.Config.Debug,
		HttpDebug: global.Config.Debug,
		Log: miniProgram.Log{
			File:  `./log/wechat.log`,
			Error: `./log/wechat_error.log`,
		},
	})

	if err != nil {
		panic("小程序配置错误[opjkgh]]: " + err.Error())
	}
}
