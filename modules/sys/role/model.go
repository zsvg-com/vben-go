package role

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
	"vben/common/r"
)

type SysRole struct {
	Id    int64       `json:"id" gorm:"type:bigint;primaryKey"` //主键ID
	Name  string      `json:"name" gorm:"type:varchar(255)"`    //名称
	Type  int         `json:"type"  gorm:"type:int"`            //类型
	Scope int         `json:"scope"  gorm:"type:int"`           //数据权限
	Notes string      `json:"notes" gorm:"type:varchar(128)"`   //备注
	Ornum int         `json:"ornum" gorm:"type:int"`            //排序号
	Crtim timex.Time  `json:"crtim"  gorm:"type:datetime"`      //创建时间
	Uptim timex.Time  `json:"uptim" gorm:"type:datetime"`       //更新时间
	Cruid string      `json:"cruid" gorm:"type:varchar(36)"`    //创建人ID
	Cruna string      `json:"cruna" gorm:"-"`                   //创建人名称
	Upuid string      `json:"upuid" gorm:"type:varchar(36)"`    //更新人ID
	Upuna string      `json:"upuna" gorm:"-"`                   //更新人名称
	Avtag boolx.Bool  `json:"avtag"`                            //可用标记
	Orgs  []r.SidName `json:"orgs" gorm:"-"`                    //包含成员
	Menus []int64     `json:"menus" gorm:"-"`                   //包含菜单
	Apis  []int64     `json:"apis" gorm:"-"`                    //包含接口
}

func (SysRole) TableName() string {
	return "sys_role"
}

type SysRoleOrg struct {
	Rid int64  `gorm:"primaryKey"`
	Oid string `gorm:"primaryKey"`
}

func (SysRoleOrg) TableName() string {
	return "sys_role_org"
}

type SysRoleMenu struct {
	Rid int64 `gorm:"primaryKey"`
	Mid int64 `gorm:"primaryKey"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

type SysRoleApi struct {
	Rid int64 `gorm:"primaryKey"`
	Aid int64 `gorm:"primaryKey"`
}

func (SysRoleApi) TableName() string {
	return "sys_role_api"
}
