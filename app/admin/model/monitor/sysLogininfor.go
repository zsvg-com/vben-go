package monitor

import (
	"time"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	utils2 "vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"

	"github.com/gin-gonic/gin"
	useragent "github.com/wenlng/go-user-agent"
)

// SysLogininfor model：数据库字段
type SysLogininfor struct {
	InfoId        int       `json:"infoId" gorm:"column:info_id;primaryKey"` //表示主键
	UserName      string    `json:"userName" gorm:"user_name"`
	Ipaddr        string    `json:"ipaddr" gorm:"ipaddr"`
	LoginLocation string    `json:"loginLocation" gorm:"login_location"`
	Browser       string    `json:"browser" gorm:"browser"`
	Os            string    `json:"os" gorm:"os"`
	Status        string    `json:"status" gorm:"status"`
	Msg           string    `json:"msg" gorm:"msg"`
	LoginTime     time.Time `json:"loginTime" gorm:"column:login_time;type:datetime"`
}

// TableName 指定数据库表名称
func (SysLogininfor) TableName() string {
	return "sys_logininfor"
}

// 分页查询
func SelectLogininforList(params tools.SearchTableDataParam) ([]SysLogininfor, int64) {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysLogininfo := params.Other.(SysLogininfor)
	var ipaddr = sysLogininfo.Ipaddr
	var userName = sysLogininfo.UserName
	var status = sysLogininfo.Status

	var orderByColumn = params.OrderByColumn
	var isAsc = params.IsAsc

	var beginTime = params.Params.BeginTime
	var endTime = params.Params.EndTime

	var total int64
	db := mysql.MysqlDb().Model(SysLogininfor{})

	if ipaddr != "" {
		db.Where("ipaddr like ?", ipaddr+"%")
	}
	if status != "" {
		db.Where("status = ?", status)
	}
	if userName != "" {
		db.Where("user_name like ?", userName+"%")
	}
	if beginTime != "" {
		//Loc, _ := time.LoadLocation("Asia/Shanghai")
		//startTime1, _ := time.ParseInLocation(constants.DateFormat, beginTime, Loc)
		//endTime = endTime + " 23:59:59"
		//endTime1, _ := time.ParseInLocation(constants.TimeFormat, endTime, Loc)
		startTime1, endTime1 := utils2.GetBeginAndEndTime(beginTime, endTime)
		//db.Where("login_time >= ? and login_time <= ?", startTime1, endTime1)
		db.Where("login_time >= ?", startTime1)
		db.Where("login_time <= ?", endTime1)
	}
	if orderByColumn != "" {
		if "ascending" == isAsc {
			if "loginTime" == orderByColumn {
				db.Order("login_time DESC")
			}
			if "userName" == orderByColumn {
				db.Order("user_name DESC")
			}
		}
		if "descending" == isAsc {
			if "loginTime" == orderByColumn {
				db.Order("login_time ASC")
			}
			if "userName" == orderByColumn {
				db.Order("user_name ASC")
			}
		}
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0
	}
	offset := (pageNum - 1) * pageSize
	var list []SysLogininfor
	db.Order("info_id DESC").Offset(offset).Limit(pageSize).Find(&list)
	return list, total
}

func LoginInfoAdd(context *gin.Context, param system.LoginParam, message string, loginSucess bool) R.Result {
	var status = "0"
	if loginSucess {
		status = "0"
	} else {
		status = "1"
	}
	userAgent := context.Request.Header.Get("User-Agent")
	Os := useragent.GetOsName(userAgent)
	browser := useragent.GetBrowserName(userAgent)
	var info = SysLogininfor{
		UserName:      param.UserName,
		Msg:           message,
		Ipaddr:        "" + context.ClientIP(),
		LoginLocation: "" + utils2.GetRealAddressByIP(context.ClientIP()),
		Browser:       "" + browser,
		Os:            "" + Os,
		Status:        status,
		LoginTime:     time.Now(),
	}
	if err := mysql.MysqlDb().Model(SysLogininfor{}).Create(&info).First(&SysLogininfor{}).Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func DelectLoginlog(operIds []int) R.Result {
	if err := mysql.MysqlDb().Model(&SysLogininfor{}).Delete("info_id in (?)", operIds).Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func ClearLoginlog() R.Result {
	if err := mysql.MysqlDb().Model(&SysLogininfor{}).Raw("truncate table sys_logininfor").Find(SysLogininfor{}).Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func UnlockByUserName(userName string) {
	/*在redis 里面删除*/

}
