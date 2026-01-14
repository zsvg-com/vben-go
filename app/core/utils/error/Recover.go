package error

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"vben/app/admin/service/monitor"

	"github.com/gin-gonic/gin"
)

/*
异常处理方法里面panic
panic(R.ReturnFailMsg("hello"))
这个也写入到数据库里面或日志里面
*/
func Recover(context *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			data, _ := json.Marshal(r)
			monitor.SysOperLog.AddSysOperLog(context, string(data))
			context.JSON(http.StatusOK, r)
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			context.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	context.Next()
}
