package monitor

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	utils2 "vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

// SysOperLog model：数据库字段
type SysOperLog struct {
	OperId        int       `json:"operId" gorm:"column:oper_id;primaryKey"` //表示主键
	Title         string    `json:"title" gorm:"title"`
	BusinessType  string    `json:"businessType" gorm:"business_type"`
	Method        string    `json:"method" gorm:"method"`
	RequestMethod string    `json:"requestMethod" gorm:"request_method"`
	OperatorType  string    `json:"operatorType" gorm:"operator_type"`
	OperName      string    `json:"operName" gorm:"oper_name"`
	DeptName      string    `json:"deptName" gorm:"dept_name"`
	OperUrl       string    `json:"operUrl" gorm:"oper_url"`
	OperIp        string    `json:"operIp" gorm:"oper_ip"`
	OperLocation  string    `json:"operLocation" gorm:"oper_location"`
	OperParam     string    `json:"operParam" gorm:"oper_param"`
	JsonResult    string    `json:"jsonResult" gorm:"json_result"`
	Status        string    `json:"status" gorm:"status"`
	ErrorMsg      string    `json:"errorMsg" gorm:"error_msg"`
	OperTime      time.Time `json:"operTime" gorm:"column:oper_time;type:datetime;autoCreateTime"`
}

// 表中的状态值
const (
	BusinessTypeOther = "0"
	BusinessTypeAdd   = "1"
	BusinessTypeEdit  = "2"
	BusinessTypeDel   = "3"

	OperatorTypeOther = "0"
	OperatorTypeAdmin = "1"
	OperatorTypePhone = "2"
)

// TableName 指定数据库表名称
func (SysOperLog) TableName() string {
	return "sys_oper_log"
}

func SelectOperLogList(params tools.SearchTableDataParam) tools.TableDataInfo {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	offset := (pageNum - 1) * pageSize

	sysOperLog := params.Other.(SysOperLog)

	var rows []SysOperLog

	var beginTime = params.Params.BeginTime
	var endTime = params.Params.EndTime

	var total int64
	var db = mysql.MysqlDb().Model(&rows)

	var title = sysOperLog.Title
	if title != "" {
		db.Where("title like concat('%', ?, '%')", title)
	}
	var businessType = sysOperLog.BusinessType
	if businessType != "" {
		db.Where("business_type = ?", businessType)
	}

	/*目前没有用*/
	//var businessTypes = sysOperLog.BusinessTypes
	//if businessTypes != nil || len(businessTypes) > 0 {
	//	//db.Table("users").Where("id in ? ",[]int{1,2,3}).Find(&user)
	//	db.Where("business_type in ?", businessTypes)
	//}

	var status = sysOperLog.Status
	if status != "" {
		db.Where("status = ?", status)
	}

	var oper_name = sysOperLog.OperName
	if oper_name != "" {
		db.Where("oper_name like concat('%', ?, '%')", oper_name)
	}
	var orderByColumn = params.OrderByColumn
	var isAsc = params.IsAsc

	if orderByColumn != "" {
		if "ascending" == isAsc {
			if "operTime" == orderByColumn {
				db.Order("oper_time DESC")
			}
			if "operName" == orderByColumn {
				db.Order("oper_name DESC")
			}
		}
		if "descending" == isAsc {
			if "operTime" == orderByColumn {
				db.Order("oper_time ASC")
			}
			if "operName" == orderByColumn {
				db.Order("oper_name ASC")
			}
		}
	}

	if beginTime != "" {
		//Loc, _ := time.LoadLocation("Asia/Shanghai")
		//startTime1, _ := time.ParseInLocation(constants.DateFormat, beginTime, Loc)
		//endTime = endTime + " 23:59:59"
		//endTime1, _ := time.ParseInLocation(constants.DateFormat, endTime, Loc)
		startTime1, endTime1 := utils2.GetBeginAndEndTime(beginTime, endTime)
		//db.Where("oper_time >= ? and oper_time <= ?", startTime1, endTime1)
		db.Where("oper_time >= ?", startTime1)
		db.Where("oper_time <= ?", endTime1)
	}

	db.Order("oper_id desc")
	if err := db.Table(sysOperLog.TableName()).Count(&total).Error; err != nil {
		return tools.Fail()
	}
	if err := db.Limit(pageSize).Offset(offset).Find(&rows).Error; err != nil {
		return tools.Fail()
	}

	if rows == nil {
		return tools.Fail()
	} else {
		return tools.Success(rows, total)
	}
}

func DelectOperlog(operIds []int) R.Result {
	if err := mysql.MysqlDb().Where("oper_id in (?)", operIds).Delete(&SysOperLog{}).Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func ClearOperlog() R.Result {
	if err := mysql.MysqlDb().Exec("truncate table sys_oper_log").Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func (param *SysOperLog) OperationLogAdd(context *gin.Context) R.Result {
	userId, _ := context.Get("userId")
	if userId != nil {
		sysUser := system.FindUserById(userId)
		dept := system.GetDeptInfo(strconv.Itoa(sysUser.DeptId))
		param.OperName = sysUser.UserName
		param.DeptName = dept.DeptName
		param.OperatorType = "1"
	}

	ip := utils2.GetRemoteClientIp(context.Request)
	method := context.Request.Method
	param.Method = method
	param.RequestMethod = method

	if method == "GET" {
		param.BusinessType = BusinessTypeOther
	}
	if method == "POST" {
		param.BusinessType = BusinessTypeAdd
	}
	if method == "DEL" {
		param.BusinessType = BusinessTypeDel
	}
	if method == "PUT" {
		param.BusinessType = BusinessTypeEdit
	}

	param.OperUrl = context.Request.URL.Path

	param.OperIp = ip
	param.OperTime = time.Now()
	param.OperLocation = "" + utils2.GetRealAddressByIP(ip)
	if err := mysql.MysqlDb().Model(&SysOperLog{}).Create(&param).Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}
