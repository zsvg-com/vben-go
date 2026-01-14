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

func LoginInformListHandler(context *gin.Context) {
	rows, total := getListLoginLog(context)
	if rows == nil {
		context.JSON(http.StatusOK, tools.Fail())
	} else {
		context.JSON(http.StatusOK, tools.Success(rows, total))
	}
}

func getListLoginLog(context *gin.Context) ([]monitor.SysLogininfor, int64) {
	/*分页*/
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))
	var ipaddr = context.DefaultQuery("ipaddr", "")
	var userName = context.DefaultQuery("userName", "")
	var status = context.DefaultQuery("status", "")

	var orderByColumn = context.DefaultQuery("orderByColumn", "")
	var isAsc = context.DefaultQuery("isAsc", "")

	var beginTime = context.DefaultQuery("params[beginTime]", "")
	var endTime = context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: monitor.SysLogininfor{
			Ipaddr:   ipaddr,
			UserName: userName,
			Status:   status,
		},
		OrderByColumn: orderByColumn,
		IsAsc:         isAsc,
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	return monitor.SelectLogininforList(param)
}

func ExportHandler(context *gin.Context) {
	list, _ := getListLoginLog(context)
	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "infoId",
		"title":  "序号",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "userName",
		"title":  "用户账号",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "status",
		"title":  "登录状态",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "ipaddr",
		"title":  "登录地址",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "loginLocation",
		"title":  "登录地点",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "browser",
		"title":  "浏览器",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "os",
		"title":  "操作系统",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "msg",
		"title":  "提示消息",
		"width":  "30",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "loginTime",
		"title":  "访问时间",
		"width":  "50",
		"is_num": "0",
	})
	//填充数据
	data := make([]map[string]interface{}, 0)
	if len(list) > 0 {
		for _, v := range list {
			var status = v.Status
			var statusStr = ""
			if status == "0" {
				statusStr = "成功"
			} else {
				statusStr = "失败"
			}

			var loginTime = v.LoginTime.Format(constants.TimeFormat)
			data = append(data, map[string]interface{}{
				"infoId":        v.InfoId,
				"userName":      v.UserName,
				"status":        statusStr,
				"ipaddr":        v.Ipaddr,
				"loginLocation": v.LoginLocation,
				"browser":       v.Browser,
				"os":            v.Os,
				"msg":           v.Msg,
				"loginTime":     loginTime,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}

func DeleteByIdHandler(context *gin.Context) {
	var operId = context.Param("infoIds")
	var result = monitor.DelectLoginlog(utils.Split(operId))
	context.JSON(http.StatusOK, result)
}

func CleanHandler(context *gin.Context) {
	monitor.ClearLoginlog()
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
	})
}

func UnlockHandler(context *gin.Context) {
	var userName = context.Param("userName")
	monitor.UnlockByUserName(userName)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
	})
}
