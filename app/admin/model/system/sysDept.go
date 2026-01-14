package system

import (
	"time"
	"vben/app/admin/model/tools"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

// SysDept model：数据库字段
type SysDept struct {
	DeptId     int       `json:"deptId" gorm:"column:dept_id;primaryKey"` //表示主键
	ParentId   int       `json:"parentId" gorm:"parent_id"`
	Ancestors  string    `json:"ancestors" gorm:"ancestors"`
	DeptName   string    `json:"deptName" gorm:"dept_name"`
	OrderNum   int       `json:"orderNum" gorm:"order_num"`
	Leader     string    `json:"leader" gorm:"leader"`
	Phone      string    `json:"phone" gorm:"phone"`
	Email      string    `json:"email" gorm:"email"`
	Status     string    `json:"status" gorm:"status"`
	DelFlag    string    `json:"delFlag" gorm:"del_flag"`
	CreateBy   string    `json:"createBy" gorm:"create_by"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;type:datetime;autoCreateTime"`
	UpdateBy   string    `json:"updateBy" gorm:"update_by"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;type:datetime;autoCreateTime"`
}

type SysDeptResult struct {
	DeptId     int             `json:"deptId" gorm:"column:dept_id;primaryKey"` //表示主键
	ParentId   int             `json:"parentId" gorm:"parent_id"`
	Ancestors  string          `json:"ancestors" gorm:"ancestors"`
	DeptName   string          `json:"deptName" gorm:"dept_name"`
	OrderNum   int             `json:"orderNum" gorm:"order_num"`
	Leader     string          `json:"leader" gorm:"leader"`
	Phone      string          `json:"phone" gorm:"phone"`
	Email      string          `json:"email" gorm:"email"`
	Status     string          `json:"status" gorm:"status"`
	ParentName string          `json:"parentName" gorm:"parent_name"`
	Children   []SysDeptResult `json:"children"`
	CreateBy   string          `json:"createBy" gorm:"create_by"`
	CreateTime time.Time       `json:"createTime" gorm:"column:create_time;type:datetime"`
	UpdateBy   string          `json:"updateBy" gorm:"update_by"`
	UpdateTime time.Time       `json:"updateTime" gorm:"column:update_time;type:datetime"`
}

// TableName 指定数据库表名称
func (SysDept) TableName() string {
	return "sys_dept"
}

type SysDeptDto struct {
	Id int `json:"id"`
	/** 节点名称 */
	Label string `json:"label"`
	/** 子节点 */
	Children []SysDeptDto `json:"children"`
}

// 分页查询
func GetDeptList(params tools.SearchTableDataParam, isPaging bool) ([]SysDeptResult, int64) {
	var total int64
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	offset := (pageNum - 1) * pageSize
	dept := params.Other.(SysDept)

	var list []SysDeptResult
	var db = mysql.MysqlDb().Model(&[]SysDept{})

	var deptId = dept.DeptId
	if deptId != 0 {
		db.Where("dept_id = ?", deptId)
	}

	var parentId = dept.ParentId
	if parentId != 0 {
		db.Where("parent_id = ?", parentId)
	}

	var deptName = dept.DeptName
	if deptName != "" {
		db.Where("dept_name like concat('%', ?, '%')", deptName)
	}

	var status = dept.Status
	if status != "" {
		db.Where("status = ?", status)
	}
	db.Order("parent_id, order_num")

	if err := db.Count(&total).Error; err != nil {
		return nil, 0
	}
	if isPaging {
		if err := db.Offset(offset).Limit(pageSize).Find(&list).Error; err != nil {
			return nil, 0
		}
	} else {
		if err := db.Find(&list).Error; err != nil {
			return nil, 0
		}
	}

	return list, total
}

func GetDeptInfo(deptId string) SysDept {
	var dept SysDept
	err := mysql.MysqlDb().Where("dept_id = ?", deptId).First(&dept).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return dept
}

func SaveDept(sysDept SysDept) R.Result {
	sysDept.DelFlag = "0"
	err := mysql.MysqlDb().Model(&sysDept).Create(&sysDept).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func UpDataDept(sysDept SysDept) R.Result {
	err := mysql.MysqlDb().Updates(&sysDept).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

/*单个删除*/
func DeleteDept(deptId string) R.Result {
	//if (deptService.hasChildByDeptId(deptId))
	//        {
	//            return AjaxResult.error("存在下级部门,不允许删除");
	//        }
	//        if (deptService.checkDeptExistUser(deptId))
	//        {
	//            return AjaxResult.error("部门存在用户,不允许删除");
	//        }
	//是否有数据权限
	err := mysql.MysqlDb().Where("dept_id = ?", deptId).Delete(&SysDept{}).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

type TreeSelect struct {
	Id       int          `json:"id"`
	Label    string       `json:"label"`
	Children []TreeSelect `json:"children"`
}

func SelectDeptTreeList() []TreeSelect {
	var deptResults []TreeSelect
	var depts []SysDept
	err := mysql.MysqlDb().Where("del_flag = '0'").Order("parent_id, order_num").Find(&depts).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	for i := 0; i < len(depts); i++ {
		dept := depts[i]
		deptId := dept.DeptId
		name := dept.DeptName
		pId := dept.ParentId
		if pId == 0 {
			tChild := getChildList(depts, deptId)
			treeSelect := TreeSelect{
				Id:       deptId,
				Label:    name,
				Children: tChild,
			}
			deptResults = append(deptResults, treeSelect)
		}
	}

	return deptResults
}

func getChildList(depts []SysDept, deptId int) []TreeSelect {
	var tlist []TreeSelect
	for i := 0; i < len(depts); i++ {
		dept1 := depts[i]
		id := dept1.DeptId
		pId := dept1.ParentId
		name := dept1.DeptName

		if pId == deptId {
			tChild := getChildList(depts, id)
			tree := TreeSelect{
				Id:       id,
				Label:    name,
				Children: tChild,
			}
			tlist = append(tlist, tree)
		}

	}
	return tlist
}
