package system

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/twbworld/dating/global"
)

func tgStart() {
	if global.Config.Debug || len(global.Config.Telegram.Token) < 1 {
		return
	}

	bot, err := tg.NewBotAPI(global.Config.Telegram.Token)
	if err != nil {
		panic("bot初始化失败[jfsertyu]: " + err.Error())
	}
	global.Bot = bot
	global.Bot.Debug = global.Config.Debug

	setCommands := tg.NewSetMyCommands(tg.BotCommand{
		Command:     "start",
		Description: "开始",
	})
	if _, err := global.Bot.Request(setCommands); err != nil {
		panic("设置Command失败[podritgfd]: " + err.Error())
	}

	if global.Config.Domain != "" {
		wh, _ := tg.NewWebhook(fmt.Sprintf(`https://%s/wh/tg/%s`, global.Config.Domain, global.Bot.Token))
		if _, err = global.Bot.Request(wh); err != nil {
			panic("设置webhook失败[oifoghe]:" + err.Error())
		}

		info, err := global.Bot.GetWebhookInfo()
		if err != nil {
			panic("获取webhook失败[iuieee]:" + err.Error())
		}

		if info.LastErrorDate != 0 {
			panic("获取tg信息错误[fosdjfoisj]:" + info.LastErrorMessage)
		}
		global.Log.Printf("成功配置tg[doiasjo]: %s", global.Bot.Self.UserName)
	}
}

func tgClear() (err error) {
	if global.Bot == nil {
		return
	}
	_, err = global.Bot.Request(tg.DeleteWebhookConfig{})
	return
}
