package api

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type SysApi struct {
	Id    int64      `json:"id" gorm:"type:bigint;primaryKey"` //主键ID
	Name  string     `json:"name" gorm:"type:varchar(255)"`    //名称
	Menid int64      `json:"menid" gorm:"type:bigint"`         //所属菜单ID
	Type  string     `json:"type"  gorm:"type:varchar(32)"`    //类型
	Perm  string     `json:"perm"  gorm:"type:varchar(64)"`    //权限字符
	Code  int64      `json:"code"  gorm:"type:bigint"`         //权限代码
	Pos   int        `json:"pos"  gorm:"type:int"`             //权限位
	Notes string     `json:"notes" gorm:"type:varchar(128)"`   //备注
	Ornum int        `json:"ornum" gorm:"type:int"`            //排序号
	Crtim timex.Time `json:"crtim"  gorm:"type:datetime"`      //创建时间
	Uptim timex.Time `json:"uptim" gorm:"type:datetime"`       //更新时间
	//Cruid string     `json:"cruid" gorm:"type:varchar(36)"`    //创建人ID
	//Upuid string     `json:"upuid" gorm:"type:varchar(36)"`    //更新人ID
	Avtag boolx.Bool `json:"avtag"` //可用标记
}

func (SysApi) TableName() string {
	return "sys_api"
}
