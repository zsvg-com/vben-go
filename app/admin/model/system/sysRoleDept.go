package system

import (
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

type SysRoleDept struct {
	RoleId int `json:"roleId" gorm:"column:role_id"`
	DeptId int `json:"deptId" gorm:"column:dept_id"`
}

// TableName 指定数据库表名称
func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}

// 删除角色与部门关联
func DeleteRoleDept(roleIds string) {
	sql := "delete from sys_role_dept where role_id in ( " + roleIds + " )"
	err := mysql.MysqlDb().Exec(sql).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
}
func DeleteRoleDeptByRole(roleId string) {
	err := mysql.MysqlDb().Where("role_id = ?", roleId).Delete(&SysRoleDept{}).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
}
