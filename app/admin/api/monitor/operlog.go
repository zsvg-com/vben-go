package monitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vben/app/admin/model/constants"
	"vben/app/admin/model/monitor"
	"vben/app/admin/model/tools"
	"vben/app/core/utils"
)

func ListOperlog(context *gin.Context) {
	var result = getListOperlog(context)
	context.JSON(http.StatusOK, result)
}

func getListOperlog(context *gin.Context) tools.TableDataInfo {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	var title = context.DefaultQuery("title", "")
	var operName = context.DefaultQuery("operName", "")
	var businessType = context.DefaultQuery("businessType", "")
	var status = context.DefaultQuery("status", "")

	var orderByColumn = context.DefaultQuery("orderByColumn", "")
	var isAsc = context.DefaultQuery("isAsc", "")

	var beginTime = context.DefaultQuery("params[beginTime]", "")
	var endTime = context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: monitor.SysOperLog{
			Title:        title,
			OperName:     operName,
			BusinessType: businessType,
			Status:       status,
		},
		OrderByColumn: orderByColumn,
		IsAsc:         isAsc,
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	result := monitor.SelectOperLogList(param)
	return result
}

func DelectOperlog(context *gin.Context) {
	var operId = context.Param("operId")
	var result = monitor.DelectOperlog(utils.Split(operId))
	context.JSON(http.StatusOK, result)
}

func ClearOperlog(context *gin.Context) {
	monitor.ClearOperlog()
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": "",
	})
}

func ExportOperlog(context *gin.Context) {
	var result = getListOperlog(context)
	var list = result.Rows.([]monitor.SysOperLog)

	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "operId",
		"title":  "操作序号",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "title",
		"title":  "操作模块",
		"width":  "15",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "businessType",
		"title":  "业务类型",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "method",
		"title":  "请求方法",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "requestMethod",
		"title":  "请求方式",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "operatorType",
		"title":  "操作类别",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "operName",
		"title":  "操作人员",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "deptName",
		"title":  "部门名称",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "operUrl",
		"title":  "请求地址",
		"width":  "60",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "operIp",
		"title":  "操作地址",
		"width":  "50",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "operLocation",
		"title":  "操作地点",
		"width":  "30",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "operParam",
		"title":  "请求参数",
		"width":  "30",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "jsonResult",
		"title":  "返回参数",
		"width":  "30",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "status",
		"title":  "状态",
		"width":  "30",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "errorMsg",
		"title":  "错误消息",
		"width":  "30",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "operTime",
		"title":  "操作时间",
		"width":  "30",
		"is_num": "0",
	})

	//填充数据
	data := make([]map[string]interface{}, 0)
	if len(list) > 0 {
		for _, v := range list {
			var businessType = v.BusinessType
			var businessTypeStr = ""
			if businessType == "0" {
				businessTypeStr = "其它"
			} else if businessType == "1" {
				businessTypeStr = "新增"
			} else if businessType == "2" {
				businessTypeStr = "修改"
			} else if businessType == "3" {
				businessTypeStr = "删除"
			} else if businessType == "4" {
				businessTypeStr = "授权"
			} else if businessType == "5" {
				businessTypeStr = "导出"
			} else if businessType == "6" {
				businessTypeStr = "导入"
			} else if businessType == "7" {
				businessTypeStr = "强退"
			} else if businessType == "8" {
				businessTypeStr = "生成代码"
			} else if businessType == "9" {
				businessTypeStr = "清空数据"
			}
			var operatorType = v.OperatorType
			var operatorTypeStr = ""
			if operatorType == "0" {
				operatorTypeStr = "其它"
			} else if operatorType == "1" {
				operatorTypeStr = "后台用户"
			} else if operatorType == "2" {
				operatorTypeStr = "手机端用户"
			}
			var status = v.Status
			var statusStr = ""
			if status == "0" {
				statusStr = "正常"
			} else {
				statusStr = "停用"
			}
			var operTime = v.OperTime.Format(constants.TimeFormat)
			data = append(data, map[string]interface{}{
				"operId":        v.OperId,
				"title":         v.Title,
				"businessType":  businessTypeStr,
				"method":        v.Method,
				"requestMethod": v.RequestMethod,
				"operatorType":  operatorTypeStr,
				"operName":      v.OperName,
				"deptName":      v.DeptName,
				"operUrl":       v.OperUrl,
				"operIp":        v.OperIp,
				"operLocation":  v.OperLocation,
				"operParam":     v.OperParam,
				"jsonResult":    v.JsonResult,
				"status":        statusStr,
				"errorMsg":      v.ErrorMsg,
				"operTime":      operTime,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}
