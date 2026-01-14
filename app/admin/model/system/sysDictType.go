package system

import (
	"time"
	"vben/app/admin/model/tools"
	utils2 "vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

// SysDictType model：数据库字段
type SysDictType struct {
	DictId     int       `json:"dictId" gorm:"column:dict_id;primaryKey"` //表示主键
	DictName   string    `json:"dictName" gorm:"dict_name"`
	DictType   string    `json:"dictType" gorm:"dict_type"`
	Status     string    `json:"status" gorm:"status"`
	CreateBy   string    `json:"createBy" gorm:"create_by"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;type:datetime;autoCreateTime"`
	UpdateBy   string    `json:"updateBy" gorm:"update_by"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;type:datetime;autoCreateTime"`
	Remark     string    `json:"remark" gorm:"remark"`
}

// TableName 指定数据库表名称
func (SysDictType) TableName() string {
	return "sys_dict_type"
}

// 分页查询
func SelectSysDictTypeList(params tools.SearchTableDataParam, isPage bool) tools.TableDataInfo {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysDictType := params.Other.(SysDictType)
	offset := (pageNum - 1) * pageSize
	var total int64
	var rows []SysDictType

	var db = mysql.MysqlDb().Model(&rows)

	var dictName = sysDictType.DictName
	if dictName != "" {
		db.Where("dict_name like ?", "%"+dictName+"%")
	}

	var status = sysDictType.Status
	if status != "" {
		db.Where("status =", status)
	}

	var dictType = sysDictType.DictType
	if dictType != "" {
		db.Where("dict_type like ?", dictType)
	}

	var beginTime = params.Params.BeginTime
	var endTime = params.Params.EndTime

	if beginTime != "" {
		//Loc, _ := time.LoadLocation("Asia/Shanghai")
		//startTime1, _ := time.ParseInLocation(constants.DateFormat, beginTime, Loc)
		//endTime = endTime + " 23:59:59"
		//endTime1, _ := time.ParseInLocation(constants.TimeFormat, endTime, Loc)
		startTime1, endTime1 := utils2.GetBeginAndEndTime(beginTime, endTime)
		db.Where("create_time >= ?", startTime1)
		db.Where("create_time <= ?", endTime1)
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

func FindTypeDictById(dictId string) SysDictType {
	var dictType SysDictType
	err := mysql.MysqlDb().Where("dict_id = ?", dictId).First(&dictType).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return dictType
}

func SaveType(dictType SysDictType) R.Result {
	err := mysql.MysqlDb().Model(&SysDictType{}).Create(&dictType).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func UpdateType(dictType SysDictType) R.Result {
	err := mysql.MysqlDb().Updates(&dictType).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func DeleteDataType(dictIds string) R.Result {
	ids := utils2.Split(dictIds)
	for i := 0; i < len(ids); i++ {
		id := ids[i]
		err := mysql.MysqlDb().Where("dict_id = ?", id).Delete(&SysDictType{}).Error
		if err != nil {
			return R.ReturnFailMsg(err.Error())
		}
	}
	return R.ReturnSuccess("操作成功")
}

func GetOptionSelect() R.Result {
	var sysdictType []SysDictType
	err := mysql.MysqlDb().Find(&sysdictType).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return R.ReturnSuccess(sysdictType)
}
