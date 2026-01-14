package system

import (
	"sort"
	"time"
	"vben/app/admin/model/tools"
	"vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

// SysDictData model：数据库字段
type SysDictData struct {
	DictCode   int       `json:"dictCode" gorm:"column:dict_code;primaryKey"` //表示主键
	DictSort   int       `json:"dictSort" gorm:"dict_sort"`
	DictLabel  string    `json:"dictLabel" gorm:"dict_label"`
	DictValue  string    `json:"dictValue" gorm:"dict_value"`
	DictType   string    `json:"dictType" gorm:"dict_type"`
	CssClass   string    `json:"CssClass" gorm:"css_class"`
	ListClass  string    `json:"listClass" gorm:"list_class"`
	IsDefault  string    `json:"isDefault" gorm:"is_default"`
	Status     string    `json:"status" gorm:"status"`
	CreateBy   string    `json:"createBy" gorm:"create_by"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;type:datetime;autoCreateTime"`
	UpdateBy   string    `json:"updateBy" gorm:"update_by"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;type:datetime;autoCreateTime"`
	Remark     string    `json:"remark" gorm:"remark"`
}

func OrderByDictSortAsc(dictDataList []SysDictData) {
	sort.Sort(ByDictSort(dictDataList))
}

type ByDictSort []SysDictData

func (d ByDictSort) Len() int      { return len(d) }
func (d ByDictSort) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

// Less 正序
func (d ByDictSort) Less(i, j int) bool { return d[i].DictSort < d[j].DictSort }

// TableName 指定数据库表名称
func (SysDictData) TableName() string {
	return "sys_dict_data"
}

func SelectDictDataByType(dictType string) []SysDictData {
	var dictData []SysDictData
	mysql.MysqlDb().Raw("select * from sys_dict_data where status = '0' and dict_type = ? order by dict_sort", dictType).Scan(&dictData)
	return dictData
}

// 分页查询
func SelectDictDataList(params tools.SearchTableDataParam, isPage bool) tools.TableDataInfo {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysDictData := params.Other.(SysDictData)
	offset := (pageNum - 1) * pageSize
	var total int64
	var rows []SysDictData

	var db = mysql.MysqlDb().Model(&rows)

	var dictLabel = sysDictData.DictLabel
	if dictLabel != "" {
		db.Where("dict_label = ?", dictLabel)
	}
	var dictType = sysDictData.DictType
	if dictType != "" {
		db.Where("dict_type like concat('%', ?, '%')", dictType)
	}
	var status = sysDictData.Status
	if status != "" {
		db.Where("status = ?", status)
	}

	db.Order("dict_sort asc")

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

func FindDictCodeById(dictCode string) SysDictData {
	var dictData SysDictData
	err := mysql.MysqlDb().Where("dict_code = ?", dictCode).First(&dictData).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return dictData
}

func SaveDictData(dictDataParam SysDictData) R.Result {
	isDefault := dictDataParam.IsDefault
	if isDefault == "" {
		dictDataParam.IsDefault = "N"
	}
	err := mysql.MysqlDb().Model(&SysDictData{}).Create(&dictDataParam).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func EditDictData(dictDataParam SysDictData) R.Result {
	err := mysql.MysqlDb().Updates(&dictDataParam).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func DeleteDictData(dictCodes string) R.Result {
	var ids = utils.Split(dictCodes)

	// 获取事务对象并保存
	tx := mysql.MysqlDb().Begin()
	if tx.Error != nil {
		return R.ReturnFailMsg(tx.Error.Error())
	}

	for i := 0; i < len(ids); i++ {
		id := ids[i]
		err := tx.Where("dict_code = ?", id).Delete(&SysDictData{}).Error
		if err != nil {
			tx.Rollback()
			return R.ReturnFailMsg(err.Error())
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return R.ReturnFailMsg(err.Error())
	}

	return R.ReturnSuccess("操作成功")
}
