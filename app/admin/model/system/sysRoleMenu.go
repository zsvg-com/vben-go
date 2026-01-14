package system

import (
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

type SysRoleMenu struct {
	RoleId int `json:"roleId" gorm:"column:role_id"`
	MenuId int `json:"menuId" gorm:"column:menu_id"`
}

// TableName 指定数据库表名称
func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

// 删除角色与菜单关联
func DeleteRoleMenu(roleIds string) {
	sql := "delete from sys_role_menu where role_id in ( " + roleIds + " )"
	err := mysql.MysqlDb().Exec(sql).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
}

// 通过角色ID删除角色和菜单关联
func DeleteRoleMenuByRoleId(roleId string) {
	sql := "delete from sys_role_menu where role_id = " + roleId
	err := mysql.MysqlDb().Exec(sql).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
}
