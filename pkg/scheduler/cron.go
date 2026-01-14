package scheduler

import (
	"vben/app/admin/model/monitor"
	"vben/app/admin/model/tools"
	"vben/app/admin/service/tesk"
	"vben/app/core/utils"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

var c *cron.Cron

// 初始化 定时
func InitCron() {
	c = cron.New()
	c.Start()
	openMethod()
	RunSqlCron()
}

func AddCronFunc(sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}

func RunCronFunc(sepc string, invokeTarget string) {
	AddCronFunc(sepc, func() {
		utils.Call(m, invokeTarget)
	})
}

var m map[string]interface{}

func openMethod() {
	m = make(map[string]interface{})
	m["ryTask.ryNoParams"] = tesk.NoParamsMethod
	m["ryTask.ryParams"] = tesk.ParamsMethod
	m["ryTask.ryMultipleParams"] = tesk.MultipleParamsMethod
}

func RunSqlCron() {

	taskList, _ := monitor.SelectJobList(tools.SearchTableDataParam{
		Other: monitor.SysJob{},
	}, false)
	if len(taskList) == 0 {
		return
	}

	for _, item := range taskList {
		var policy = item.MisfirePolicy

		concurrent := item.Concurrent
		invokeTarget := item.InvokeTarget
		expression := item.CronExpression
		// 获取参数
		if concurrent == 0 {
			if policy == 1 {
				RunCronFunc(expression, invokeTarget)
			} else if policy == 2 {
				utils.Call(m, invokeTarget)
			}
		}
	}
}
