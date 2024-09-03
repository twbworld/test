package system

import (
	"github.com/robfig/cron/v3"
	"github.com/twbworld/dating/global"
	"github.com/twbworld/dating/task"
)

var c *cron.Cron

func timerStart() {
	var option []cron.Option
	// option = append(option, cron.WithSeconds()) //精确到秒
	c = cron.New(option...)

	_, err := c.AddFunc("0 3 * * *", func() {
		if err := task.Clean(); err != nil {
			global.Log.Errorf("任务出错[osjd]: %s", err)
		}
	})
	if err != nil {
		panic("添加定时出错[oirfgio]:" + err.Error())
	}

	c.Start() //已含协程
}

func timerStop() {
	c.Stop()
}
