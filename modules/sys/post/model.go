package post

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
	"vben/modules/sys/org"
)

type SysPost struct {
	//主键ID
	Id string `json:"id" gorm:"primaryKey"`
	//名称
	Name string `json:"name"`
	//部门ID
	Depid string `json:"depid"`
	//部门名称
	Depna string `json:"depna" gorm:"-"`
	//层级,以“_”隔开
	Tier string `json:"tier"`
	//标签
	Label string `json:"label"`
	//备注
	Notes string `json:"notes"`
	//排序号
	Ornum int `json:"ornum"`
	//可用标记
	Avtag boolx.Bool `json:"avtag"`
	//创建时间
	Crtim timex.Time `json:"crtim"`
	//更新时间
	Uptim timex.Time `json:"uptim"`
	//创建人ID
	Cruid string `json:"cruid"`
	//更新人ID
	Upuid string `json:"upuid"`
	//包含成员
	Users []org.SysOrg `json:"users" gorm:"many2many:sys_post_org;foreignKey:Id;joinForeignKey:Pid;References:Id;joinReferences:Oid"`
}

func (SysPost) TableName() string {
	return "sys_post"
}

type SysPostOrg struct {
	Pid string `gorm:"primaryKey"`
	Oid string `gorm:"primaryKey"`

	//Post SysPost    `gorm:"foreignKey:pid;references:id"`
	//User org.SysOrg `gorm:"foreignKey:oid;references:id"`
}

func (SysPostOrg) TableName() string {
	return "sys_post_org"
}
