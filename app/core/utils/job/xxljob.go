package job

import (
	"strconv"
	"vben/common/config"

	"github.com/xxl-job/xxl-job-executor-go"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
)

/*
*	https://github.com/gin-middleware/xxl-job-executor（最新版本）
*	https://github.com/PGshen/go-xxl-executor(之前版本)
 */
func InitXxlJobCron() xxl.Executor {

	//初始化执行器
	exec := xxl.NewExecutor(
		xxl.ServerAddr(config.XxlJob.AdminAddress),
		xxl.AccessToken(config.XxlJob.AccessToken),         //请求令牌(默认为空)
		xxl.ExecutorIp(config.XxlJob.Ip),                   //可自动获取
		xxl.ExecutorPort(strconv.Itoa(config.XxlJob.Port)), //默认9999（此处要与gin服务启动port必需一至）
		xxl.RegistryKey(config.XxlJob.AppName),             //执行器名称
	)

	if config.XxlJob.Enabled {
		exec.Init()

		//注册任务handler（测试）这里是JobHandler*的名字
		exec.RegTask("task.test", task.Test)
		exec.RegTask("task.test2", task.Test2)
		exec.RegTask("task.panic", task.Panic)
	}

	return exec
}
