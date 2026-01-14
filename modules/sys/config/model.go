package config

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type SysConfig struct {
	Id    int64      `json:"id" gorm:"type:bigint;primaryKey"` //主键ID
	Name  string     `json:"name" gorm:"type:varchar(255)"`    //名称
	Kenam string     `json:"kenam" gorm:"type:varchar(36)"`    //参数键名
	Keval string     `json:"keval" gorm:"type:varchar(128)"`   //参数键值
	Notes string     `json:"notes" gorm:"type:varchar(128)"`   //备注
	Ornum int        `json:"ornum" gorm:"type:int"`            //排序号
	Crtim timex.Time `json:"crtim"  gorm:"type:datetime"`      //创建时间
	Uptim timex.Time `json:"uptim" gorm:"type:datetime"`       //更新时间
	Intag boolx.Bool `json:"intag"`                            //内置标记
	Avtag boolx.Bool `json:"avtag"`
}

func (SysConfig) TableName() string {
	return "sys_config"
}
