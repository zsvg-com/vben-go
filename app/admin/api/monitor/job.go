package monitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vben/app/admin/model/monitor"
	"vben/app/admin/model/tools"
	"vben/app/core/utils/R"
)

func ListJob(context *gin.Context) {
	rows, total := getListJob(context, true)
	if rows == nil {
		context.JSON(http.StatusOK, tools.Fail())
	} else {
		context.JSON(http.StatusOK, tools.Success(rows, total))
	}
}

func getListJob(context *gin.Context, isPage bool) ([]monitor.SysJob, int64) {
	/*分页*/
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	var jobName = context.DefaultQuery("jobName", "")
	var jobGroup = context.DefaultQuery("jobGroup", "")
	var status = context.DefaultQuery("status", "")
	var invokeTarget = context.DefaultQuery("invokeTarget", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: monitor.SysJob{
			JobName:      jobName,
			JobGroup:     jobGroup,
			Status:       status,
			InvokeTarget: invokeTarget,
		},
	}
	return monitor.SelectJobList(param, isPage)
}

func ExportJob(context *gin.Context) {
	list, _ := getListJob(context, false)
	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "jobId",
		"title":  "任务序号",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "jobName",
		"title":  "任务名称",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "jobGroup",
		"title":  "任务组名",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "invokeTarget",
		"title":  "调用目标字符串",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "cronExpression",
		"title":  "执行表达式",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "misfirePolicy",
		"title":  "计划策略",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "concurrent",
		"title":  "并发执行",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "status",
		"title":  "任务状态",
		"width":  "10",
		"is_num": "0",
	})
	//填充数据
	data := make([]map[string]interface{}, 0)
	if len(list) > 0 {
		for _, v := range list {
			misfirePolicyKey := v.MisfirePolicy
			var misfirePolicy = ""
			if 0 == misfirePolicyKey {
				misfirePolicy = "默认"
			}
			if 1 == misfirePolicyKey {
				misfirePolicy = "立即触发执行"
			}
			if 2 == misfirePolicyKey {
				misfirePolicy = "触发一次执行"
			}
			if 3 == misfirePolicyKey {
				misfirePolicy = "不触发立即执行"
			}
			concurrentKey := v.Concurrent
			var concurrent = ""
			if 0 == concurrentKey {
				concurrent = "允许"
			}
			if 1 == concurrentKey {
				concurrent = "禁止"
			}
			statusKey := v.Concurrent
			var status = ""
			if 0 == statusKey {
				status = "正常"
			}
			if 1 == statusKey {
				status = "暂停"
			}
			data = append(data, map[string]interface{}{
				"jobId":          v.JobId,
				"jobName":        v.JobName,
				"jobGroup":       v.JobGroup,
				"cronExpression": v.CronExpression,
				"invokeTarget":   v.InvokeTarget,
				"misfirePolicy":  misfirePolicy,
				"concurrent":     concurrent,
				"status":         status,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}

func GetJobById(context *gin.Context) {
	jobId := context.Param("jobId")
	result := monitor.FindJobById(jobId)
	context.JSON(http.StatusOK, result)
}

func SaveJob(context *gin.Context) {
	var jobParam monitor.SysJobParam
	if err := context.ShouldBind(&jobParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	userId, _ := context.Get("userId")
	result := monitor.SaveJob(jobParam, userId)
	context.JSON(http.StatusOK, result)
}

func UploadJob(context *gin.Context) {
	var jobParam monitor.SysJob
	if err := context.ShouldBind(&jobParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	userId, _ := context.Get("userId")
	result := monitor.UploadJob(jobParam, userId)
	context.JSON(http.StatusOK, result)
}

func ChangeStatus(context *gin.Context) {
	jobId := context.Param("jobIds")
	status := context.Param("status")
	monitor.ChangeStatus(jobId, status)
	context.JSON(http.StatusOK, R.ReturnSuccess("操作成功"))
}

func RunJob(context *gin.Context) {
	//jobId := context.Param("jobId")
	//result := monitor.FindJobById(jobId)
	//println(result)

	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": "",
	})
}

func DelectJob(context *gin.Context) {
	var jobIds = context.Param("jobIds")
	monitor.DetectJob(jobIds)
	context.JSON(http.StatusOK, R.ReturnSuccess("操作成功"))
}

func ListJobLog(context *gin.Context) {
	rows, total := getListJobLogList(context)
	if rows == nil {
		context.JSON(http.StatusOK, tools.Fail())
	} else {
		context.JSON(http.StatusOK, tools.Success(rows, total))
	}
}

func getListJobLogList(context *gin.Context) ([]monitor.SysJobLog, int64) {
	/*分页*/
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	var jobName = context.DefaultQuery("jobName", "")
	var jobGroup = context.DefaultQuery("jobGroup", "")
	var status = context.DefaultQuery("status", "")

	var beginTime = context.DefaultQuery("params[beginTime]", "")
	var endTime = context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: monitor.SysJobLog{
			JobName:  jobName,
			JobGroup: jobGroup,
			Status:   status,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	return monitor.SelectJobLogList(param)
}

func ExportJobLog(context *gin.Context) {
	list, _ := getListJobLogList(context)
	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "jobLogId",
		"title":  "日志序号",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "jobName",
		"title":  "任务名称",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "jobGroup",
		"title":  "任务组名",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "invokeTarget",
		"title":  "调用目标字符串",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "jobMessage",
		"title":  "日志信息",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "status",
		"title":  "执行状态",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "exceptionInfo",
		"title":  "异常信息",
		"width":  "10",
		"is_num": "0",
	})
	//填充数据
	data := make([]map[string]interface{}, 0)
	if len(list) > 0 {
		for _, v := range list {
			statusStr := ""
			status := v.Status
			if status == "0" {
				statusStr = "正常"
			}
			if status == "1" {
				statusStr = "失败"
			}
			data = append(data, map[string]interface{}{
				"jobLogId":      v.JobLogId,
				"jobName":       v.JobName,
				"jobGroup":      v.JobGroup,
				"invokeTarget":  v.InvokeTarget,
				"jobMessage":    v.JobMessage,
				"status":        statusStr,
				"exceptionInfo": v.ExceptionInfo,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}

func GetJobLog(context *gin.Context) {
	var configId = context.Param("configId")
	result := monitor.FindJobLogById(configId)
	context.JSON(http.StatusOK, result)
}

func DetectJobLog(context *gin.Context) {
	var jobLogIds = context.Param("jobLogIds")
	monitor.DetectJobLog(jobLogIds)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": "",
	})
}

func ClearJobLog(context *gin.Context) {
	result := monitor.ClearJobLog()
	context.JSON(http.StatusOK, result)
}
