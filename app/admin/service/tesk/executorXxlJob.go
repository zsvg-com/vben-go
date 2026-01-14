package tesk

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/PGshen/go-xxl-executor/biz"
	"github.com/PGshen/go-xxl-executor/common"
	"github.com/PGshen/go-xxl-executor/handler"
)

// JobHandler实现要求
// 1. "继承"handler.MethodJobHandler
// 2. 业务逻辑实现写在Execute
// 3. Init(), Destroy()属于钩子方法
// 4. 关于receiver.Log,他将输出到xxl日志中。由于go没有提供类似threadLocal变量，只能自己携带了。。。
// 5. 关于common.Log，他是执行器自身的日志，需要同receiver.Log区分开

type XxlJobHandler struct {
	handler.MethodJobHandler
}

// xxl-job
func (receiver *XxlJobHandler) Execute(param handler.Param) biz.ReturnT {
	receiver.MethodJobHandler.Execute(param)
	common.Log.Info("Test...")
	jobParams := make(map[string]interface{})
	_ = json.Unmarshal([]byte(param.JobParam), &jobParams)
	if len(jobParams) < 1 {
		common.Log.Info("暂无参数 ")
	} else {
		times := int(jobParams["times"].(float64))
		common.Log.Info("It will cycle " + strconv.Itoa(times) + " times")
		for i := 0; i < times; i++ {
			common.Log.Info("Test running: " + strconv.Itoa(i))
			receiver.Log.Info("Test running: " + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}

	receiver.Log.Info("Info...")
	receiver.Log.Warn("Warn...")
	receiver.Log.Debug("Debug...")
	receiver.Log.Error("Error...")
	receiver.Log.Trace("Trace...")
	receiver.Log.Fatal("Fatal...")
	common.Log.Info("Finish work!!!")

	return biz.NewReturnT(common.SuccessCode, "Test JobHandler")
}

// 结束定时
func (receiver XxlJobHandler) Destroy() {
	common.Log.Info("destroy...")
}
