package post

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
	"vben/modules/sys/actor"
)

type SysPost struct {
	//主键ID
	Id string `json:"id" gorm:"primaryKey"`
	//名称
	Name string `json:"name"`
	//部门ID
	Orgid string `json:"orgid"`
	//部门名称
	Orgna string `json:"orgna" gorm:"-"`
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
	Users []actor.SysActor `json:"users" gorm:"many2many:sys_post_actor;foreignKey:Id;joinForeignKey:Pid;References:Id;joinReferences:Aid"`
}

func (SysPost) TableName() string {
	return "sys_post"
}

type SysPostActor struct {
	Pid string `gorm:"primaryKey"`
	Aid string `gorm:"primaryKey"`

	//Post SysPost    `gorm:"foreignKey:pid;references:id"`
	//User actor.SysActor `gorm:"foreignKey:oid;references:id"`
}

func (SysPostActor) TableName() string {
	return "sys_post_actor"
}
