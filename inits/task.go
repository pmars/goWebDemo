package inits

import (
	"fmt"

	"github.com/pmars/beego/logs"

	"github.com/robfig/cron"
)

// crontab 执行定时任务
/**
f1 f2 f3 f4 f5 f6 program
其中 f1 是表示秒，f2表示分钟，f3 表示小时，f4 表示一个月份中的第几日，f5 表示月份，f6表示一个星期中的第几天。
*/
func initTask() {
	fmt.Println("Start initTask")
	c := cron.New()

	c.AddFunc("1 1 1 * * 1", func() {
		// 将任务放到这里就可以
		logs.Debug("这段文字将会被定期执行打印")
	})

	// 启动定时任务
	c.Start()
}
