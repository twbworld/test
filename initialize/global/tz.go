package global

import (
	"time"

	"github.com/twbworld/dating/global"
)

func initTz() {
	Location, err := time.LoadLocation(global.Config.Tz)
	if err != nil {
		panic("时区配置失败[siortuj]: " + err.Error())
	}
	global.Tz = Location
}
