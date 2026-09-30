package group

import (
	"vben/common/json/timex"
	"vben/modules/sys/actor"
)

type SysGroup struct {
	Id      string           `json:"id" gorm:"type:varchar(36);primaryKey"` //主键ID
	Name    string           `json:"name" gorm:"type:varchar(255)"`         //名称
	Catid   string           `json:"catid" gorm:"type:varchar(36)"`         //分类ID
	Label   string           `json:"label"  gorm:"type:varchar(32)"`        //标签
	Notes   string           `json:"notes" gorm:"type:varchar(128)"`        //备注
	Ornum   int              `json:"ornum" gorm:"type:int"`                 //排序号
	Crtim   timex.Time       `json:"crtim"  gorm:"type:datetime"`           //创建时间
	Uptim   timex.Time       `json:"uptim" gorm:"type:datetime"`            //更新时间
	Cruid   string           `json:"cruid" gorm:"type:varchar(36)"`         //创建人ID
	Upuid   string           `json:"upuid" gorm:"type:varchar(36)"`         //更新人ID
	Avtag   bool             `json:"avtag"`
	Members []actor.SysActor `json:"members"  gorm:"many2many:sys_group_actor;foreignKey:Id;joinForeignKey:Gid;References:Id;joinReferences:Aid"` //包含成员
}

func (SysGroup) TableName() string {
	return "sys_group"
}

type SysGroupActor struct {
	Gid string `gorm:"primaryKey"`
	Aid string `gorm:"primaryKey"`
}

func (SysGroupActor) TableName() string {
	return "sys_group_actor"
}
