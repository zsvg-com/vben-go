package system

import (
	"time"
	"vben/app/admin/model/tools"
	utils2 "vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

type SysNotice struct {
	NoticeId      int       `json:"noticeId" gorm:"column:notice_id;primaryKey"` //表示主键
	NoticeTitle   string    `json:"noticeTitle" gorm:"notice_title"`
	NoticeType    string    `json:"noticeType" gorm:"notice_type"`
	NoticeContent string    `json:"noticeContent" gorm:"notice_content"`
	Status        string    `json:"status" gorm:"status"`
	CreateBy      string    `json:"createBy" gorm:"create_by"`
	CreateTime    time.Time `json:"createTime" gorm:"column:create_time;type:datetime;autoCreateTime"`
	UpdateBy      string    `json:"updateBy" gorm:"update_by"`
	UpdateTime    time.Time `json:"updateTime" gorm:"column:update_time;type:datetime;autoCreateTime"`
	Remark        string    `json:"remark" gorm:"remark"`
}

// TableName 指定数据库表名称
func (SysNotice) TableName() string {
	return "sys_notice"
}

func SelectSysNoticeList(params tools.SearchTableDataParam, isPage bool) tools.TableDataInfo {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysRoles := params.Other.(SysNotice)

	offset := (pageNum - 1) * pageSize
	var total int64
	var rows []SysNotice
	var db = mysql.MysqlDb().Model(&rows)

	var noticeTitle = sysRoles.NoticeTitle
	if noticeTitle != "" {
		db.Where("notice_title like concat('%', ?, '%')", noticeTitle)
	}

	var noticeType = sysRoles.NoticeType
	if noticeType != "" {
		db.Where("notice_type = ?", noticeType)
	}

	var createBy = sysRoles.CreateBy
	if createBy != "" {
		db.Where("create_by like concat('%', ?, '%')", createBy)
	}

	var beginTime = params.Params.BeginTime
	var endTime = params.Params.EndTime

	if beginTime != "" {
		//Loc, _ := time.LoadLocation("Asia/Shanghai")
		//startTime1, _ := time.ParseInLocation(constants.DateFormat, beginTime, Loc)
		//endTime = endTime + " 23:59:59"
		//endTime1, _ := time.ParseInLocation(constants.TimeFormat, endTime, Loc)
		startTime1, endTime1 := utils2.GetBeginAndEndTime(beginTime, endTime)
		db.Where("`sys_role`.create_time >= ?", startTime1)
		db.Where("`sys_role`.create_time <= ?", endTime1)
	}

	if err := db.Count(&total).Error; err != nil {
		return tools.Fail()
	}

	if isPage {
		if err := db.Limit(pageSize).Offset(offset).Find(&rows).Error; err != nil {
			return tools.Fail()
		}
	} else {
		if err := db.Find(&rows).Error; err != nil {
			return tools.Fail()
		}
	}

	if rows == nil {
		return tools.Fail()
	} else {
		return tools.Success(rows, total)
	}

}

func FindNoticeInfoById(noticeId string) SysNotice {
	var notice SysNotice
	err := mysql.MysqlDb().Where("notice_id = ?", noticeId).First(&notice).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return notice
}

func SaveNotice(sysNotice SysNotice) R.Result {
	err := mysql.MysqlDb().Model(&SysNotice{}).Create(&sysNotice).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func UploadNotice(sysNotice SysNotice) R.Result {
	err := mysql.MysqlDb().Updates(&sysNotice).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func DeleteNotice(noticeIds string) R.Result {
	ids := utils2.Split(noticeIds)
	for i := 0; i < len(ids); i++ {
		id := ids[i]
		err := mysql.MysqlDb().Where("notice_id = ?", id).Delete(&SysNotice{}).Error
		if err != nil {
			return R.ReturnFailMsg(err.Error())
		}
	}
	return R.ReturnSuccess("操作成功")
}
